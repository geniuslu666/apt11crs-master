package logic_basics

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/token"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/simple"
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

type sBasicsAdminSite struct{}

func NewBasicsAdminSite() *sBasicsAdminSite {
	return &sBasicsAdminSite{}
}

func init() {
	service.RegisterBasicsAdminSite(NewBasicsAdminSite())
}

// Register 账号注册
func (s *sBasicsAdminSite) Register(ctx context.Context, in *input_basics.RegisterInp) (err error) {
	config, err := service.BasicsConfig().GetLogin(ctx)
	if err != nil {
		return
	}

	if config.ForceInvite == 1 && in.InviteCode == "" {
		err = gerror.New("请填写邀请码")
		return
	}

	var data input_basics.MemberAddInp

	// 默认上级
	data.Pid = 1

	// 存在邀请人
	if in.InviteCode != "" {
		pmb, err := service.BasicsAdminMember().GetIdByCode(ctx, &input_basics.GetIdByCodeInp{Code: in.InviteCode})
		if err != nil {
			return err
		}

		if pmb == nil {
			err = gerror.New("邀请人信息不存在")
			return err
		}

		data.Pid = pmb.Id
	}

	if config.RegisterSwitch != 1 {
		err = gerror.New("管理员未开放注册")
		return
	}

	if config.RoleId < 1 {
		err = gerror.New("管理员未配置默认角色")
		return
	}

	if config.DeptId < 1 {
		err = gerror.New("管理员未配置默认部门")
		return
	}

	// 验证唯一性
	err = service.BasicsAdminMember().VerifyUnique(ctx, &input_basics.VerifyUniqueInp{
		Where: g.Map{
			dao.AdminMember.Columns().Username: in.Username,
			dao.AdminMember.Columns().Mobile:   in.Mobile,
		},
	})
	if err != nil {
		return
	}

	// 验证短信验证码
	err = service.BasicsSmsLog().VerifyCode(ctx, &input_basics.VerifyCodeInp{
		Event:  consts.SmsTemplateRegister,
		Mobile: in.Mobile,
		Code:   in.Code,
	})
	if err != nil {
		return
	}

	data.MemberEditInp = &input_basics.MemberEditInp{
		Id:       0,
		RoleId:   config.RoleId,
		PostIds:  config.PostIds,
		DeptId:   config.DeptId,
		Username: in.Username,
		Password: in.Password,
		RealName: "",
		Avatar:   config.Avatar,
		Sex:      3, // 保密
		Mobile:   in.Mobile,
		Status:   consts.StatusEnabled,
	}
	data.Salt = grand.S(6)
	data.InviteCode = grand.S(12)
	data.PasswordHash = gmd5.MustEncryptString(data.Password + data.Salt)
	data.Level, data.Tree, err = service.BasicsAdminMember().GenTree(ctx, data.Pid)
	if err != nil {
		return
	}

	// 提交注册信息
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		id, err := dao.AdminMember.Ctx(ctx).Data(data).OmitEmptyData().InsertAndGetId()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}

		// 更新岗位
		if err = service.BasicsAdminMemberPost().UpdatePostIds(ctx, id, config.PostIds); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
		}
		return
	})
}

// AccountLogin 账号登录
func (s *sBasicsAdminSite) AccountLogin(ctx context.Context, in *input_basics.AccountLoginInp) (res *input_basics.LoginModel, err error) {
	defer func() {
		service.BasicsLoginLog().Push(ctx, &input_basics.LoginLogPushInp{Response: res, Err: err})
	}()

	var mb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Where("username", in.Username).Scan(&mb); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if mb == nil {
		err = gerror.New("账号不存在")
		return
	}

	res = new(input_basics.LoginModel)
	res.Id = mb.Id
	res.Username = mb.Username
	if mb.Salt == "" {
		err = gerror.New("用户信息错误")
		return
	}

	if err = simple.CheckPassword(in.Password, mb.Salt, mb.PasswordHash); err != nil {
		return
	}

	if mb.Status != consts.StatusEnabled {
		err = gerror.New("账号已被禁用")
		return
	}
	// 系统账号登录
	res, err = s.handleLogin(ctx, mb)
	return
}

