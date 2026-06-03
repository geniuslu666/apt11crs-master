package logic_basics

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/handler"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/tree"
	"APT/utility/validate"
	"context"
	"sync"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"
)

// SuperAdmin 超级管理员用户
type SuperAdmin struct {
	sync.RWMutex
	RoleId    int64              // 超管角色ID
	MemberIds map[int64]struct{} // 超管用户ID
}

type sBasicsAdminMember struct {
	superAdmin *SuperAdmin
}

func NewBasicsAdminMember() *sBasicsAdminMember {
	return &sBasicsAdminMember{
		superAdmin: new(SuperAdmin),
	}
}

func init() {
	service.RegisterBasicsAdminMember(NewBasicsAdminMember())
}

// UpdateEmail 换绑邮箱
func (s *sBasicsAdminMember) UpdateEmail(ctx context.Context, in *input_basics.MemberUpdateEmailInp) (err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		err = gerror.New("获取用户信息失败！")
		return
	}
	memberId := user.Id
	if memberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return
	}

	var mb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).WherePri(memberId).Scan(&mb); err != nil {
		err = gerror.Wrap(err, "获取用户信息失败，请稍后重试！")
		return
	}

	if mb == nil {
		err = gerror.New("用户信息不存在")
		return
	}

	if mb.Email == in.Email {
		err = gerror.New("新旧邮箱不能一样")
		return
	}

	if !validate.IsEmail(in.Email) {
		err = gerror.New("邮箱地址不正确")
		return
	}

	// 存在原绑定号码，需要进行验证
	if mb.Email != "" {
		err = service.BasicsEmsLog().VerifyCode(ctx, &input_basics.VerifyEmsCodeInp{
			Event: consts.EmsTemplateBind,
			Email: mb.Email,
			Code:  in.Code,
		})
		if err != nil {
			return
		}
	}

	update := g.Map{
		dao.AdminMember.Columns().Email: in.Email,
	}

	if _, err = dao.AdminMember.Ctx(ctx).WherePri(memberId).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, "换绑邮箱失败，请稍后重试！")
		return
	}
	return
}

// UpdateMobile 换绑手机号
func (s *sBasicsAdminMember) UpdateMobile(ctx context.Context, in *input_basics.MemberUpdateMobileInp) (err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		err = gerror.New("获取用户信息失败！")
		return
	}
	memberId := user.Id
	if memberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return
	}

	var mb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).WherePri(memberId).Scan(&mb); err != nil {
		err = gerror.Wrap(err, "获取用户信息失败，请稍后重试！")
		return
	}

	if mb == nil {
		err = gerror.New("用户信息不存在")
		return
	}

	if mb.Mobile == in.Mobile {
		err = gerror.New("新旧手机号不能一样")
		return
	}

	if !validate.IsMobile(in.Mobile) {
		err = gerror.New("手机号码不正确")
		return
	}

	// 存在原绑定号码，需要进行验证
	if mb.Mobile != "" {
		err = service.BasicsSmsLog().VerifyCode(ctx, &input_basics.VerifyCodeInp{
			Event:  consts.SmsTemplateBind,
			Mobile: mb.Mobile,
			Code:   in.Code,
		})
		if err != nil {
			return
		}
	}

	update := g.Map{
		dao.AdminMember.Columns().Mobile: in.Mobile,
	}

	if _, err = dao.AdminMember.Ctx(ctx).WherePri(memberId).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, "换绑手机号失败，请稍后重试！")
		return
	}
	return
}

// UpdateProfile 更新用户资料
func (s *sBasicsAdminMember) UpdateProfile(ctx context.Context, in *input_basics.MemberUpdateProfileInp) (err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		err = gerror.New("获取用户信息失败！")
		return
	}
	memberId := user.Id
	if memberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return
	}

	var mb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).WherePri(memberId).Scan(&mb); err != nil {
		err = gerror.Wrap(err, "获取用户信息失败，请稍后重试！")
		return
	}

	if mb == nil {
		err = gerror.New("用户信息不存在")
		return
	}

	cols := dao.AdminMember.Columns()
	update := g.Map{
		cols.Avatar:   in.Avatar,
		cols.RealName: in.RealName,
		cols.Qq:       in.Qq,
		cols.Birthday: in.Birthday,
		cols.Sex:      in.Sex,
		cols.CityId:   in.CityId,
		cols.Address:  in.Address,
	}

	if _, err = dao.AdminMember.Ctx(ctx).WherePri(memberId).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, "更新资料失败，请稍后重试！")
		return
	}
	return
}

