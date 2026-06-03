package app

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/token"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/encrypt"
	"APT/utility/validate"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/app/member"
)

func (c *ControllerMember) Login(ctx context.Context, req *member.LoginReq) (res *member.LoginRes, err error) {
	var (
		tx              gdb.TX
		orm             *gdb.Model
		ormWhereBuilder *gdb.WhereBuilder
		insertResult    sql.Result
		lastInsertId    int64
		MemberTokenInfo *model.MemberIdentity
		smsEvent        = "login"
		smsMobile       string
		savePhone       string
	)
	if tx, err = g.DB().Begin(ctx); err != nil {
		goto ERR
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	orm = dao.PmsMember.Ctx(ctx)
	ormWhereBuilder = orm.Builder()

	savePhone = req.Phone

	res = new(member.LoginRes)
	switch req.LoginType {
	case "email":
		if err = service.BasicsEmsLog().VerifyCode(ctx, &input_basics.VerifyEmsCodeInp{
			Event: smsEvent,
			Email: req.Email,
			Code:  req.Code,
		}); err != nil {
			return
		}
		orm = orm.Where(dao.PmsMember.Columns().Mail, req.Email)
		break
	case "phone":
		// 根据区号来判断手机号是否正确，并进行标准化处理
		if !g.IsEmpty(req.Phone) && !g.IsEmpty(req.AreaNo) {
			// 使用手机号验证工具进行验证和标准化
			phoneResult := validate.IsPhoneNumberWithCountryCode(req.AreaNo, req.Phone)
			if !phoneResult.IsValid {
				// 手机号格式错误
				err = gerror.New(gi18n.T(ctx, "phone_format_error"))
				return
			}
		}

		switch req.AreaNo {
		case "+81":
			smsEvent = "login-ja"
			smsMobile = fmt.Sprintf("%s%s", req.AreaNo, req.Phone)
			break
		case "+82":
			smsEvent = "login-ko"
			smsMobile = fmt.Sprintf("%s%s", req.AreaNo, req.Phone)
			break
		case "+86":
			smsEvent = "login"
			smsMobile = fmt.Sprintf("%s%s", req.AreaNo, req.Phone)
			break
		default:
			smsEvent = "login-en"
			smsMobile = fmt.Sprintf("%s%s", req.AreaNo, req.Phone)
			break
		}
		if err = service.BasicsSmsLog().VerifyCode(ctx, &input_basics.VerifyCodeInp{
			Event:  smsEvent,
			Mobile: smsMobile,
			Code:   req.Code,
		}); err != nil {
			return
		}
		orm = orm.Where(dao.PmsMember.Columns().Phone, req.Phone)
		orm = orm.Where(dao.PmsMember.Columns().PhoneArea, req.AreaNo)
		break
	case "password":
		ormWhereBuilder = ormWhereBuilder.WhereOr(dao.PmsMember.Columns().Mail, req.Email)
		ormWhereBuilder = ormWhereBuilder.WhereOr(g.MapStrAny{
			dao.PmsMember.Columns().Phone:     req.Phone,
			dao.PmsMember.Columns().PhoneArea: req.AreaNo,
		})
		orm = orm.Where(ormWhereBuilder)
		orm = orm.Where(dao.PmsMember.Columns().Password, encrypt.Md5([]byte(req.Password)))
		break
	case "yahooOauth":
	case "googleOauth":
		orm = orm.Where(dao.PmsMember.Columns().Mail, req.Email)
		break
	}
	if err = orm.OmitEmptyWhere().Scan(&MemberTokenInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
		goto ERR
	}

	if req.LoginType == "phone" {
		// 如果没有找到用户，判断如果是日本手机并且手机号只有10位，则往前补0，再去查找一次
		if g.IsEmpty(MemberTokenInfo) && (req.AreaNo == "+81" || req.AreaNo == "81") && len(req.Phone) == 10 {
			savePhone = "0" + req.Phone
			if err = dao.PmsMember.Ctx(ctx).OmitEmptyWhere().Where(g.MapStrAny{
				dao.PmsMember.Columns().Phone:     "0" + req.Phone,
				dao.PmsMember.Columns().PhoneArea: req.AreaNo,
			}).Scan(&MemberTokenInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
				goto ERR
			}
		}

		// 如果没有找到用户，判断如果是日本手机并且手机号有11位，则往前减0，再去查找一次
		if g.IsEmpty(MemberTokenInfo) && (req.AreaNo == "+81" || req.AreaNo == "81") && len(req.Phone) == 11 {
			if err = dao.PmsMember.Ctx(ctx).OmitEmptyWhere().Where(g.MapStrAny{
				dao.PmsMember.Columns().Phone:     req.Phone[1:],
				dao.PmsMember.Columns().PhoneArea: req.AreaNo,
			}).Scan(&MemberTokenInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
				goto ERR
			}
		}

		// 如果没有找到用户，判断如果是台湾手机并且手机号只有9位，则往前补0，再去查找一次
		if g.IsEmpty(MemberTokenInfo) && (req.AreaNo == "+886" || req.AreaNo == "886") && len(req.Phone) == 9 {
			savePhone = "0" + req.Phone
			if err = dao.PmsMember.Ctx(ctx).OmitEmptyWhere().Where(g.MapStrAny{
				dao.PmsMember.Columns().Phone:     "0" + req.Phone,
				dao.PmsMember.Columns().PhoneArea: req.AreaNo,
			}).Scan(&MemberTokenInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
				goto ERR
			}
		}

		// 如果没有找到用户，判断如果是台湾手机并且手机号有10位，则往前减0，再去查找一次
		if g.IsEmpty(MemberTokenInfo) && (req.AreaNo == "+886" || req.AreaNo == "886") && len(req.Phone) == 10 {
			if err = dao.PmsMember.Ctx(ctx).OmitEmptyWhere().Where(g.MapStrAny{
				dao.PmsMember.Columns().Phone:     req.Phone[1:],
				dao.PmsMember.Columns().PhoneArea: req.AreaNo,
			}).Scan(&MemberTokenInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
				goto ERR
			}
		}
	}

	if g.IsEmpty(MemberTokenInfo) && req.LoginType != "password" {
		// 查询是否已注销
		if err = dao.PmsMemberCancel.Ctx(ctx).
			Where(&entity.PmsMemberCancel{
				Phone:       req.Phone,
				PhoneArea:   req.AreaNo,
				Mail:        req.Email,
				AuditStatus: 2,
			}).OmitEmptyWhere().
			Scan(&MemberTokenInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
			goto ERR
		}

		// 如果没有找到用户，判断如果是日本手机并且手机号只有10位，则往前补0，再去查找一次
		if !g.IsEmpty(req.AreaNo) && !g.IsEmpty(req.Phone) && g.IsEmpty(MemberTokenInfo) && (req.AreaNo == "+81" || req.AreaNo == "81") && len(req.Phone) == 10 {
			if err = dao.PmsMemberCancel.Ctx(ctx).
				Where(&entity.PmsMemberCancel{
					Phone:       "0" + req.Phone,
					PhoneArea:   req.AreaNo,
					Mail:        req.Email,
					AuditStatus: 2,
				}).OmitEmptyWhere().
				Scan(&MemberTokenInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
				goto ERR
			}
		}

		// 如果没有找到用户，判断如果是日本手机并且手机号有11位，则往前减0，再去查找一次
		if !g.IsEmpty(req.AreaNo) && !g.IsEmpty(req.Phone) && g.IsEmpty(MemberTokenInfo) && (req.AreaNo == "+81" || req.AreaNo == "81") && len(req.Phone) == 11 {
			if err = dao.PmsMemberCancel.Ctx(ctx).
				Where(&entity.PmsMemberCancel{
					Phone:       req.Phone[1:],
					PhoneArea:   req.AreaNo,
					Mail:        req.Email,
					AuditStatus: 2,
				}).OmitEmptyWhere().
				Scan(&MemberTokenInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
				goto ERR
			}
		}

		// 如果没有找到用户，判断如果是台湾手机并且手机号只有9位，则往前补0，再去查找一次
		if !g.IsEmpty(req.AreaNo) && !g.IsEmpty(req.Phone) && g.IsEmpty(MemberTokenInfo) && (req.AreaNo == "+886" || req.AreaNo == "886") && len(req.Phone) == 9 {
			if err = dao.PmsMemberCancel.Ctx(ctx).
				Where(&entity.PmsMemberCancel{
					Phone:       "0" + req.Phone,
					PhoneArea:   req.AreaNo,
					Mail:        req.Email,
					AuditStatus: 2,
				}).OmitEmptyWhere().
				Scan(&MemberTokenInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
				goto ERR
			}
		}

		// 如果没有找到用户，判断如果是台湾手机并且手机号有10位，则往前减0，再去查找一次
		if !g.IsEmpty(req.AreaNo) && !g.IsEmpty(req.Phone) && g.IsEmpty(MemberTokenInfo) && (req.AreaNo == "+886" || req.AreaNo == "886") && len(req.Phone) == 10 {
			if err = dao.PmsMemberCancel.Ctx(ctx).
				Where(&entity.PmsMemberCancel{
					Phone:       req.Phone[1:],
					PhoneArea:   req.AreaNo,
					Mail:        req.Email,
					AuditStatus: 2,
				}).OmitEmptyWhere().
				Scan(&MemberTokenInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
				goto ERR
			}
		}

		if !g.IsEmpty(MemberTokenInfo) {
			// 获取会员注销设置
			var Config *input_basics.GetConfigModel
			Config, err = service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{
				Group: "membercancelsetting",
			})
			if err != nil {
				// 未找到注销设置
				err = gerror.Wrap(err, gi18n.T(ctx, "not_found_unregister_settings"))
				return
			}

			// 1-开启的话默认不能继续注册  2-关闭的话可以继续注册
			IsAllowedRegister := Config.List["isAllowedRegister"]
			if IsAllowedRegister == 1 {
				// 会员已注销
				err = gerror.New(gi18n.T(ctx, "member_has_been_cancelled"))
				return
			}
		}
		res.Replenish = true
		// 注册逻辑
		InsertData := &entity.PmsMember{}
		InsertData = &entity.PmsMember{
			Phone:           savePhone,
			PhoneArea:       req.AreaNo,
			Mail:            req.Email,
			Source:          req.Source,
			LoginMode:       req.LoginType,
			LastLogin:       gtime.Now(),
			LastLoginIp:     ghttp.RequestFromCtx(ctx).GetClientIp(),
			RegisterIp:      ghttp.RequestFromCtx(ctx).GetClientIp(),
			RegisterTime:    gtime.Now(),
			RegisterMdCode:  req.MdCode,
			RegisterMpModel: req.MpModel,
			Referrer:        req.Referrer,
			LastReferrer:    req.Referrer,
		}

		if insertResult, err = orm.OmitEmptyData().Insert(InsertData); err != nil {
			goto ERR
		}
		if lastInsertId, err = insertResult.LastInsertId(); err != nil {
			goto ERR
		}
		if lastInsertId <= 0 {
			goto ERR
		}
		// 更新会员号
		MemberNo := fmt.Sprintf("888%05d", lastInsertId)
		if _, err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, lastInsertId).Update(g.MapStrAny{
			dao.PmsMember.Columns().MemberNo: MemberNo,
		}); err != nil {
			return
		}
		// 记录推荐记录
		if !g.IsEmpty(req.Referrer) {
			if _, err = dao.PmsReferrerLog.Ctx(ctx).Data(g.Slice{
				&entity.PmsReferrerLog{
					Referrer:     req.Referrer,
					MemberId:     gvar.New(lastInsertId).Int(),
					LastReferrer: req.Referrer,
				},
			}).InsertAndGetId(); err != nil {
				return
			}
		}

		MemberInfo := g.MapStrAny{
			dao.PmsMember.Columns().Id:       gvar.New(lastInsertId).Int(),
			dao.PmsMember.Columns().Phone:    req.Phone,
			dao.PmsMember.Columns().Mail:     req.Email,
			dao.PmsMember.Columns().Source:   req.Source,
			dao.PmsMember.Columns().MemberNo: MemberNo,
		}
		if err = gvar.New(MemberInfo).Struct(&MemberTokenInfo); err != nil {
			goto ERR
		}
		err = service.AppMember().RegisterMemberAward(ctx, tx, gvar.New(lastInsertId).Int())
	} else {

		// 获取会员信息
		var PmsMember entity.PmsMember
		if err = dao.PmsMember.Ctx(ctx).WherePri(MemberTokenInfo.Id).Scan(&PmsMember); err != nil {
			return
		}

		if g.IsEmpty(PmsMember) {
			// 会员信息不存在或已注销
			err = gerror.New(gi18n.T(ctx, "member_info_does_not_exist_or_cancelled"))
			return
		}

		// 判断会员启用禁用状态
		if PmsMember.Status == 2 {
			// 会员已禁用
			err = gerror.New(gi18n.T(ctx, "member_disabled"))
			return
		}

		PmsMemberInfoUpdate := g.MapStrAny{
			dao.PmsMember.Columns().LastLogin:   gtime.Now(),
			dao.PmsMember.Columns().LoginMode:   req.LoginType,
			dao.PmsMember.Columns().LastLoginIp: ghttp.RequestFromCtx(ctx).GetClientIp(),
		}

		if !g.IsEmpty(req.Referrer) && req.Referrer != MemberTokenInfo.Id {

			PmsReferrerLogUpdate := g.MapStrAny{
				dao.PmsReferrerLog.Columns().MemberId:     MemberTokenInfo.Id,
				dao.PmsReferrerLog.Columns().LastReferrer: req.Referrer,
			}

			if g.IsEmpty(MemberTokenInfo.Referrer) {
				PmsMemberInfoUpdate[dao.PmsMember.Columns().Referrer] = req.Referrer
				PmsReferrerLogUpdate[dao.PmsReferrerLog.Columns().Referrer] = req.Referrer
			}
			PmsMemberInfoUpdate[dao.PmsMember.Columns().LastReferrer] = req.Referrer

			// 插入记录
			if _, err = dao.PmsReferrerLog.Ctx(ctx).Data(PmsReferrerLogUpdate).InsertAndGetId(); err != nil {
				return
			}
		}
		// 更新登录时间和IP
		if _, err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, MemberTokenInfo.Id).Update(PmsMemberInfoUpdate); err != nil {
			goto ERR
		}
		res.Replenish = MemberTokenInfo.Replenish == "N"
	}
	MemberTokenInfo.LoginAt = gtime.Now()
	MemberTokenInfo.App = consts.AppMember
	if res.Token, res.Expires, err = token.MemberLogin(ctx, MemberTokenInfo); err != nil {
		goto ERR
	}
	// 记录登录记录
	if _, err = dao.PmsMemberLog.Ctx(ctx).Data(&entity.PmsMemberLog{
		MemberId:  MemberTokenInfo.Id,
		LoginTime: gtime.Now(),
		LoginType: req.LoginType,
		LoginIp:   ghttp.RequestFromCtx(ctx).GetClientIp(),
		Token:     res.Token,
		ExpirTime: gtime.New(res.Expires + gtime.Now().Unix()),
		MdCode:    req.MdCode,
		MpModel:   req.MpModel,
	}).Insert(); err != nil {
		return
	}
	return
