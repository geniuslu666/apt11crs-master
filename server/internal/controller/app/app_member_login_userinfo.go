package app

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
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
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"

	"APT/api/app/member"
)

func (c *ControllerMember) GetUserInfo(ctx context.Context, _ *member.GetUserInfoReq) (res *member.GetUserInfoRes, err error) {
	var (
		UserInfo *model.MemberIdentity
	)
	UserInfo = contexts.GetMemberUser(ctx)
	res = new(member.GetUserInfoRes)
	if err = g.Model(dao.PmsMember.Table()).Ctx(ctx).WithAll().Where(dao.PmsMember.Columns().Id, UserInfo.Id).Scan(&res); err != nil && !errors.Is(err, sql.ErrNoRows) {
		// 获取用户信息失败,请重新登录
		err = gerror.NewCode(gcode.CodeNotAuthorized, gi18n.T(ctx, "failed_to_get_member_info"))
		return
	}

	// 获取配置
	var (
		AppDataBoardViewConfig *model.AppDataBoardViewConfig
	)
	if AppDataBoardViewConfig, err = service.BasicsConfig().GetAppDataBoardViewConfig(ctx); err == nil {
		if AppDataBoardViewConfig.TestType == 1 {
			// 会员分组
			pmsMemberGroupIds := AppDataBoardViewConfig.CanTestGroupIds
			pmsMemberGroupIdsList := strings.Split(pmsMemberGroupIds, ",")
			for _, v := range pmsMemberGroupIdsList {
				if gvar.New(v).Int() == res.GroupId {
					res.ShowDataBoard = true
					return
				}
			}
		} else {
			// 指定会员
			pmsMemberIds := AppDataBoardViewConfig.CanTestMemberIds
			pmsMemberIdsArr := strings.Split(pmsMemberIds, ",")
			for _, v := range pmsMemberIdsArr {
				if gvar.New(v).Int() == UserInfo.Id {
					res.ShowDataBoard = true
					return
				}
			}
		}
	}

	if g.IsEmpty(res) {
		// 获取用户信息失败,请重新登录。
		err = gerror.NewCode(gcode.CodeNotAuthorized, gi18n.T(ctx, "failed_to_get_member_info"))
		return
	}
	return
}

func (c *ControllerMember) BindEmail(ctx context.Context, req *member.BindEmailReq) (res *member.BindEmailRes, err error) {
	var (
		MemberContextInfo = contexts.GetMemberUser(ctx)
		MemberInfo        *entity.PmsMember
	)
	if err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, MemberContextInfo.Id).Scan(&MemberInfo); err != nil {
		return
	}
	if g.IsEmpty(MemberInfo) {
		// 用户信息不存在
		err = gerror.New(gi18n.T(ctx, "user_info_does_not_exist"))
		return
	}
	// 验证邮箱验证码
	if err = service.BasicsEmsLog().VerifyCode(ctx, &input_basics.VerifyEmsCodeInp{
		Event: "login",
		Email: req.Email,
		Code:  req.Code,
	}); err != nil {
		return
	}

	// 查询是否已存在该邮箱用户
	value, err := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Mail, req.Email).Value(dao.PmsMember.Columns().Id)
	if err != nil {
		return nil, err
	}
	if !value.IsEmpty() {
		// 该邮箱已被占用，无法绑定
		err = gerror.New(gi18n.T(ctx, "email_already_occupied"))
		return
	}
	// 更新邮箱数据到会员表中
	if _, err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, MemberContextInfo.Id).Update(g.MapStrAny{
		dao.PmsMember.Columns().Mail: req.Email,
	}); err != nil {
		return
	}
	return
}