// UpdatePwd 修改登录密码
func (s *sBasicsAdminMember) UpdatePwd(ctx context.Context, in *input_basics.MemberUpdatePwdInp) (err error) {
	var mb entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).WherePri(in.Id).Scan(&mb); err != nil {
		err = gerror.Wrap(err, "获取用户信息失败，请稍后重试！")
		return
	}

	if gmd5.MustEncryptString(in.OldPassword+mb.Salt) != mb.PasswordHash {
		err = gerror.New("原密码不正确")
		return
	}

	update := g.Map{
		dao.AdminMember.Columns().PasswordHash: gmd5.MustEncryptString(in.NewPassword + mb.Salt),
	}

	if _, err = dao.AdminMember.Ctx(ctx).WherePri(in.Id).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, "更新登录密码失败，请稍后重试！")
		return
	}
	return
}

// ResetPwd 重置密码
func (s *sBasicsAdminMember) ResetPwd(ctx context.Context, in *input_basics.MemberResetPwdInp) (err error) {
	var (
		mb       *entity.AdminMember
		memberId = contexts.GetUserId(ctx)
	)

	if err = s.FilterAuthModel(ctx, memberId).WherePri(in.Id).Scan(&mb); err != nil {
		err = gerror.Wrap(err, "获取用户信息失败，请稍后重试！")
		return
	}

	if mb == nil {
		err = gerror.New("用户信息不存在")
		return
	}

	update := g.Map{
		dao.AdminMember.Columns().PasswordHash: gmd5.MustEncryptString(in.Password + mb.Salt),
	}

	if _, err = s.FilterAuthModel(ctx, memberId).WherePri(in.Id).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, "用户密码重置失败，请稍后重试！")
		return
	}
	return
}

// VerifyUnique 验证管理员唯一属性
func (s *sBasicsAdminMember) VerifyUnique(ctx context.Context, in *input_basics.VerifyUniqueInp) (err error) {
	if in.Where == nil {
		return
	}

	cols := dao.AdminMember.Columns()
	msgMap := g.MapStrStr{
		cols.Username:   "用户名已存在，请换一个",
		cols.Email:      "邮箱已存在，请换一个",
		cols.Mobile:     "手机号已存在，请换一个",
		cols.InviteCode: "邀请码已存在，请换一个",
	}

	for k, v := range in.Where {
		if v == "" {
			continue
		}
		message, ok := msgMap[k]
		if !ok {
			err = gerror.Newf("字段 [ %v ] 未配置唯一属性验证", k)
			return
		}
		if err = hgorm.IsUnique(ctx, &dao.AdminMember, g.Map{k: v}, message, in.Id); err != nil {
			return
		}
	}
	return
}

// Delete 删除用户
func (s *sBasicsAdminMember) Delete(ctx context.Context, in *input_basics.MemberDeleteInp) (err error) {
	memberId := contexts.GetUserId(ctx)
	if memberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return
	}

	var list []*entity.AdminMember
	if err = s.FilterAuthModel(ctx, memberId).WherePri(in.Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取用户信息失败，请稍后重试！")
		return
	}

	if len(list) == 0 {
		err = gerror.New("需要删除的用户不存在或已删除！")
		return
	}

	for _, v := range list {
		if s.VerifySuperId(ctx, v.Id) {
			err = gerror.New("超管账号禁止删除！")
			return
		}
		count, err := dao.AdminMember.Ctx(ctx).Where("pid", v.Id).Count()
		if err != nil {
			err = gerror.Wrap(err, "删除用户检查失败，请稍后重试！")
			return err
		}
		if count > 0 {
			err = gerror.Newf("用户[%v]存在下级，请先删除TA的下级用户！", v.Id)
			return err
		}
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = s.FilterAuthModel(ctx, memberId).WherePri(in.Id).Delete(); err != nil {
			err = gerror.Wrap(err, "删除用户失败，请稍后重试！")
			return
		}

		if _, err = dao.AdminMemberPost.Ctx(ctx).Where("member_id", in.Id).Delete(); err != nil {
			err = gerror.Wrap(err, "删除用户岗位失败，请稍后重试！")
		}

		// 这里如果需要，可以加入更多删除用户的相关处理
		// ...
		return
	})
}