ERR:
	g.Log().Error(ctx, err)
	err = errors.New("loginFail")
	return
}
func (c *ControllerMember) AuthIdBindLogin(ctx context.Context, req *member.AuthIdBindLoginReq) (res *member.AuthIdBindLoginRes, err error) {
	var (
		PmsMemberAuth        *entity.PmsMemberAuth
		PmsMemberAuthIsExist []*entity.PmsMemberAuth
		PmsMember            *entity.PmsMember
		insertResult         sql.Result
		lastInsertId         int64
		MemberTokenInfo      *model.MemberIdentity
		tx                   gdb.TX
		MemberInfo           g.MapStrAny
		MobileNo             string
		smsEvent             string
	)
	if tx, err = g.DB().Begin(ctx); err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	// 手动校验手机号和邮箱
	// PHONE,phone,EMAIL,email,WX_MINI,wx_mini
	if req.BindType == "PHONE" || req.BindType == "phone" || req.BindType == "WX_MINI" || req.BindType == "wx_mini" {
		if g.IsEmpty(gstr.Trim(req.Phone, "+=-")) || g.IsEmpty(gstr.Trim(req.AreaNo, "+-=")) {
			// 请选择手机号
			err = gerror.New(gi18n.T(ctx, "select_phone"))
			return
		}
	}
	if req.BindType == "EMAIL" || req.BindType == "email" {
		if g.IsEmpty(gstr.Trim(req.Email)) {
			// 请选择邮箱
			err = gerror.New(gi18n.T(ctx, "select_email"))
			return
		}
	}

	res = new(member.AuthIdBindLoginRes)
	// 查询是否存在该AuthId
	if err = dao.PmsMemberAuth.Ctx(ctx).Where(dao.PmsMemberAuth.Columns().AuthId, req.AuthId).Scan(&PmsMemberAuth); err != nil {
		return
	}
	if g.IsEmpty(PmsMemberAuth) {
		// 绑定失败
		err = gerror.New(gi18n.T(ctx, "bind_failed"))
		return
	}
	if !g.IsEmpty(PmsMemberAuth.MemberId) {
		// 该账号已被绑定
		err = gerror.New(gi18n.T(ctx, "account_already_bound"))
		return
	}
	switch req.BindType {
	case "PHONE", "phone":
		switch req.AreaNo {
		case "+81":
			smsEvent = "login-ja"
			MobileNo = fmt.Sprintf("%s%s", req.AreaNo, req.Phone)
			break
		case "+82":
			smsEvent = "login-ko"
			MobileNo = fmt.Sprintf("%s%s", req.AreaNo, req.Phone)
			break
		case "+86":
			smsEvent = "login"
			MobileNo = fmt.Sprintf("%s%s", req.AreaNo, req.Phone)
			break
		default:
			smsEvent = "login-en"
			MobileNo = fmt.Sprintf("%s%s", req.AreaNo, req.Phone)
			break
		}
		if err = service.BasicsSmsLog().VerifyCode(ctx, &input_basics.VerifyCodeInp{
			Event:  smsEvent,
			Mobile: MobileNo,
			Code:   req.Code,
		}); err != nil {
			return
		}
		break
	case "EMAIL", "email":
		if err = service.BasicsEmsLog().VerifyCode(ctx, &input_basics.VerifyEmsCodeInp{
			Event: "login",
			Email: req.Email,
			Code:  req.Code,
		}); err != nil {
			return
		}
		break
	}
	// 查询是否存在用户信息
	if !g.IsEmpty(req.Email) {
		req.AreaNo = ""
	}

	savePhone := req.Phone

	// 根据区号来判断手机号是否正确，并进行标准化处理
	if !g.IsEmpty(req.Phone) && !g.IsEmpty(req.AreaNo) {
		// 使用手机号验证工具进行验证和标准化
		phoneResult := validate.IsPhoneNumberWithCountryCode(req.AreaNo, req.Phone)
		if !phoneResult.IsValid {
			// 手机号格式错误
			err = gerror.New(gi18n.T(ctx, "phone_format_error"))
			return
		}
	}

	if err = dao.PmsMember.Ctx(ctx).OmitEmptyWhere().Where(g.MapStrAny{
		dao.PmsMember.Columns().Phone:     req.Phone,
		dao.PmsMember.Columns().PhoneArea: req.AreaNo,
		dao.PmsMember.Columns().Mail:      req.Email,
	}).Scan(&PmsMember); err != nil {
		return
	}

	// 如果没有找到用户，判断如果是日本手机并且手机号只有10位，则往前补0，再去查找一次
	if g.IsEmpty(PmsMember) && (req.AreaNo == "+81" || req.AreaNo == "81") && len(req.Phone) == 10 {
		savePhone = "0" + req.Phone
		if err = dao.PmsMember.Ctx(ctx).OmitEmptyWhere().Where(g.MapStrAny{
			dao.PmsMember.Columns().Phone:     "0" + req.Phone,
			dao.PmsMember.Columns().PhoneArea: req.AreaNo,
			dao.PmsMember.Columns().Mail:      req.Email,
		}).Scan(&PmsMember); err != nil {
			return
		}
	}

	// 如果没有找到用户，判断如果是日本手机并且手机号有11位，则往前减0，再去查找一次
	if g.IsEmpty(PmsMember) && (req.AreaNo == "+81" || req.AreaNo == "81") && len(req.Phone) == 11 {
		if err = dao.PmsMember.Ctx(ctx).OmitEmptyWhere().Where(g.MapStrAny{
			dao.PmsMember.Columns().Phone:     req.Phone[1:],
			dao.PmsMember.Columns().PhoneArea: req.AreaNo,
			dao.PmsMember.Columns().Mail:      req.Email,
		}).Scan(&PmsMember); err != nil {
			return
		}
	}

	// 如果没有找到用户，判断如果是台湾手机并且手机号只有9位，则往前补0，再去查找一次
	if g.IsEmpty(PmsMember) && (req.AreaNo == "+886" || req.AreaNo == "886") && len(req.Phone) == 9 {
		savePhone = "0" + req.Phone
		if err = dao.PmsMember.Ctx(ctx).OmitEmptyWhere().Where(g.MapStrAny{
			dao.PmsMember.Columns().Phone:     "0" + req.Phone,
			dao.PmsMember.Columns().PhoneArea: req.AreaNo,
			dao.PmsMember.Columns().Mail:      req.Email,
		}).Scan(&PmsMember); err != nil {
			return
		}
	}

	// 如果没有找到用户，判断如果是台湾手机并且手机号有10位，则往前减0，再去查找一次
	if g.IsEmpty(PmsMember) && (req.AreaNo == "+886" || req.AreaNo == "886") && len(req.Phone) == 10 {
		if err = dao.PmsMember.Ctx(ctx).OmitEmptyWhere().Where(g.MapStrAny{
			dao.PmsMember.Columns().Phone:     req.Phone[1:],
			dao.PmsMember.Columns().PhoneArea: req.AreaNo,
			dao.PmsMember.Columns().Mail:      req.Email,
		}).Scan(&PmsMember); err != nil {
			return
		}
	}

	if g.IsEmpty(PmsMember) {
		// 查询是否已注销
		if err = dao.PmsMemberCancel.Ctx(ctx).
			Where(&entity.PmsMemberCancel{
				Phone:       req.Phone,
				PhoneArea:   req.AreaNo,
				Mail:        req.Email,
				AuditStatus: 2,
			}).OmitEmptyWhere().
			Scan(&MemberTokenInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
			g.Log().Error(ctx, err)
			// 登录失败
			err = gerror.New(gi18n.T(ctx, "login_failed"))
			return
		}

		// 如果没有找到用户，判断如果是日本手机并且手机号只有10位，则往前补0，再去查找一次
		if g.IsEmpty(MemberTokenInfo) && (req.AreaNo == "+81" || req.AreaNo == "81") && len(req.Phone) == 10 {
			if err = dao.PmsMemberCancel.Ctx(ctx).
				Where(&entity.PmsMemberCancel{
					Phone:       "0" + req.Phone,
					PhoneArea:   req.AreaNo,
					Mail:        req.Email,
					AuditStatus: 2,
				}).OmitEmptyWhere().
				Scan(&MemberTokenInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
				g.Log().Error(ctx, err)
				// 登录失败
				err = gerror.New(gi18n.T(ctx, "login_failed"))
				return
			}
		}

		// 如果没有找到用户，判断如果是日本手机并且手机号有11位，则往前减0，再去查找一次
		if g.IsEmpty(MemberTokenInfo) && (req.AreaNo == "+81" || req.AreaNo == "81") && len(req.Phone) == 11 {
			if err = dao.PmsMemberCancel.Ctx(ctx).
				Where(&entity.PmsMemberCancel{
					Phone:       req.Phone[1:],
					PhoneArea:   req.AreaNo,
					Mail:        req.Email,
					AuditStatus: 2,
				}).OmitEmptyWhere().
				Scan(&MemberTokenInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
				g.Log().Error(ctx, err)
				// 登录失败
				err = gerror.New(gi18n.T(ctx, "login_failed"))
				return
			}
		}

		// 如果没有找到用户，判断如果是台湾手机并且手机号只有9位，则往前补0，再去查找一次
		if g.IsEmpty(MemberTokenInfo) && (req.AreaNo == "+886" || req.AreaNo == "886") && len(req.Phone) == 9 {
			if err = dao.PmsMemberCancel.Ctx(ctx).
				Where(&entity.PmsMemberCancel{
					Phone:       "0" + req.Phone,
					PhoneArea:   req.AreaNo,
					Mail:        req.Email,
					AuditStatus: 2,
				}).OmitEmptyWhere().
				Scan(&MemberTokenInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
				g.Log().Error(ctx, err)
				// 登录失败
				err = gerror.New(gi18n.T(ctx, "login_failed"))
				return
			}
		}

		// 如果没有找到用户，判断如果是台湾手机并且手机号有10位，则往前减0，再去查找一次
		if g.IsEmpty(MemberTokenInfo) && (req.AreaNo == "+886" || req.AreaNo == "886") && len(req.Phone) == 10 {
			if err = dao.PmsMemberCancel.Ctx(ctx).
				Where(&entity.PmsMemberCancel{
					Phone:       req.Phone[1:],
					PhoneArea:   req.AreaNo,
					Mail:        req.Email,
					AuditStatus: 2,
				}).OmitEmptyWhere().
				Scan(&MemberTokenInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
				g.Log().Error(ctx, err)
				// 登录失败
				err = gerror.New(gi18n.T(ctx, "login_failed"))
				return
			}
		}

		if !g.IsEmpty(MemberTokenInfo) {
			// 获取会员注销设置
			var Config *input_basics.GetConfigModel
			Config, err = service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{
				Group: "membercancelsetting",
			})
			if err != nil {
				// 未找到注销设置
				err = gerror.Wrap(err, gi18n.T(ctx, "not_found_unregister_settings"))
				return
			}

			IsAllowedRegister := Config.List["isAllowedRegister"]
			if IsAllowedRegister == 1 {
				// 会员已注销
				err = gerror.New(gi18n.T(ctx, "member_has_been_cancelled"))
				return
			}
		}
		// 注册逻辑
		InsertData := &entity.PmsMember{
			Phone:           savePhone,
			PhoneArea:       req.AreaNo,
			Mail:            req.Email,
			Source:          req.Source,
			LoginMode:       PmsMemberAuth.Channel,
			LastLogin:       gtime.Now(),
			LastLoginIp:     ghttp.RequestFromCtx(ctx).GetClientIp(),
			RegisterIp:      ghttp.RequestFromCtx(ctx).GetClientIp(),
			RegisterTime:    gtime.Now(),
			RegisterMdCode:  req.MdCode,
			RegisterMpModel: req.MpModel,
			Referrer:        req.Referrer,
			LastReferrer:    req.Referrer,
		}
		if insertResult, err = dao.PmsMember.Ctx(ctx).TX(tx).OmitEmptyData().Insert(InsertData); err != nil {
			return
		}
		if lastInsertId, err = insertResult.LastInsertId(); err != nil {
			return
		}
		if lastInsertId <= 0 {
			// 绑定失败
			err = gerror.New(gi18n.T(ctx, "bind_failed"))
			return
		}
		// 更新会员号
		MemberNo := fmt.Sprintf("888%05d", lastInsertId)
		if _, err = dao.PmsMember.Ctx(ctx).TX(tx).Where(dao.PmsMember.Columns().Id, lastInsertId).Update(g.MapStrAny{
			dao.PmsMember.Columns().MemberNo: MemberNo,
		}); err != nil {
			return
		}

		// 记录推荐记录
		if !g.IsEmpty(req.Referrer) {
			if _, err = dao.PmsReferrerLog.Ctx(ctx).Data(g.Slice{
				&entity.PmsReferrerLog{
					Referrer:     req.Referrer,
					MemberId:     gvar.New(lastInsertId).Int(),
					LastReferrer: req.Referrer,
				},
			}).InsertAndGetId(); err != nil {
				return
			}
		}

		MemberInfo = g.MapStrAny{
			dao.PmsMember.Columns().Id:       gvar.New(lastInsertId).Int(),
			dao.PmsMember.Columns().Phone:    savePhone,
			dao.PmsMember.Columns().Mail:     req.Email,
			dao.PmsMember.Columns().Source:   req.Source,
			dao.PmsMember.Columns().MemberNo: MemberNo,
		}
		res.Replenish = true
		err = service.AppMember().RegisterMemberAward(ctx, tx, gvar.New(lastInsertId).Int())
	} else {
		// 查询是否同时存在该用户的其他绑定关系
		if err = dao.PmsMemberAuth.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsMemberAuth.Columns().MemberId: PmsMember.Id,
			dao.PmsMemberAuth.Columns().Channel:  PmsMemberAuth.Channel,
		}).Scan(&PmsMemberAuthIsExist); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		// 如果存在用户绑定关系，则移除绑定关系
		if !g.IsEmpty(PmsMemberAuthIsExist) {
			// 该账号已被绑定
			err = gerror.New(gi18n.T(ctx, "account_already_bound"))
			return
			//if _, err = dao.PmsMemberAuth.Ctx(ctx).
			//	Where(g.MapStrAny{
			//		dao.PmsMemberAuth.Columns().MemberId: PmsMember.Id,
			//		dao.PmsMemberAuth.Columns().Channel:  PmsMemberAuth.Channel,
			//	}).Delete(); err != nil {
			//	return
			//}
		}

		if _, err = dao.PmsMemberAuth.Ctx(ctx).Where(dao.PmsMemberAuth.Columns().AuthId, req.AuthId).Data(g.MapStrAny{
			dao.PmsMemberAuth.Columns().MemberId: PmsMember.Id,
		}).Update(); err != nil {
			return
		}
		MemberInfo = g.MapStrAny{
			dao.PmsMember.Columns().Id:       PmsMember.Id,
			dao.PmsMember.Columns().Phone:    req.Phone,
			dao.PmsMember.Columns().Mail:     req.Email,
			dao.PmsMember.Columns().Source:   req.Source,
			dao.PmsMember.Columns().MemberNo: PmsMember.MemberNo,
		}
		res.Replenish = PmsMember.Replenish == "N"

		PmsMemberInfoUpdate := g.MapStrAny{
			dao.PmsMember.Columns().LastLogin:   gtime.Now(),
			dao.PmsMember.Columns().LoginMode:   PmsMemberAuth.Channel,
			dao.PmsMember.Columns().LastLoginIp: ghttp.RequestFromCtx(ctx).GetClientIp(),
		}

		if !g.IsEmpty(req.Referrer) && req.Referrer != PmsMember.Id {

			PmsReferrerLogUpdate := g.MapStrAny{
				dao.PmsReferrerLog.Columns().MemberId:     PmsMember.Id,
				dao.PmsReferrerLog.Columns().LastReferrer: req.Referrer,
			}

			if g.IsEmpty(PmsMember.Referrer) {
				PmsMemberInfoUpdate[dao.PmsMember.Columns().Referrer] = req.Referrer
				PmsReferrerLogUpdate[dao.PmsReferrerLog.Columns().Referrer] = req.Referrer
			}
			PmsMemberInfoUpdate[dao.PmsMember.Columns().LastReferrer] = req.Referrer

			// 插入记录
			if _, err = dao.PmsReferrerLog.Ctx(ctx).Data(PmsReferrerLogUpdate).InsertAndGetId(); err != nil {
				return
			}
		}

		if _, err = dao.PmsMember.Ctx(ctx).TX(tx).
			Where(dao.PmsMember.Columns().Id, PmsMember.Id).
			OmitEmptyData().Update(PmsMemberInfoUpdate); err != nil {
			return
		}

	}

	if err = gvar.New(MemberInfo).Struct(&MemberTokenInfo); err != nil {
		return
	}
	MemberTokenInfo.Channel = PmsMemberAuth.Channel
	if req.Source == "H5_FX" {
		MemberTokenInfo.IsFx = true
	}
	MemberTokenInfo.PmsMemberAuth = PmsMemberAuth
	if res.Token, res.Expires, err = token.MemberLogin(ctx, MemberTokenInfo); err != nil {
		return
	}
	if _, err = dao.PmsMemberAuth.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.PmsMemberAuth.Columns().AuthId: req.AuthId,
	}).Update(g.MapStrAny{
		dao.PmsMemberAuth.Columns().MemberId: MemberTokenInfo.Id,
	}); err != nil {
		return
	}
	// 记录登录记录
	if _, err = dao.PmsMemberLog.Ctx(ctx).TX(tx).Data(&entity.PmsMemberLog{
		MemberId:  MemberTokenInfo.Id,
		LoginTime: gtime.Now(),
		LoginType: PmsMemberAuth.Channel,
		LoginIp:   ghttp.RequestFromCtx(ctx).GetClientIp(),
		Token:     res.Token,
		ExpirTime: gtime.New(res.Expires + gtime.Now().Unix()),
		MdCode:    req.MdCode,
		MpModel:   req.MpModel,
	}).Insert(); err != nil {
		return
	}

	return
}