func (c *ControllerMember) BindPhone(ctx context.Context, req *member.BindPhoneReq) (res *member.BindPhoneRes, err error) {
	var (
		MemberContextInfo = contexts.GetMemberUser(ctx)
		MemberInfo        *entity.PmsMember
		MobileNo          string
		smsEvent          string
	)
	// 查询用户是否存在
	if err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, MemberContextInfo.Id).Scan(&MemberInfo); err != nil {
		return
	}
	if g.IsEmpty(MemberInfo) {
		// 用户信息不存在
		err = gerror.New(gi18n.T(ctx, "user_info_does_not_exist"))
		return
	}

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

	// 查询是否存在该手机号
	value, err := dao.PmsMember.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsMember.Columns().Phone:     req.Phone,
		dao.PmsMember.Columns().PhoneArea: req.AreaNo,
	}).Value(dao.PmsMember.Columns().Id)
	if err != nil {
		return nil, err
	}

	savePhone := req.Phone

	// 如果没有找到用户，判断如果是日本手机并且手机号只有10位，则往前补0，再去查找一次
	if value.IsEmpty() && (req.AreaNo == "+81" || req.AreaNo == "81") && len(req.Phone) == 10 {
		savePhone = "0" + req.Phone
		value, err = dao.PmsMember.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsMember.Columns().Phone:     "0" + req.Phone,
			dao.PmsMember.Columns().PhoneArea: req.AreaNo,
		}).Value(dao.PmsMember.Columns().Id)
		if err != nil {
			return nil, err
		}
	}

	// 如果没有找到用户，判断如果是日本手机并且手机号有11位，则往前减0，再去查找一次
	if value.IsEmpty() && (req.AreaNo == "+81" || req.AreaNo == "81") && len(req.Phone) == 11 {
		value, err = dao.PmsMember.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsMember.Columns().Phone:     req.Phone[1:],
			dao.PmsMember.Columns().PhoneArea: req.AreaNo,
		}).Value(dao.PmsMember.Columns().Id)
		if err != nil {
			return nil, err
		}
	}

	// 如果没有找到用户，判断如果是台湾手机并且手机号只有9位，则往前补0，再去查找一次
	if value.IsEmpty() && (req.AreaNo == "+886" || req.AreaNo == "886") && len(req.Phone) == 9 {
		savePhone = "0" + req.Phone
		value, err = dao.PmsMember.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsMember.Columns().Phone:     "0" + req.Phone,
			dao.PmsMember.Columns().PhoneArea: req.AreaNo,
		}).Value(dao.PmsMember.Columns().Id)
		if err != nil {
			return nil, err
		}
	}

	// 如果没有找到用户，判断如果是台湾手机并且手机号有10位，则往前减0，再去查找一次
	if value.IsEmpty() && (req.AreaNo == "+886" || req.AreaNo == "886") && len(req.Phone) == 10 {
		value, err = dao.PmsMember.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsMember.Columns().Phone:     req.Phone[1:],
			dao.PmsMember.Columns().PhoneArea: req.AreaNo,
		}).Value(dao.PmsMember.Columns().Id)
		if err != nil {
			return nil, err
		}
	}

	if !value.IsEmpty() {
		// 该手机号已被占用，无法绑定
		err = gerror.New(gi18n.T(ctx, "phone_already_occupied"))
		return
	}

	if err = service.BasicsSmsLog().VerifyCode(ctx, &input_basics.VerifyCodeInp{
		Event:  smsEvent,
		Mobile: MobileNo,
		Code:   req.Code,
	}); err != nil {
		return
	}

	// 更新会员你手机号信息
	if _, err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, MemberContextInfo.Id).Update(g.MapStrAny{
		dao.PmsMember.Columns().Phone:     savePhone,
		dao.PmsMember.Columns().PhoneArea: req.AreaNo,
	}); err != nil {
		return
	}
	return
}

func (c *ControllerMember) StaffInfo(ctx context.Context, req *member.StaffInfoReq) (res *member.StaffInfoRes, err error) {
	res = new(member.StaffInfoRes)
	if err = dao.PmsStaff.Ctx(ctx).Where(dao.PmsStaff.Columns().Id, req.StaffId).Scan(&res); err != nil {
		return
	}
	return
}
func (c *ControllerMember) ChannelInfo(ctx context.Context, req *member.ChannelInfoReq) (res *member.ChannelInfoRes, err error) {
	res = new(member.ChannelInfoRes)
	if err = dao.PmsChannel.Ctx(ctx).Where(dao.PmsChannel.Columns().Id, req.ChannelId).Scan(&res); err != nil {
		return
	}
	return
}

func (c *ControllerMember) UserLevelRule(ctx context.Context, _ *member.UserLevelRuleReq) (res *member.UserLevelRuleRes, err error) {
	var (
		MemberLevelMap *input_basics.GetConfigModel
	)
	MemberLevelMap = new(input_basics.GetConfigModel)
	if MemberLevelMap, err = service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{
		Group: "memberlevel",
	}); err != nil {
		return
	}
	res = new(member.UserLevelRuleRes)
	res.LevelRule = gvar.New(MemberLevelMap.List["rule_"+contexts.GetLanguage(ctx)]).String()
	return
}