// Edit 修改/新增用户
func (s *sBasicsAdminMember) Edit(ctx context.Context, in *input_basics.MemberEditInp) (err error) {
	opMemberId := contexts.GetUserId(ctx)
	if opMemberId <= 0 {
		err = gerror.New("获取用户信息失败！")
		return
	}

	if in.Username == "" {
		err = gerror.New("帐号不能为空")
		return
	}

	cols := dao.AdminMember.Columns()
	err = s.VerifyUnique(ctx, &input_basics.VerifyUniqueInp{
		Id: in.Id,
		Where: g.Map{
			cols.Username: in.Username,
			cols.Mobile:   in.Mobile,
			cols.Email:    in.Email,
		},
	})
	if err != nil {
		return
	}

	// 验证角色ID
	if err = service.BasicsAdminRole().VerifyRoleId(ctx, in.RoleId); err != nil {
		return
	}

	// 验证部门ID
	if err = service.BasicsAdminDept().VerifyDeptId(ctx, in.DeptId); err != nil {
		return
	}

	config, err := service.BasicsConfig().GetLogin(ctx)
	if err != nil {
		return
	}

	needLoadSuperAdmin := false
	defer func() {
		if needLoadSuperAdmin {
			// 本地先更新
			s.LoadSuperAdmin(ctx)
		}
	}()

	// 修改
	if in.Id > 0 {
		if s.VerifySuperId(ctx, in.Id) {
			err = gerror.New("超管账号禁止编辑！")
			return
		}

		mod := s.FilterAuthModel(ctx, opMemberId)

		if in.Password != "" {
			// 修改密码，需要获取到密码盐
			salt, err := s.FilterAuthModel(ctx, opMemberId).Fields(cols.Salt).WherePri(in.Id).Value()
			if err != nil {
				err = gerror.Wrap(err, "获取用户信息失败，请稍后重试！")
				return err
			}
			if salt.IsEmpty() {
				err = gerror.New("该用户没有设置密码盐，请联系管理员！")
				return err
			}
			in.PasswordHash = gmd5.MustEncryptString(in.Password + salt.String())
		} else {
			mod = mod.FieldsEx(cols.PasswordHash)
		}

		return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			if _, err = mod.WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改用户信息失败，请稍后重试！")
				return
			}

			// 更新岗位
			if err = service.BasicsAdminMemberPost().UpdatePostIds(ctx, in.Id, in.PostIds); err != nil {
				err = gerror.Wrap(err, "更新用户岗位失败，请稍后重试！")
			}

			// 编辑客服账号(接待模式、接待固定值)
			if _, err = g.DB("chat").Model("user").Where(dao.User.Columns().Name, in.KefuAccount).Data(g.Map{
				dao.User.Columns().Reception:      in.Reception,
				dao.User.Columns().ReceptionValue: in.ReceptionValue,
			}).OmitEmptyData().Update(); err != nil {
				return
			}

			needLoadSuperAdmin = in.RoleId == s.superAdmin.RoleId
			return
		})
	}

	// 新增用户时的额外属性
	var data input_basics.MemberAddInp
	data.MemberEditInp = in
	data.KefuAccount = in.Username
	data.Salt = grand.S(6)
	data.InviteCode = grand.S(12)
	data.PasswordHash = gmd5.MustEncryptString(data.Password + data.Salt)

	// 关系树
	data.Pid = opMemberId
	data.Level, data.Tree, err = s.GenTree(ctx, opMemberId)
	if err != nil {
		return
	}

	// 默认头像
	if data.Avatar == "" {
		data.Avatar = config.Avatar
	}
	var kefu_user_id int64
	// 创建客服账号
	if kefu_user_id, err = g.DB("chat").Model("user").Data(&entity.User{
		Pid:            4,
		Name:           in.Username,
		Password:       gmd5.MustEncrypt(in.Password),
		Nickname:       in.RealName,
		CreatedAt:      gtime.Now(),
		ExpiredAt:      gtime.New("2038-01-01"),
		UpdatedAt:      gtime.Now(),
		Avator:         "/static/images/0.jpg",
		AgentNum:       1000,
		Email:          in.Email,
		Reception:      in.Reception,
		ReceptionValue: in.ReceptionValue,
		Uuid:           "231d415be776fbc6e84a9c38a7b312ac",
	}).OmitEmptyData().InsertAndGetId(); err != nil {
		return
	}

	// 创建客服账号对应的权限
	if _, err = g.DB("chat").Model("user_role").Data(&entity.UserRole{
		UserId: gvar.New(kefu_user_id).Int(),
		RoleId: 3,
	}).InsertAndGetId(); err != nil {
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		id, err := dao.AdminMember.Ctx(ctx).Data(data).OmitEmptyData().InsertAndGetId()
		if err != nil {
			err = gerror.Wrap(err, "新增用户失败，请稍后重试！")
			return
		}

		// 更新岗位
		if err = service.BasicsAdminMemberPost().UpdatePostIds(ctx, id, in.PostIds); err != nil {
			err = gerror.Wrap(err, "新增用户岗位失败，请稍后重试！")
		}

		needLoadSuperAdmin = in.RoleId == s.superAdmin.RoleId
		return
	})
}