// MobileLogin 手机号登录
func (s *sBasicsAdminSite) MobileLogin(ctx context.Context, in *input_basics.MobileLoginInp) (res *input_basics.LoginModel, err error) {
	defer func() {
		service.BasicsLoginLog().Push(ctx, &input_basics.LoginLogPushInp{Response: res, Err: err})
	}()

	var mb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Where("mobile ", in.Mobile).Scan(&mb); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if mb == nil {
		err = gerror.New("账号不存在")
		return
	}

	res = new(input_basics.LoginModel)
	res.Id = mb.Id
	res.Username = mb.Username

	err = service.BasicsSmsLog().VerifyCode(ctx, &input_basics.VerifyCodeInp{
		Event:  consts.SmsTemplateLogin,
		Mobile: in.Mobile,
		Code:   in.Code,
	})

	if err != nil {
		return
	}

	if mb.Status != consts.StatusEnabled {
		err = gerror.New("账号已被禁用")
		return
	}

	res, err = s.handleLogin(ctx, mb)
	return
}

// handleLogin .
func (s *sBasicsAdminSite) handleLogin(ctx context.Context, mb *entity.AdminMember) (res *input_basics.LoginModel, err error) {
	role, dept, err := s.getLoginRoleAndDept(ctx, mb.RoleId, mb.DeptId)
	if err != nil {
		return nil, err
	}

	user := &model.Identity{
		Id:          mb.Id,
		Pid:         mb.Pid,
		DeptId:      dept.Id,
		DeptType:    dept.Type,
		RoleId:      role.Id,
		RoleKey:     role.Key,
		Username:    mb.Username,
		RealName:    mb.RealName,
		Avatar:      mb.Avatar,
		Email:       mb.Email,
		Mobile:      mb.Mobile,
		App:         consts.AppAdmin,
		KefuAccount: mb.KefuAccount,
		LoginAt:     gtime.Now(),
	}

	lt, expires, err := token.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	res = &input_basics.LoginModel{
		Username: user.Username,
		Id:       user.Id,
		Token:    lt,
		Expires:  expires,
	}
	return
}

// getLoginRoleAndDept 获取登录的角色和部门信息
func (s *sBasicsAdminSite) getLoginRoleAndDept(ctx context.Context, roleId, deptId int64) (role *entity.AdminRole, dept *entity.AdminDept, err error) {
	if err = dao.AdminRole.Ctx(ctx).Fields("id,key,status").WherePri(roleId).Scan(&role); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if role == nil {
		err = gerror.New("角色不存在或已被删除")
		return
	}

	if role.Status != consts.StatusEnabled {
		err = gerror.New("角色已被禁用，如有疑问请联系管理员")
		return
	}

	if err = dao.AdminDept.Ctx(ctx).Fields("id,type,status").WherePri(deptId).Scan(&dept); err != nil {
		err = gerror.Wrap(err, "获取部门信息失败，请稍后重试！")
		return
	}

	if dept == nil {
		err = gerror.New("部门不存在或已被删除")
		return
	}

	if dept.Status != consts.StatusEnabled {
		err = gerror.New("部门已被禁用，如有疑问请联系管理员")
		return
	}
	return
}

// BindUserContext 绑定用户上下文
func (s *sBasicsAdminSite) BindUserContext(ctx context.Context, claims *model.Identity) (err error) {
	//// 如果不想每次访问都重新加载用户信息，可以放开注释。但在本次登录未失效前，用户信息不会刷新
	//contexts.SetUser(ctx, claims)
	//return

	var mb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).WherePri(claims.Id).Scan(&mb); err != nil {
		err = gerror.Wrap(err, "获取用户信息失败，请稍后重试！")
		return
	}

	if mb == nil {
		err = gerror.Wrap(err, "账号不存在或已被删除！")
		return
	}

	if mb.Status != consts.StatusEnabled {
		err = gerror.New("账号已被禁用，如有疑问请联系管理员")
		return
	}

	role, dept, err := s.getLoginRoleAndDept(ctx, mb.RoleId, mb.DeptId)
	if err != nil {
		return err
	}

	user := &model.Identity{
		Id:          mb.Id,
		Pid:         mb.Pid,
		DeptId:      dept.Id,
		DeptType:    dept.Type,
		RoleId:      mb.RoleId,
		RoleKey:     role.Key,
		Username:    mb.Username,
		RealName:    mb.RealName,
		Avatar:      mb.Avatar,
		Email:       mb.Email,
		Mobile:      mb.Mobile,
		App:         claims.App,
		KefuAccount: mb.KefuAccount,
		LoginAt:     claims.LoginAt,
	}

	contexts.SetUser(ctx, user)
	return
}