func (c *ControllerMember) EditUserInfo(ctx context.Context, req *member.EditUserInfoReq) (res *member.EditUserInfoRes, err error) {
	var (
		MemberUser   = contexts.GetMemberUser(ctx)
		orm          = g.Model(dao.PmsMember.Table()).Ctx(ctx).Safe()
		updateResult sql.Result
		updateRow    int64
	)
	res = new(member.EditUserInfoRes)
	if g.IsEmpty(req.FullName) && (!g.IsEmpty(req.FirstName) || !g.IsEmpty(req.LastName)) {
		req.FullName = fmt.Sprintf("%s %s", req.LastName, req.FirstName)
	}
	updateData := &entity.PmsMember{
		Avatar:      req.Avatar,
		Sex:         req.Sex,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		FullName:    req.FullName,
		Birthday:    gtime.New(req.Birthday),
		Nationality: req.Nationality,
		Password:    encrypt.Md5([]byte(req.Password)),
		Mail:        req.Email,
		Replenish:   "Y",
	}
	if updateResult, err = orm.Where(dao.PmsMember.Columns().Id, MemberUser.Id).OmitEmptyData().Update(updateData); err != nil {
		return
	}
	if updateRow, err = updateResult.RowsAffected(); err != nil {
		return
	}
	if updateRow == 0 {
		// 用户信息无改动
		err = gerror.New(gi18n.T(ctx, "user_info_remain_unchanged"))
		return
	}
	if updateRow != 1 {
		// 补全用户信息失败
		err = gerror.New(gi18n.T(ctx, "failed_to_complete_user_info"))
		return
	}

	// 发送消息
	// var (
	// 	msgTitle   map[string]string
	// 	msgContent map[string]string
	// )
	// msgTitle = map[string]string{
	// 	"zh":    "系统通知",
	// 	"en":    "System Notification",
	// 	"ja":    "システム通知",
	// 	"ko":    "시스템 알림",
	// 	"zh_CN": "系統通知",
	// }
	// if !g.IsEmpty(req.Birthday) {
	// 	msgContent = map[string]string{
	// 		"zh":    "生日修改成功",
	// 		"en":    "Birthday updated successfully",
	// 		"ja":    "誕生日が正常に更新されました",
	// 		"ko":    "생일이 성공적으로 업데이트되었습니다",
	// 		"zh_CN": "生日修改成功",
	// 	}
	// } else if !g.IsEmpty(req.Sex) {
	// 	msgContent = map[string]string{
	// 		"zh":    "性别修改成功",
	// 		"en":    "Gender updated successfully",
	// 		"ja":    "性別が正常に更新されました",
	// 		"ko":    "성별이 성공적으로 업데이트되었습니다",
	// 		"zh_CN": "性别修改成功",
	// 	}
	// } else if !g.IsEmpty(req.Avatar) {
	// 	msgContent = map[string]string{
	// 		"zh":    "头像修改成功",
	// 		"en":    "Avatar updated successfully",
	// 		"ja":    "アバターが正常に更新されました",
	// 		"ko":    "아바ター가 성공적으로 업데이트되었습니다",
	// 		"zh_CN": "頭像修改成功",
	// 	}
	// } else if !g.IsEmpty(req.FirstName) || !g.IsEmpty(req.LastName) {
	// 	msgContent = map[string]string{
	// 		"zh":    "昵称修改成功",
	// 		"en":    "Nickname updated successfully",
	// 		"ja":    "ニックネームが正常に更新されました",
	// 		"ko":    "닉네임이 성공적으로 업데이트되었습니다",
	// 		"zh_CN": "暱稱修改成功",
	// 	}
	// } else {
	// 	msgContent = map[string]string{
	// 		"zh":    "用户信息修改成功",
	// 		"en":    "User information successfully modified.",
	// 		"ja":    "ユーザー情報が正常に変更されました。",
	// 		"ko":    "사용자 정보가 성공적으로 수정되었습니다.",
	// 		"zh_CN": "使用者資訊修改成功",
	// 	}
	// }
	// service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
	// 	SystemMessageTitle:   msgTitle,
	// 	SystemMessageContent: msgContent,
	// 	Scene:                "system",
	// 	Type:                 "im",
	// 	MemberId:             MemberUser.Id,
	// 	EnablePush:           false,
	// 	EnableSms:            false,
	// 	OperatorId:           int(contexts.GetUserId(ctx)),
	// 	OperatorRole:         "ADMIN",
	// })
	return
}