// View 获取用户信息
func (s *sBasicsAdminMember) View(ctx context.Context, in *input_basics.MemberViewInp) (res *input_basics.MemberViewModel, err error) {
	if err = s.FilterAuthModel(ctx, contexts.GetUserId(ctx)).Hook(hook.MemberInfo).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取用户信息失败，请稍后重试！")
	}
	return
}

// List 获取用户列表
func (s *sBasicsAdminMember) List(ctx context.Context, in *input_basics.MemberListInp) (list []*input_basics.MemberListModel, totalCount int, err error) {
	mod := s.FilterAuthModel(ctx, contexts.GetUserId(ctx)).WithAll()
	cols := dao.AdminMember.Columns()

	if in.RealName != "" {
		mod = mod.WhereLike(cols.RealName, "%"+in.RealName+"%")
	}

	if in.Username != "" {
		mod = mod.WhereLike(cols.Username, "%"+in.Username+"%")
	}

	if in.Mobile > 0 {
		mod = mod.Where(cols.Mobile, in.Mobile)
	}

	if in.Status > 0 {
		mod = mod.Where(cols.Status, in.Status)
	}

	if in.DeptId > 0 {
		mod = mod.Where(cols.DeptId, in.DeptId)
	}

	if in.RoleId > 0 {
		mod = mod.Where(cols.RoleId, in.RoleId)
	}

	if in.Id > 0 {
		mod = mod.Where(cols.Id, in.Id)
	}

	if in.Pid > 0 {
		mod = mod.Where(cols.Pid, in.Pid)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(cols.CreatedAt, gtime.New(in.CreatedAt[0]), gtime.New(in.CreatedAt[1]))
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取用户数据行失败！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Hook(hook.MemberInfo).Page(in.Page, in.PerPage).OrderDesc(cols.Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取用户列表失败！")
		return
	}

	for _, v := range list {
		columns, err := dao.AdminMemberPost.Ctx(ctx).Fields(dao.AdminMemberPost.Columns().PostId).Where(dao.AdminMemberPost.Columns().MemberId, v.Id).Array()
		if err != nil {
			err = gerror.Wrap(err, "获取用户岗位数据失败！")
			return nil, 0, err
		}
		v.PostIds = g.NewVar(columns).Int64s()

		if !g.IsEmpty(v.KefuAccount) {
			var (
				User *entity.User
			)
			err = g.DB("chat").Model("user").Where(dao.User.Columns().Name, v.KefuAccount).Scan(&User)
			if err != nil {
				return nil, 0, err
			}
			if !g.IsEmpty(User) {
				v.Reception = User.Reception
				v.ReceptionValue = User.ReceptionValue
			}
		}

	}
	return
}

// Status 更新状态
func (s *sBasicsAdminMember) Status(ctx context.Context, in *input_basics.MemberStatusInp) (err error) {
	if s.VerifySuperId(ctx, in.Id) {
		err = gerror.New("超管账号不能更改状态")
		return
	}

	if _, err = s.FilterAuthModel(ctx, contexts.GetUserId(ctx)).WherePri(in.Id).Data(dao.AdminMember.Columns().Status, in.Status).Update(); err != nil {
		err = gerror.Wrap(err, "更新用户状态失败，请稍后重试！")
	}
	return
}

// GenTree 生成关系树
func (s *sBasicsAdminMember) GenTree(ctx context.Context, pid int64) (level int, newTree string, err error) {
	var pmb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).WherePri(pid).Scan(&pmb); err != nil {
		return
	}

	if pmb == nil {
		err = gerror.New("上级信息不存在")
		return
	}

	level = pmb.Level + 1
	newTree = tree.GenLabel(pmb.Tree, pmb.Id)
	return
}

// LoginMemberInfo 获取登录用户信息
func (s *sBasicsAdminMember) LoginMemberInfo(ctx context.Context) (res *input_basics.LoginMemberInfoModel, err error) {
	var memberId = contexts.GetUserId(ctx)
	if memberId <= 0 {
		err = gerror.New("用户身份异常，请重新登录！")
		return
	}

	if err = dao.AdminMember.Ctx(ctx).Hook(hook.MemberInfo).WherePri(memberId).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取用户信息失败，请稍后重试！")
		return
	}

	if res == nil {
		err = gerror.New("用户不存在！")
		return
	}

	// 细粒度权限
	permissions, err := service.BasicsAdminMenu().LoginPermissions(ctx, memberId)
	if err != nil {
		return
	}
	res.Permissions = permissions

	// 登录统计
	stat, err := s.MemberLoginStat(ctx, &input_basics.MemberLoginStatInp{MemberId: memberId})
	if err != nil {
		return
	}

	res.MemberLoginStatModel = stat
	res.Mobile = gstr.HideStr(res.Mobile, 40, `*`)
	res.Email = gstr.HideStr(res.Email, 40, `*`)
	//res.OpenId, _ = service.CommonWechat().GetOpenId(ctx)
	res.DeptType = contexts.GetDeptType(ctx)
	return
}

// MemberLoginStat 用户登录统计
func (s *sBasicsAdminMember) MemberLoginStat(ctx context.Context, in *input_basics.MemberLoginStatInp) (res *input_basics.MemberLoginStatModel, err error) {
	var (
		models *entity.SysLoginLog
		cols   = dao.SysLoginLog.Columns()
	)

	err = dao.SysLoginLog.Ctx(ctx).Fields(cols.LoginAt, cols.LoginIp).
		Where(cols.MemberId, in.MemberId).
		Where(cols.Status, consts.StatusEnabled).
		OrderDesc(cols.Id).
		Scan(&models)

	if err != nil {
		return
	}

	res = new(input_basics.MemberLoginStatModel)
	if models == nil {
		return
	}

	res.LastLoginAt = models.LoginAt
	res.LastLoginIp = models.LoginIp
	res.LoginCount, err = dao.SysLoginLog.Ctx(ctx).
		Where(cols.MemberId, in.MemberId).
		Where(cols.Status, consts.StatusEnabled).
		Count()
	return
}

// GetIdByCode 通过邀请码获取用户ID
func (s *sBasicsAdminMember) GetIdByCode(ctx context.Context, in *input_basics.GetIdByCodeInp) (res *input_basics.GetIdByCodeModel, err error) {
	if err = dao.AdminMember.Ctx(ctx).Fields(input_basics.GetIdByCodeModel{}).Where(dao.AdminMember.Columns().InviteCode, in.Code).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取用户信息失败，请稍后重试！")
	}
	return
}

// Select 获取可选的用户选项
func (s *sBasicsAdminMember) Select(ctx context.Context, in *input_basics.MemberSelectInp) (res []*input_basics.MemberSelectModel, err error) {
	err = dao.AdminMember.Ctx(ctx).Fields("id as value,real_name as label,username,avatar").
		Handler(handler.FilterAuthWithField("id")).
		Scan(&res)
	if err != nil {
		err = gerror.Wrap(err, "获取可选用户选项失败，请稍后重试！")
	}
	return
}

// GetIdsByKeyword 根据关键词查找符合条件的用户ID
func (s *sBasicsAdminMember) GetIdsByKeyword(ctx context.Context, ks string) (res []int64, err error) {
	ks = gstr.Trim(ks)
	if len(ks) == 0 {
		return
	}
	array, err := dao.AdminMember.Ctx(ctx).Fields("id").
		Where("`id` = ? or `real_name` = ? or `username` = ? or `mobile` = ?", ks, ks, ks, ks).
		Array()
	if err != nil {
		err = gerror.Wrap(err, "根据关键词获取用户ID失败，请稍后重试！")
	}
	res = gvar.New(array).Int64s()
	return
}

// VerifySuperId 验证是否为超管
func (s *sBasicsAdminMember) VerifySuperId(ctx context.Context, verifyId int64) bool {
	s.superAdmin.RLock()
	defer s.superAdmin.RUnlock()

	if s.superAdmin == nil || s.superAdmin.MemberIds == nil {
		g.Log().Error(ctx, "superAdmin is not initialized.")
		return false
	}

	_, ok := s.superAdmin.MemberIds[verifyId]
	return ok
}

// LoadSuperAdmin 加载超管数据
func (s *sBasicsAdminMember) LoadSuperAdmin(ctx context.Context) {
	value, err := dao.AdminRole.Ctx(ctx).Where(dao.AdminRole.Columns().Key, consts.SuperRoleKey).Value()
	if err != nil {
		g.Log().Errorf(ctx, "LoadSuperAdmin AdminRole err:%+v", err)
		return
	}

	if value.IsEmpty() || value.IsNil() {
		g.Log().Error(ctx, "the superAdmin role must be configured.")
		return
	}

	array, err := dao.AdminMember.Ctx(ctx).Fields(dao.AdminMember.Columns().Id).Where(dao.AdminMember.Columns().RoleId, value).Array()
	if err != nil {
		g.Log().Errorf(ctx, "LoadSuperAdmin AdminMember err:%+v", err)
		return
	}

	s.superAdmin.Lock()
	defer s.superAdmin.Unlock()

	s.superAdmin.MemberIds = make(map[int64]struct{}, len(array))
	for _, v := range array {
		s.superAdmin.MemberIds[v.Int64()] = struct{}{}
	}
	s.superAdmin.RoleId = value.Int64()
}

// ClusterSyncSuperAdmin 集群同步
func (s *sBasicsAdminMember) ClusterSyncSuperAdmin(ctx context.Context, message *gredis.Message) {
	s.LoadSuperAdmin(ctx)
}

// FilterAuthModel 过滤用户操作权限
// 非超管用户只能操作自己的下级角色用户，并且需要满足自身角色的数据权限设置
func (s *sBasicsAdminMember) FilterAuthModel(ctx context.Context, memberId int64) *gdb.Model {
	m := dao.AdminMember.Ctx(ctx)
	if s.VerifySuperId(ctx, memberId) {
		return m
	}

	var roleId int64
	if contexts.GetUserId(ctx) == memberId {
		// 当前登录用户直接从上下文中取角色ID
		roleId = contexts.GetRoleId(ctx)
	} else {
		ro, err := dao.AdminMember.Ctx(ctx).Fields("role_id").Where("id", memberId).Value()
		if err != nil {
			g.Log().Panicf(ctx, "failed to get role information, err:%+v", err)
			return nil
		}
		roleId = ro.Int64()
	}

	roleIds, err := service.BasicsAdminRole().GetSubRoleIds(ctx, roleId, false)
	if err != nil {
		g.Log().Panicf(ctx, "get the subordinate role permission exception, err:%+v", err)
		return nil
	}
	return m.Where("id <> ?", memberId).WhereIn("role_id", roleIds).Handler(handler.FilterAuthWithField("id"))
}
