// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"APT/api/admin/basics"
	"APT/internal/library/hgorm/handler"
	"APT/internal/library/ws"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"APT/utility/tree"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IBasicsIndexBanner interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_basics.IndexBannerListInp) (list []*input_basics.IndexBannerListModel, totalCount int, err error)
		Edit(ctx context.Context, in *input_basics.IndexBannerEditInp) (err error)
		Delete(ctx context.Context, in *input_basics.IndexBannerDeleteInp) (err error)
		// MaxSort 获取Banner最大排序
		MaxSort(ctx context.Context, in *input_basics.IndexBannerMaxSortInp) (res *input_basics.IndexBannerMaxSortModel, err error)
		View(ctx context.Context, in *input_basics.IndexBannerViewInp) (res *input_basics.IndexBannerViewModel, err error)
		// Status 更新Banner状态
		Status(ctx context.Context, in *input_basics.IndexBannerStatusInp) (err error)
		// BannerSort 编辑排序
		BannerSort(ctx context.Context, in *input_basics.IndexBannerSortInp) (err error)
	}
	IBasicsIndexNav interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_basics.PmsIndexNavListInp) (list []*input_basics.PmsIndexNavListModel, totalCount int, err error)
		ApiNavAll(ctx context.Context, in *input_basics.PmsIndexNavAllInp) (list []*input_basics.PmsIndexNavApiListModel, err error)
		All(ctx context.Context, in *input_basics.PmsIndexNavAllInp) (list []*input_basics.PmsIndexNavAllModel, err error)
		Edit(ctx context.Context, in *input_basics.PmsIndexNavEditInp) (err error)
		Delete(ctx context.Context, in *input_basics.PmsIndexNavDeleteInp) (err error)
		View(ctx context.Context, in *input_basics.PmsIndexNavViewInp) (res *input_basics.PmsIndexNavViewModel, err error)
		Switch(ctx context.Context, in *input_basics.PmsIndexNavSwitchInp) (err error)
		// MinappStatus 更新小程序显示状态
		MinappStatus(ctx context.Context, in *input_basics.PmsIndexNavSwitchInp) (err error)
		// NavSort 编辑排序
		NavSort(ctx context.Context, in *input_basics.PmsIndexNavSortInp) (err error)
	}
	IBasicsTestNav interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_basics.PmsTestNavListInp) (list []*input_basics.PmsTestNavListModel, totalCount int, err error)
		ApiNavAll(ctx context.Context, in *input_basics.PmsTestNavAllInp) (list []*input_basics.PmsTestNavApiListModel, err error)
		All(ctx context.Context, in *input_basics.PmsTestNavAllInp) (list []*input_basics.PmsTestNavAllModel, err error)
		Edit(ctx context.Context, in *input_basics.PmsTestNavEditInp) (err error)
		Delete(ctx context.Context, in *input_basics.PmsTestNavDeleteInp) (err error)
		View(ctx context.Context, in *input_basics.PmsTestNavViewInp) (res *input_basics.PmsTestNavViewModel, err error)
		Switch(ctx context.Context, in *input_basics.PmsTestNavSwitchInp) (err error)
		// NavSort 编辑排序
		NavSort(ctx context.Context, in *input_basics.PmsTestNavSortInp) (err error)
	}
	IBasicsAdminMember interface {
		// UpdateEmail 换绑邮箱
		UpdateEmail(ctx context.Context, in *input_basics.MemberUpdateEmailInp) (err error)
		// UpdateMobile 换绑手机号
		UpdateMobile(ctx context.Context, in *input_basics.MemberUpdateMobileInp) (err error)
		// UpdateProfile 更新用户资料
		UpdateProfile(ctx context.Context, in *input_basics.MemberUpdateProfileInp) (err error)
		// UpdatePwd 修改登录密码
		UpdatePwd(ctx context.Context, in *input_basics.MemberUpdatePwdInp) (err error)
		// ResetPwd 重置密码
		ResetPwd(ctx context.Context, in *input_basics.MemberResetPwdInp) (err error)
		// VerifyUnique 验证管理员唯一属性
		VerifyUnique(ctx context.Context, in *input_basics.VerifyUniqueInp) (err error)
		// Delete 删除用户
		Delete(ctx context.Context, in *input_basics.MemberDeleteInp) (err error)
		// Edit 修改/新增用户
		Edit(ctx context.Context, in *input_basics.MemberEditInp) (err error)
		// View 获取用户信息
		View(ctx context.Context, in *input_basics.MemberViewInp) (res *input_basics.MemberViewModel, err error)
		// List 获取用户列表
		List(ctx context.Context, in *input_basics.MemberListInp) (list []*input_basics.MemberListModel, totalCount int, err error)
		// Status 更新状态
		Status(ctx context.Context, in *input_basics.MemberStatusInp) (err error)
		// GenTree 生成关系树
		GenTree(ctx context.Context, pid int64) (level int, newTree string, err error)
		// LoginMemberInfo 获取登录用户信息
		LoginMemberInfo(ctx context.Context) (res *input_basics.LoginMemberInfoModel, err error)
		// MemberLoginStat 用户登录统计
		MemberLoginStat(ctx context.Context, in *input_basics.MemberLoginStatInp) (res *input_basics.MemberLoginStatModel, err error)
		// GetIdByCode 通过邀请码获取用户ID
		GetIdByCode(ctx context.Context, in *input_basics.GetIdByCodeInp) (res *input_basics.GetIdByCodeModel, err error)
		// Select 获取可选的用户选项
		Select(ctx context.Context, in *input_basics.MemberSelectInp) (res []*input_basics.MemberSelectModel, err error)
		// GetIdsByKeyword 根据关键词查找符合条件的用户ID
		GetIdsByKeyword(ctx context.Context, ks string) (res []int64, err error)
		// VerifySuperId 验证是否为超管
		VerifySuperId(ctx context.Context, verifyId int64) bool
		// LoadSuperAdmin 加载超管数据
		LoadSuperAdmin(ctx context.Context)
		// ClusterSyncSuperAdmin 集群同步
		ClusterSyncSuperAdmin(ctx context.Context, message *gredis.Message)
		// FilterAuthModel 过滤用户操作权限
		// 非超管用户只能操作自己的下级角色用户，并且需要满足自身角色的数据权限设置
		FilterAuthModel(ctx context.Context, memberId int64) *gdb.Model
	}
	IBasicsAdminMemberPost interface {
		// UpdatePostIds 更新用户岗位
		UpdatePostIds(ctx context.Context, memberId int64, postIds []int64) (err error)
	}
	IBasicsAdminMenu interface {
		// Model Orm模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// Delete 删除
		Delete(ctx context.Context, in *input_basics.MenuDeleteInp) (err error)
		// VerifyUnique 验证菜单唯一属性
		VerifyUnique(ctx context.Context, in *input_basics.VerifyUniqueInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *input_basics.MenuEditInp) (err error)
		// List 获取菜单列表
		List(ctx context.Context, in *input_basics.MenuListInp) (res *input_basics.MenuListModel, err error)
		// GetMenuList 获取菜单列表
		GetMenuList(ctx context.Context, memberId int64) (res *basics.RoleDynamicRes, err error)
		// LoginPermissions 获取登录成功后的细粒度权限
		LoginPermissions(ctx context.Context, memberId int64) (lists input_basics.MemberLoginPermissions, err error)
		// GetFastList 获取菜单列表
		GetFastList(ctx context.Context) (res map[int64]*entity.AdminMenu, err error)
	}
	IBasicsAdminMonitor interface {
		// StartMonitor 启动服务监控
		StartMonitor(ctx context.Context)
		// GetMeta 获取监控元数据
		GetMeta(ctx context.Context) *model.MonitorData
	}
	IBasicsAdminNotice interface {
		// Model Orm模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// Delete 删除
		Delete(ctx context.Context, in *input_basics.NoticeDeleteInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *input_basics.NoticeEditInp) (err error)
		// Status 更新部门状态
		Status(ctx context.Context, in *input_basics.NoticeStatusInp) (err error)
		// MaxSort 最大排序
		MaxSort(ctx context.Context, in *input_basics.NoticeMaxSortInp) (res *input_basics.NoticeMaxSortModel, err error)
		// View 获取指定字典类型信息
		View(ctx context.Context, in *input_basics.NoticeViewInp) (res *input_basics.NoticeViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *input_basics.NoticeListInp) (list []*input_basics.NoticeListModel, totalCount int, err error)
		// PullMessages 拉取未读消息列表
		PullMessages(ctx context.Context, in *input_basics.PullMessagesInp) (res *input_basics.PullMessagesModel, err error)
		// UnreadCount 获取所有类型消息的未读数量
		UnreadCount(ctx context.Context, in *input_basics.NoticeUnreadCountInp) (res *input_basics.NoticeUnreadCountModel, err error)
		// UpRead 更新已读
		UpRead(ctx context.Context, in *input_basics.NoticeUpReadInp) (err error)
		// ReadAll 已读全部
		ReadAll(ctx context.Context, in *input_basics.NoticeReadAllInp) (err error)
		// MessageList 我的消息列表
		MessageList(ctx context.Context, in *input_basics.NoticeMessageListInp) (list []*input_basics.NoticeMessageListModel, totalCount int, err error)
	}
	IBasicsAdminPost interface {
		// Delete 删除
		Delete(ctx context.Context, in *input_basics.PostDeleteInp) (err error)
		// VerifyUnique 验证部门唯一属性
		VerifyUnique(ctx context.Context, in *input_basics.VerifyUniqueInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *input_basics.PostEditInp) (err error)
		// MaxSort 最大排序
		MaxSort(ctx context.Context, in *input_basics.PostMaxSortInp) (res *input_basics.PostMaxSortModel, err error)
		// View 获取指定岗位信息
		View(ctx context.Context, in *input_basics.PostViewInp) (res *input_basics.PostViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *input_basics.PostListInp) (list []*input_basics.PostListModel, totalCount int, err error)
		// Option 岗位选项
		Option(ctx context.Context) (opts []*model.Option, err error)
		// GetMemberByStartName 获取指定用户的第一岗位
		GetMemberByStartName(ctx context.Context, memberId int64) (name string, err error)
		// Status 更新状态
		Status(ctx context.Context, in *input_basics.PostStatusInp) (err error)
	}
	IBasicsAdminRole interface {
		// Verify 验证权限
		Verify(ctx context.Context, path string, method string) bool
		// List 获取列表
		List(ctx context.Context, in *input_basics.RoleListInp) (res *input_basics.RoleListModel, totalCount int, err error)
		// GetName 获取指定角色的名称
		GetName(ctx context.Context, id int64) (name string, err error)
		// GetMemberList 获取指定用户的岗位列表
		GetMemberList(ctx context.Context, id int64) (list []*input_basics.RoleListModel, err error)
		// GetPermissions 更改角色菜单权限
		GetPermissions(ctx context.Context, in *input_basics.GetPermissionsInp) (res *input_basics.GetPermissionsModel, err error)
		// UpdatePermissions 更改角色菜单权限
		UpdatePermissions(ctx context.Context, in *input_basics.UpdatePermissionsInp) (err error)
		Edit(ctx context.Context, in *input_basics.RoleEditInp) (err error)
		Delete(ctx context.Context, in *input_basics.RoleDeleteInp) (err error)
		DataScopeSelect() (res input_form.Selects)
		DataScopeEdit(ctx context.Context, in *input_basics.DataScopeEditInp) (err error)
		// VerifyRoleId 验证角色ID
		VerifyRoleId(ctx context.Context, id int64) (err error)
		// GetSubRoleIds 获取所有下级角色ID
		GetSubRoleIds(ctx context.Context, roleId int64, isSuper bool) (ids []int64, err error)
	}
	IBasicsAdminSite interface {
		// Register 账号注册
		Register(ctx context.Context, in *input_basics.RegisterInp) (err error)
		// AccountLogin 账号登录
		AccountLogin(ctx context.Context, in *input_basics.AccountLoginInp) (res *input_basics.LoginModel, err error)
		// MobileLogin 手机号登录
		MobileLogin(ctx context.Context, in *input_basics.MobileLoginInp) (res *input_basics.LoginModel, err error)
		// BindUserContext 绑定用户上下文
		BindUserContext(ctx context.Context, claims *model.Identity) (err error)
	}
	IBasicsAppNotify interface {
		NotifyList(ctx context.Context, in *input_basics.PmsNotifyListInp) (list []*input_basics.PmsNotifyListModel, totalCount int, err error)
		NotifyEdit(ctx context.Context, in *input_basics.PmsNotifyEditInp) (err error)
		NotifyDelete(ctx context.Context, in *input_basics.PmsNotifyDeleteInp) (err error)
		NotifyView(ctx context.Context, in *input_basics.PmsNotifyViewInp) (res *input_basics.PmsNotifyViewModel, err error)
	}
	IBasicsPmsAppconfig interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_basics.PmsAppconfigListInp) (list []*input_basics.PmsAppconfigListModel, totalCount int, err error)
		Edit(ctx context.Context, in *input_basics.PmsAppconfigEditInp) (err error)
		Delete(ctx context.Context, in *input_basics.PmsAppconfigDeleteInp) (err error)
		View(ctx context.Context, in *input_basics.PmsAppconfigViewInp) (res *input_basics.PmsAppconfigViewModel, err error)
	}
	IBasicsAttachment interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		Delete(ctx context.Context, in *input_basics.AttachmentDeleteInp) (err error)
		View(ctx context.Context, in *input_basics.AttachmentViewInp) (res *input_basics.AttachmentViewModel, err error)
		List(ctx context.Context, in *input_basics.AttachmentListInp) (list []*input_basics.AttachmentListModel, totalCount int, err error)
		ClearKind(ctx context.Context, in *input_basics.AttachmentClearKindInp) (err error)
	}
	IBasicsAttachmentKind interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_basics.SysAttachmentKindListInp) (list []*input_basics.SysAttachmentKindListModel, totalCount int, err error)
		Edit(ctx context.Context, in *input_basics.SysAttachmentKindEditInp) (err error)
		Delete(ctx context.Context, in *input_basics.SysAttachmentKindDeleteInp) (err error)
		View(ctx context.Context, in *input_basics.SysAttachmentKindViewInp) (res *input_basics.SysAttachmentKindViewModel, err error)
	}
	IBasicsBanner interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_basics.PmsBannerListInp) (list []*input_basics.PmsBannerListModel, totalCount int, err error)
		Export(ctx context.Context, in *input_basics.PmsBannerListInp) (err error)
		Edit(ctx context.Context, in *input_basics.PmsBannerEditInp) (err error)
		Delete(ctx context.Context, in *input_basics.PmsBannerDeleteInp) (err error)
		View(ctx context.Context, in *input_basics.PmsBannerViewInp) (res *input_basics.PmsBannerViewModel, err error)
		BannerAppList(ctx context.Context, in *input_basics.PmsBannerListInp) (list []*input_basics.PmsBannerAppListModel, totalCount int, err error)
	}
	IBasicsBlacklist interface {
		Delete(ctx context.Context, in *input_basics.BlacklistDeleteInp) (err error)
		Edit(ctx context.Context, in *input_basics.BlacklistEditInp) (err error)
		Status(ctx context.Context, in *input_basics.BlacklistStatusInp) (err error)
		View(ctx context.Context, in *input_basics.BlacklistViewInp) (res *input_basics.BlacklistViewModel, err error)
		List(ctx context.Context, in *input_basics.BlacklistListInp) (list []*input_basics.BlacklistListModel, totalCount int, err error)
		VariableLoad(ctx context.Context, err error)
		Load(ctx context.Context)
		VerifyRequest(r *ghttp.Request) (err error)
		ClusterSync(ctx context.Context, message *gredis.Message)
	}
	IBasicsChannel interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_basics.PmsChannelListInp) (list []*input_basics.PmsChannelListModel, totalCount int, err error)
		Edit(ctx context.Context, in *input_basics.PmsChannelEditInp) (err error)
		Bind(ctx context.Context, in *input_basics.PmsChannelBindInp) (err error)
		Unbind(ctx context.Context, in *input_basics.PmsChannelUnbindInp) (err error)
		Delete(ctx context.Context, in *input_basics.PmsChannelDeleteInp) (err error)
		View(ctx context.Context, in *input_basics.PmsChannelViewInp) (res *input_basics.PmsChannelViewModel, err error)
		Status(ctx context.Context, in *input_basics.PmsChannelStatusInp) (err error)
	}
	IBasicsCollect interface {
		Add(ctx context.Context, in *input_basics.PmsCollectAddInp) (err error)
		List(ctx context.Context, in *input_basics.PmsCollectListInp) (list []*input_basics.PmsCollectListModel, totalCount int, err error)
		Delete(ctx context.Context, in *input_basics.PmsCollectDeleteInp) (err error)
	}
	IBasicsConfig interface {
		InitConfig(ctx context.Context)
		LoadConfig(ctx context.Context) (err error)
		GetLogin(ctx context.Context) (conf *model.LoginConfig, err error)
		GetPay(ctx context.Context) (conf *model.PayConfig, err error)
		GetSms(ctx context.Context) (conf *model.SmsConfig, err error)
		GetGeeTest(ctx context.Context) (conf *model.GeeTestConfig, err error)
		GetSpaSmsConfig(ctx context.Context) (conf *model.OrderSmsConfig, err error)
		GetLanguagePackSetting(ctx context.Context) (conf *model.LanguagePackSetting, err error)
		GetCarSmsConfig(ctx context.Context) (conf *model.OrderSmsConfig, err error)
		GetYYConfig(ctx context.Context) (conf *model.YYConfig, err error)
		GetWXShareConfig(ctx context.Context) (conf *model.WXShareConfig, err error)
		GetMemberRegRewardConfig(ctx context.Context) (conf *model.MemberRegRewardConfig, err error)
		GetInviteNewRewardConfig(ctx context.Context) (conf *model.InviteNewRewardConfig, err error)
		GetMemberIntentionConfig(ctx context.Context) (conf *model.MemberIntentionConfig, err error)
		GetCarPrinterSettingConfig(ctx context.Context) (conf *model.PrinterSettingConfig, err error)
		GetSpaPrinterSettingConfig(ctx context.Context) (conf *model.PrinterSettingConfig, err error)
		GetFoodPrinterSettingConfig(ctx context.Context) (conf *model.PrinterSettingConfig, err error)
		GetAppDataBoardViewConfig(ctx context.Context) (conf *model.AppDataBoardViewConfig, err error)
		GetGeo(ctx context.Context) (conf *model.GeoConfig, err error)
		GetUpload(ctx context.Context) (conf *model.UploadConfig, err error)
		GetPmsPrice(ctx context.Context) (conf *model.PmsPriceConfig, err error)
		GetSmtp(ctx context.Context) (conf *model.EmailConfig, err error)
		GetBasic(ctx context.Context) (conf *model.BasicConfig, err error)
		GetLoadTCP(ctx context.Context) (conf *model.TCPConfig, err error)
		GetApp(ctx context.Context) (conf *model.AppConfig, err error)
		GetLoadGenerate(ctx context.Context) (conf *model.GenerateConfig, err error)
		GetLoadToken(ctx context.Context) (conf *model.TokenConfig, err error)
		GetLoadLog(ctx context.Context) (conf *model.LogConfig, err error)
		GetLoadServeLog(ctx context.Context) (conf *model.ServeLogConfig, err error)
		GetConfigByGroup(ctx context.Context, in *input_basics.GetConfigInp) (res *input_basics.GetConfigModel, err error)
		ConversionType(ctx context.Context, models *entity.SysConfig) (value interface{}, err error)
		UpdateConfigByGroup(ctx context.Context, in *input_basics.UpdateConfigInp) (err error)
		SpaSmsConfigUpdateReq(ctx context.Context, in *input_basics.UpdateSpaSmsConfigInp) (err error)
		ClusterSync(ctx context.Context, message *gredis.Message)
		UpdateOrderMember(ctx context.Context) (err error)
		GetCabinetApi(ctx context.Context) (conf *model.CabinetApiConfig, err error)
		GetToretaApi(ctx context.Context) (conf *model.ToretaApiConfig, err error)
		GetXpyunApi(ctx context.Context) (conf *model.XpyunConfig, err error)
		RefreshToretaAccessToken(ctx context.Context) (err error)
		GetSystemMessageConfig(ctx context.Context) (conf *model.SystemMessageConfig, err error)
		GetHotelSettingConfig(ctx context.Context) (conf *model.HotelSettingConfig, err error)
	}
	IBasicsCoupon interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		View(ctx context.Context, in *input_basics.PmsCouponViewInp) (res *input_basics.PmsCouponViewModel, err error)
		List(ctx context.Context, in *input_basics.PmsCouponListInp) (list []*input_basics.PmsCouponListModel, totalCount int, err error)
		AppList(ctx context.Context, in *input_basics.PmsCouponListInp) (list []*input_basics.PmsCouponListModel, totalCount int, err error)
		Stat(ctx context.Context, in *input_basics.PmsCouponStatInp) (res *input_basics.PmsCouponStatModel, err error)
		Recycle(ctx context.Context, in *input_basics.PmsCouponRecycleInp) (err error)
		InvalidCoupon(ctx context.Context, in *input_basics.PmsCouponInvalidInp) (err error)
	}
	IBasicsCouponType interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_basics.PmsCouponTypeListInp) (list []*input_basics.PmsCouponTypeListModel, totalCount int, err error)
		All(ctx context.Context, in *input_basics.PmsCouponTypeListInp) (list []*input_basics.PmsCouponTypeAllListModel, err error)
		Edit(ctx context.Context, in *input_basics.PmsCouponTypeEditInp) (err error)
		Delete(ctx context.Context, in *input_basics.PmsCouponTypeDeleteInp) (err error)
		MaxSort(ctx context.Context, in *input_basics.PmsCouponTypeMaxSortInp) (res *input_basics.PmsCouponTypeMaxSortModel, err error)
		View(ctx context.Context, in *input_basics.PmsCouponTypeViewInp) (res *input_basics.PmsCouponTypeViewModel, err error)
		AppView(ctx context.Context, in *input_basics.PmsCouponTypeAppViewInp) (res *input_basics.PmsCouponTypeAppViewModel, err error)
		Status(ctx context.Context, in *input_basics.PmsCouponTypeStatusInp) (err error)
		SendMemberCoupon(ctx context.Context, in *input_basics.PmsSendMemberCouponInp, source int) (err error)
		SendMemberCouponGetId(ctx context.Context, in *input_basics.PmsSendMemberCouponInp, source int) (memberCouponId int64, couponNameLanguage []*input_hotel.LanguageType, err error)
		AppReceiveCouponType(ctx context.Context, in *input_basics.PmsCouponTypeAppReceiveInp) (err error)
	}
	IBasicsDashboard interface {
		GetBasics(ctx context.Context, in *input_basics.BasicDashboardInp) (res *input_basics.BasicDashboardModel, err error)
		GetRanking(ctx context.Context, in *input_basics.RankingInp) (out *input_basics.RankingModel, err error)
	}
	IBasicsAdminDept interface {
		// Model 部门ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// Delete 删除
		Delete(ctx context.Context, in *input_basics.DeptDeleteInp) (err error)
		// VerifyUnique 验证部门唯一属性
		VerifyUnique(ctx context.Context, in *input_basics.VerifyUniqueInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *input_basics.DeptEditInp) (err error)
		// MaxSort 最大排序
		MaxSort(ctx context.Context, in *input_basics.DeptMaxSortInp) (res *input_basics.DeptMaxSortModel, err error)
		// View 获取指定部门信息
		View(ctx context.Context, in *input_basics.DeptViewInp) (res *input_basics.DeptViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *input_basics.DeptListInp) (res *input_basics.DeptListModel, err error)
		// GetName 获取部门名称
		GetName(ctx context.Context, id int64) (name string, err error)
		// VerifyDeptId 验证部门ID
		VerifyDeptId(ctx context.Context, id int64) (err error)
		// Option 获取当前登录用户可选的部门选项
		Option(ctx context.Context, in *input_basics.DeptOptionInp) (res *input_basics.DeptOptionModel, totalCount int, err error)
		// TreeOption 获取部门关系树选项
		TreeOption(ctx context.Context) (nodes []tree.Node, err error)
	}
	IBasicsDictData interface {
		Delete(ctx context.Context, in *input_basics.DictDataDeleteInp) error
		Edit(ctx context.Context, in *input_basics.DictDataEditInp) (err error)
		List(ctx context.Context, in *input_basics.DictDataListInp) (list []*input_basics.DictDataListModel, totalCount int, err error)
		GetId(ctx context.Context, t string) (id int64, err error)
		GetType(ctx context.Context, id int64) (types string, err error)
		GetTypes(ctx context.Context, id int64) (types []string, err error)
		Select(ctx context.Context, in *input_basics.DataSelectInp) (list input_basics.DataSelectModel, err error)
	}
	IBasicsDictType interface {
		Tree(ctx context.Context) (list []*input_basics.DictTypeTree, err error)
		Delete(ctx context.Context, in *input_basics.DictTypeDeleteInp) (err error)
		Edit(ctx context.Context, in *input_basics.DictTypeEditInp) (err error)
		TreeSelect(ctx context.Context, in *input_basics.DictTreeSelectInp) (list []*input_basics.DictTypeTree, err error)
		BuiltinSelect() (list []*input_basics.DictTypeTree)
	}
	IBasicsEmsLog interface {
		Delete(ctx context.Context, in *input_basics.EmsLogDeleteInp) (err error)
		Edit(ctx context.Context, in *input_basics.EmsLogEditInp) (err error)
		Status(ctx context.Context, in *input_basics.EmsLogStatusInp) (err error)
		View(ctx context.Context, in *input_basics.EmsLogViewInp) (res *input_basics.EmsLogViewModel, err error)
		List(ctx context.Context, in *input_basics.EmsLogListInp) (list []*input_basics.EmsLogListModel, totalCount int, err error)
		Send(ctx context.Context, in *input_basics.SendEmsInp) (err error)
		// GetTemplate 获取指定邮件模板
		GetTemplate(ctx context.Context, template string, config *model.EmailConfig) (val string, err error)
		// AllowSend 是否允许发送
		AllowSend(ctx context.Context, models *entity.SysEmsLog, config *model.EmailConfig) (err error)
		// NowDayIpSendCount 当天 IP 累计发送次数
		NowDayIpSendCount(ctx context.Context, event ...string) (count int, err error)
		// VerifyCode 效验验证码
		VerifyCode(ctx context.Context, in *input_basics.VerifyEmsCodeInp) (err error)
	}
	IBasicsFinance interface {
		Stat(ctx context.Context, in *input_basics.FinanceStatInp) (out *input_basics.FinanceStatModel, err error)
		List(ctx context.Context, in *input_basics.FinanceListInp) (out *input_basics.FinanceListModel, err error)
		Export(ctx context.Context, in *input_basics.FinanceListInp) (err error)
	}
	IBasicsGenCodes interface {
		Delete(ctx context.Context, in *input_basics.GenCodesDeleteInp) (err error)
		Edit(ctx context.Context, in *input_basics.GenCodesEditInp) (res *input_basics.GenCodesEditModel, err error)
		Status(ctx context.Context, in *input_basics.GenCodesStatusInp) (err error)
		MaxSort(ctx context.Context, in *input_basics.GenCodesMaxSortInp) (res *input_basics.GenCodesMaxSortModel, err error)
		View(ctx context.Context, in *input_basics.GenCodesViewInp) (res *input_basics.GenCodesViewModel, err error)
		List(ctx context.Context, in *input_basics.GenCodesListInp) (list []*input_basics.GenCodesListModel, totalCount int, err error)
		Selects(ctx context.Context, in *input_basics.GenCodesSelectsInp) (res *input_basics.GenCodesSelectsModel, err error)
		TableSelect(ctx context.Context, in *input_basics.GenCodesTableSelectInp) (res []*input_basics.GenCodesTableSelectModel, err error)
		ColumnSelect(ctx context.Context, in *input_basics.GenCodesColumnSelectInp) (res []*input_basics.GenCodesColumnSelectModel, err error)
		ColumnList(ctx context.Context, in *input_basics.GenCodesColumnListInp) (res []*input_basics.GenCodesColumnListModel, err error)
		Preview(ctx context.Context, in *input_basics.GenCodesPreviewInp) (res *input_basics.GenCodesPreviewModel, err error)
		Build(ctx context.Context, in *input_basics.GenCodesBuildInp) (err error)
	}
	IBasicsHelpcenter interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_basics.PmsHelpcenterListInp) (list []*input_basics.PmsHelpcenterListModel, totalCount int, err error)
		Edit(ctx context.Context, in *input_basics.PmsHelpcenterEditInp) (err error)
		Delete(ctx context.Context, in *input_basics.PmsHelpcenterDeleteInp) (err error)
		View(ctx context.Context, in *input_basics.PmsHelpcenterViewInp) (res *input_basics.PmsHelpcenterViewModel, err error)
	}
	IBasicsHelpcenterCategory interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_basics.PmsHelpcenterCategoryListInp) (list []*input_basics.PmsHelpcenterCategoryListModel, totalCount int, err error)
		All(ctx context.Context, in *input_basics.PmsHelpcenterCategoryAllInp) (list []*input_basics.PmsHelpcenterCategoryAllModel, err error)
		Edit(ctx context.Context, in *input_basics.PmsHelpcenterCategoryEditInp) (err error)
		Delete(ctx context.Context, in *input_basics.PmsHelpcenterCategoryDeleteInp) (err error)
		View(ctx context.Context, in *input_basics.PmsHelpcenterCategoryViewInp) (res *input_basics.PmsHelpcenterCategoryViewModel, err error)
	}
	IBasicsLanguage interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_language.PmsLanguageListInp) (list []*input_language.PmsLanguageListModel, totalCount int, err error)
		Export(ctx context.Context, in *input_language.PmsLanguageListInp) (err error)
		Edit(ctx context.Context, in *input_language.PmsLanguageEditInp) (err error)
		Delete(ctx context.Context, in *input_language.PmsLanguageDeleteInp) (err error)
		View(ctx context.Context, in *input_language.PmsLanguageViewInp) (res *input_language.PmsLanguageViewModel, err error)
		Sync(ctx context.Context, LanguageStruct input_language.LanguageModel, In *input_language.LoadLanguage) (err error)
		GetUuids(ctx context.Context, content string, key string) (uuids []string, err error)
	}
	IBasicsLog interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		Export(ctx context.Context, in *input_basics.LogListInp) (err error)
		RealWrite(ctx context.Context, log entity.SysLog) (err error)
		AutoLog(ctx context.Context) error
		AnalysisLog(ctx context.Context) entity.SysLog
		View(ctx context.Context, in *input_basics.LogViewInp) (res *input_basics.LogViewModel, err error)
		Delete(ctx context.Context, in *input_basics.LogDeleteInp) (err error)
		List(ctx context.Context, in *input_basics.LogListInp) (list []*input_basics.LogListModel, totalCount int, err error)
	}
	IBasicsLoginLog interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_basics.LoginLogListInp) (list []*input_basics.LoginLogListModel, totalCount int, err error)
		Export(ctx context.Context, in *input_basics.LoginLogListInp) (err error)
		Delete(ctx context.Context, in *input_basics.LoginLogDeleteInp) (err error)
		Push(ctx context.Context, in *input_basics.LoginLogPushInp)
		RealWrite(ctx context.Context, models entity.SysLoginLog) (err error)
	}
	IBasicsPDF interface {
		OrderReceipt(ctx context.Context, ipt *input_basics.PdfLssIpt) (pdfFailPath string, err error)
	}
	IBasicsPrinter interface {
		// Model 打印机ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取打印机列表
		List(ctx context.Context, in *input_basics.PrinterListInp) (list []*input_basics.PrinterListModel, totalCount int, err error)
		// Edit 修改/新增打印机
		Edit(ctx context.Context, in *input_basics.PrinterEditInp) (err error)
		// Delete 删除打印机
		Delete(ctx context.Context, in *input_basics.PrinterDeleteInp) (err error)
		// MaxSort 获取打印机最大排序
		MaxSort(ctx context.Context, in *input_basics.PrinterMaxSortInp) (res *input_basics.PrinterMaxSortModel, err error)
		// View 获取打印机指定信息
		View(ctx context.Context, in *input_basics.PrinterViewInp) (res *input_basics.PrinterViewModel, err error)
		// Status 更新司机状态
		Status(ctx context.Context, in *input_basics.PrinterStatusInp) (err error)
		// PrinterCarOrder 打印接送机订单
		PrinterCarOrder(ctx context.Context, in *input_basics.PrinterCarOrderInp) (err error)
		// PrinterSpaOrder 打印按摩订单
		PrinterSpaOrder(ctx context.Context, in *input_basics.PrinterCarOrderInp) (err error)
		// PrinterFoodOrder 打印餐厅订单
		PrinterFoodOrder(ctx context.Context, in *input_basics.PrinterCarOrderInp) (err error)
		PrintOrderContent(ctx context.Context, in *input_basics.PrintContentInp) (err error)
		GetYILINKCarOrderContent(ctx context.Context, in *input_basics.GetYILINKCarOrderContentInp) (res *input_basics.GetPrintOrderContentModel, err error)
		GetXYUNCarOrderContent(ctx context.Context, in *input_basics.GetYILINKCarOrderContentInp) (res *input_basics.GetPrintOrderContentModel, err error)
		GetYILINKSpaOrderContent(ctx context.Context, in *input_basics.GetYILINKSpaOrderContentInp) (res *input_basics.GetPrintOrderContentModel, err error)
		GetXYUNSpaOrderContent(ctx context.Context, in *input_basics.GetYILINKSpaOrderContentInp) (res *input_basics.GetPrintOrderContentModel, err error)
		GetYILINKFoodOrderContent(ctx context.Context, in *input_basics.GetYILINKFoodOrderContentInp) (res *input_basics.GetPrintOrderContentModel, err error)
		GetXYUNFoodOrderContent(ctx context.Context, in *input_basics.GetYILINKFoodOrderContentInp) (res *input_basics.GetPrintOrderContentModel, err error)
	}
	IBasicsProvinces interface {
		Tree(ctx context.Context) (list []*input_basics.ProvincesTree, err error)
		Delete(ctx context.Context, in *input_basics.ProvincesDeleteInp) (err error)
		Edit(ctx context.Context, in *input_basics.ProvincesEditInp) (err error)
		Status(ctx context.Context, in *input_basics.ProvincesStatusInp) (err error)
		MaxSort(ctx context.Context, in *input_basics.ProvincesMaxSortInp) (res *input_basics.ProvincesMaxSortModel, err error)
		View(ctx context.Context, in *input_basics.ProvincesViewInp) (res *input_basics.ProvincesViewModel, err error)
		List(ctx context.Context, in *input_basics.ProvincesListInp) (list []*input_basics.ProvincesListModel, totalCount int, err error)
		ChildrenList(ctx context.Context, in *input_basics.ProvincesChildrenListInp) (list []*input_basics.ProvincesChildrenListModel, totalCount int, err error)
		UniqueId(ctx context.Context, in *input_basics.ProvincesUniqueIdInp) (res *input_basics.ProvincesUniqueIdModel, err error)
		Select(ctx context.Context, in *input_basics.ProvincesSelectInp) (res *input_basics.ProvincesSelectModel, err error)
	}
	IBasicsServeLog interface {
		Model(ctx context.Context) *gdb.Model
		List(ctx context.Context, in *input_basics.ServeLogListInp) (list []*input_basics.ServeLogListModel, totalCount int, err error)
		Export(ctx context.Context, in *input_basics.ServeLogListInp) (err error)
		Delete(ctx context.Context, in *input_basics.ServeLogDeleteInp) (err error)
		View(ctx context.Context, in *input_basics.ServeLogViewInp) (res *input_basics.ServeLogViewModel, err error)
		RealWrite(ctx context.Context, models entity.SysServeLog) (err error)
	}
	IBasicsSmsLog interface {
		Delete(ctx context.Context, in *input_basics.SmsLogDeleteInp) (err error)
		View(ctx context.Context, in *input_basics.SmsLogViewInp) (res *input_basics.SmsLogViewModel, err error)
		List(ctx context.Context, in *input_basics.SmsLogListInp) (list []*input_basics.SmsLogListModel, totalCount int, err error)
		SendCode(ctx context.Context, in *input_basics.SendCodeInp) (err error)
		// SendPlaceOrderMsg 发送下单预警短信
		SendPlaceOrderMsg(ctx context.Context, in *input_basics.SendMsgInp) (err error)
		GetTemplate(ctx context.Context, template string, config *model.SmsConfig) (val string, err error)
		AllowSend(ctx context.Context, models *entity.SysSmsLog, config *model.SmsConfig) (err error)
		NowDayIpSendCount(ctx context.Context) (count int, err error)
		VerifyCode(ctx context.Context, in *input_basics.VerifyCodeInp) (err error)
	}
	IBasicsSmsTemplate interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_basics.SmsTemplateListInp) (list []*input_basics.SmsTemplateListModel, totalCount int, err error)
		Edit(ctx context.Context, in *input_basics.SmsTemplateEditInp) (err error)
		Delete(ctx context.Context, in *input_basics.SmsTemplateDeleteInp) (err error)
		View(ctx context.Context, in *input_basics.SmsTemplateViewInp) (res *input_basics.SmsTemplateViewModel, err error)
		Log(ctx context.Context, in *input_basics.SmsTemplateLogInp) (list []*input_basics.SmsTemplateLogModel, totalCount int, err error)
		SendTemplate(ctx context.Context, in *input_basics.SendTemplateInp) (err error)
	}
	IBasicsStaff interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_basics.PmsStaffListInp) (list []*input_basics.PmsStaffListModel, totalCount int, err error)
		Edit(ctx context.Context, in *input_basics.PmsStaffEditInp) (err error)
		Bind(ctx context.Context, in *input_basics.PmsStaffBindInp) (err error)
		Unbind(ctx context.Context, in *input_basics.PmsStaffUnbindInp) (err error)
		Delete(ctx context.Context, in *input_basics.PmsStaffDeleteInp) (err error)
		View(ctx context.Context, in *input_basics.PmsStaffViewInp) (res *input_basics.PmsStaffViewModel, err error)
		Status(ctx context.Context, in *input_basics.PmsStaffStatusInp) (err error)
	}
	IBasicsSystemMessage interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		SendMessage(ctx context.Context, in *input_basics.SendMessageInp) (err error)
		AppList(ctx context.Context, in *input_basics.MessageAppListInp) (list []*input_basics.MessageAppListModel, totalCount int, err error)
		Read(ctx context.Context, in *input_basics.MessageAppReadInp) (err error)
		AppLatest(ctx context.Context, memberId int) (res *input_basics.MessageAppLatestModel, err error)
	}
	IBasicsTerminal interface {
		// Model 终端ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取终端列表
		List(ctx context.Context, in *input_basics.TerminalListInp) (list []*input_basics.TerminalListModel, totalCount int, err error)
		// Edit 修改/新增终端
		Edit(ctx context.Context, in *input_basics.TerminalEditInp) (err error)
		// Delete 删除终端
		Delete(ctx context.Context, in *input_basics.TerminalDeleteInp) (err error)
		// View 获取终端指定信息
		View(ctx context.Context, in *input_basics.TerminalViewInp) (res *input_basics.TerminalViewModel, err error)
		// BrandList 获取品牌型号列表
		BrandList(ctx context.Context, in *input_basics.BrandListInp) (list []*input_basics.BrandListModel, totalCount int, err error)
		// BrandEdit 修改/新增品牌型号
		BrandEdit(ctx context.Context, in *input_basics.BrandEditInp) (err error)
		// BrandView 获取终端指定信息
		BrandView(ctx context.Context, in *input_basics.BrandViewInp) (res *input_basics.BrandViewModel, err error)
		// PrinterTest 打印测试
		PrinterTest(ctx context.Context, in *input_basics.PrinterTestInp) (err error)
		// Printer 打印方法
		Printer(ctx context.Context, in *input_basics.PrinterInp) (err error)
	}
	IBasicsUpload interface {
		// UploadFile 上传文件
		UploadFile(ctx context.Context, uploadType string, file *ghttp.UploadFile, kind string) (res *input_basics.AttachmentListModel, err error)
		// CheckMultipart 检查文件分片
		CheckMultipart(ctx context.Context, in *input_basics.CheckMultipartInp) (res *input_basics.CheckMultipartModel, err error)
		// UploadPart 上传分片
		UploadPart(ctx context.Context, in *input_basics.UploadPartInp) (res *input_basics.UploadPartModel, err error)
	}
	IBasicsWithdraw interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_basics.PmsWithdrawListInp) (list []*input_basics.PmsWithdrawListModel, totalCount int, err error)
		View(ctx context.Context, in *input_basics.PmsWithdrawViewInp) (res *input_basics.PmsWithdrawViewModel, err error)
		Agree(ctx context.Context, in *input_basics.PmsWithdrawAgreeInp) (err error)
		Disagree(ctx context.Context, in *input_basics.PmsWithdrawDisagreeInp) (err error)
		Transfer(ctx context.Context, in *input_basics.PmsWithdrawTransferInp) (err error)
	}
	IBasicsWs interface {
		SendMemberWebsocketMessage(ctx context.Context, in *ws.SendWebsocketMessageInp) (err error)
	}
)

var (
	localBasicsIndexBanner        IBasicsIndexBanner
	localBasicsIndexNav           IBasicsIndexNav
	localBasicsTestNav            IBasicsTestNav
	localBasicsAdminMember        IBasicsAdminMember
	localBasicsAdminMemberPost    IBasicsAdminMemberPost
	localBasicsAdminMenu          IBasicsAdminMenu
	localBasicsAdminMonitor       IBasicsAdminMonitor
	localBasicsAdminNotice        IBasicsAdminNotice
	localBasicsAdminPost          IBasicsAdminPost
	localBasicsAdminRole          IBasicsAdminRole
	localBasicsAdminSite          IBasicsAdminSite
	localBasicsAppNotify          IBasicsAppNotify
	localBasicsPmsAppconfig       IBasicsPmsAppconfig
	localBasicsAttachment         IBasicsAttachment
	localBasicsAttachmentKind     IBasicsAttachmentKind
	localBasicsBanner             IBasicsBanner
	localBasicsBlacklist          IBasicsBlacklist
	localBasicsChannel            IBasicsChannel
	localBasicsCollect            IBasicsCollect
	localBasicsConfig             IBasicsConfig
	localBasicsCoupon             IBasicsCoupon
	localBasicsCouponType         IBasicsCouponType
	localBasicsDashboard          IBasicsDashboard
	localBasicsAdminDept          IBasicsAdminDept
	localBasicsDictData           IBasicsDictData
	localBasicsDictType           IBasicsDictType
	localBasicsEmsLog             IBasicsEmsLog
	localBasicsFinance            IBasicsFinance
	localBasicsGenCodes           IBasicsGenCodes
	localBasicsHelpcenter         IBasicsHelpcenter
	localBasicsHelpcenterCategory IBasicsHelpcenterCategory
	localBasicsLanguage           IBasicsLanguage
	localBasicsLog                IBasicsLog
	localBasicsLoginLog           IBasicsLoginLog
	localBasicsPDF                IBasicsPDF
	localBasicsPrinter            IBasicsPrinter
	localBasicsProvinces          IBasicsProvinces
	localBasicsServeLog           IBasicsServeLog
	localBasicsSmsLog             IBasicsSmsLog
	localBasicsSmsTemplate        IBasicsSmsTemplate
	localBasicsStaff              IBasicsStaff
	localBasicsSystemMessage      IBasicsSystemMessage
	localBasicsTerminal           IBasicsTerminal
	localBasicsUpload             IBasicsUpload
	localBasicsWithdraw           IBasicsWithdraw
	localBasicsWs                 IBasicsWs
)

func BasicsIndexBanner() IBasicsIndexBanner {
	if localBasicsIndexBanner == nil {
		panic("implement not found for interface IBasicsIndexBanner, forgot register?")
	}
	return localBasicsIndexBanner
}

func RegisterBasicsIndexBanner(i IBasicsIndexBanner) {
	localBasicsIndexBanner = i
}

func BasicsIndexNav() IBasicsIndexNav {
	if localBasicsIndexNav == nil {
		panic("implement not found for interface IBasicsIndexNav, forgot register?")
	}
	return localBasicsIndexNav
}

func RegisterBasicsIndexNav(i IBasicsIndexNav) {
	localBasicsIndexNav = i
}

func BasicsTestNav() IBasicsTestNav {
	if localBasicsTestNav == nil {
		panic("implement not found for interface IBasicsTestNav, forgot register?")
	}
	return localBasicsTestNav
}

func RegisterBasicsTestNav(i IBasicsTestNav) {
	localBasicsTestNav = i
}

func BasicsAdminMember() IBasicsAdminMember {
	if localBasicsAdminMember == nil {
		panic("implement not found for interface IBasicsAdminMember, forgot register?")
	}
	return localBasicsAdminMember
}

func RegisterBasicsAdminMember(i IBasicsAdminMember) {
	localBasicsAdminMember = i
}

func BasicsAdminMemberPost() IBasicsAdminMemberPost {
	if localBasicsAdminMemberPost == nil {
		panic("implement not found for interface IBasicsAdminMemberPost, forgot register?")
	}
	return localBasicsAdminMemberPost
}

func RegisterBasicsAdminMemberPost(i IBasicsAdminMemberPost) {
	localBasicsAdminMemberPost = i
}

func BasicsAdminMenu() IBasicsAdminMenu {
	if localBasicsAdminMenu == nil {
		panic("implement not found for interface IBasicsAdminMenu, forgot register?")
	}
	return localBasicsAdminMenu
}

func RegisterBasicsAdminMenu(i IBasicsAdminMenu) {
	localBasicsAdminMenu = i
}

func BasicsAdminMonitor() IBasicsAdminMonitor {
	if localBasicsAdminMonitor == nil {
		panic("implement not found for interface IBasicsAdminMonitor, forgot register?")
	}
	return localBasicsAdminMonitor
}

func RegisterBasicsAdminMonitor(i IBasicsAdminMonitor) {
	localBasicsAdminMonitor = i
}

func BasicsAdminNotice() IBasicsAdminNotice {
	if localBasicsAdminNotice == nil {
		panic("implement not found for interface IBasicsAdminNotice, forgot register?")
	}
	return localBasicsAdminNotice
}

func RegisterBasicsAdminNotice(i IBasicsAdminNotice) {
	localBasicsAdminNotice = i
}

func BasicsAdminPost() IBasicsAdminPost {
	if localBasicsAdminPost == nil {
		panic("implement not found for interface IBasicsAdminPost, forgot register?")
	}
	return localBasicsAdminPost
}

func RegisterBasicsAdminPost(i IBasicsAdminPost) {
	localBasicsAdminPost = i
}

func BasicsAdminRole() IBasicsAdminRole {
	if localBasicsAdminRole == nil {
		panic("implement not found for interface IBasicsAdminRole, forgot register?")
	}
	return localBasicsAdminRole
}

func RegisterBasicsAdminRole(i IBasicsAdminRole) {
	localBasicsAdminRole = i
}

func BasicsAdminSite() IBasicsAdminSite {
	if localBasicsAdminSite == nil {
		panic("implement not found for interface IBasicsAdminSite, forgot register?")
	}
	return localBasicsAdminSite
}

func RegisterBasicsAdminSite(i IBasicsAdminSite) {
	localBasicsAdminSite = i
}

func BasicsAppNotify() IBasicsAppNotify {
	if localBasicsAppNotify == nil {
		panic("implement not found for interface IBasicsAppNotify, forgot register?")
	}
	return localBasicsAppNotify
}

func RegisterBasicsAppNotify(i IBasicsAppNotify) {
	localBasicsAppNotify = i
}

func BasicsPmsAppconfig() IBasicsPmsAppconfig {
	if localBasicsPmsAppconfig == nil {
		panic("implement not found for interface IBasicsPmsAppconfig, forgot register?")
	}
	return localBasicsPmsAppconfig
}

func RegisterBasicsPmsAppconfig(i IBasicsPmsAppconfig) {
	localBasicsPmsAppconfig = i
}

func BasicsAttachment() IBasicsAttachment {
	if localBasicsAttachment == nil {
		panic("implement not found for interface IBasicsAttachment, forgot register?")
	}
	return localBasicsAttachment
}

func RegisterBasicsAttachment(i IBasicsAttachment) {
	localBasicsAttachment = i
}

func BasicsAttachmentKind() IBasicsAttachmentKind {
	if localBasicsAttachmentKind == nil {
		panic("implement not found for interface IBasicsAttachmentKind, forgot register?")
	}
	return localBasicsAttachmentKind
}

func RegisterBasicsAttachmentKind(i IBasicsAttachmentKind) {
	localBasicsAttachmentKind = i
}

func BasicsBanner() IBasicsBanner {
	if localBasicsBanner == nil {
		panic("implement not found for interface IBasicsBanner, forgot register?")
	}
	return localBasicsBanner
}

func RegisterBasicsBanner(i IBasicsBanner) {
	localBasicsBanner = i
}

func BasicsBlacklist() IBasicsBlacklist {
	if localBasicsBlacklist == nil {
		panic("implement not found for interface IBasicsBlacklist, forgot register?")
	}
	return localBasicsBlacklist
}

func RegisterBasicsBlacklist(i IBasicsBlacklist) {
	localBasicsBlacklist = i
}

func BasicsChannel() IBasicsChannel {
	if localBasicsChannel == nil {
		panic("implement not found for interface IBasicsChannel, forgot register?")
	}
	return localBasicsChannel
}

func RegisterBasicsChannel(i IBasicsChannel) {
	localBasicsChannel = i
}

func BasicsCollect() IBasicsCollect {
	if localBasicsCollect == nil {
		panic("implement not found for interface IBasicsCollect, forgot register?")
	}
	return localBasicsCollect
}

func RegisterBasicsCollect(i IBasicsCollect) {
	localBasicsCollect = i
}

func BasicsConfig() IBasicsConfig {
	if localBasicsConfig == nil {
		panic("implement not found for interface IBasicsConfig, forgot register?")
	}
	return localBasicsConfig
}

func RegisterBasicsConfig(i IBasicsConfig) {
	localBasicsConfig = i
}

func BasicsCoupon() IBasicsCoupon {
	if localBasicsCoupon == nil {
		panic("implement not found for interface IBasicsCoupon, forgot register?")
	}
	return localBasicsCoupon
}

func RegisterBasicsCoupon(i IBasicsCoupon) {
	localBasicsCoupon = i
}

func BasicsCouponType() IBasicsCouponType {
	if localBasicsCouponType == nil {
		panic("implement not found for interface IBasicsCouponType, forgot register?")
	}
	return localBasicsCouponType
}

func RegisterBasicsCouponType(i IBasicsCouponType) {
	localBasicsCouponType = i
}

func BasicsDashboard() IBasicsDashboard {
	if localBasicsDashboard == nil {
		panic("implement not found for interface IBasicsDashboard, forgot register?")
	}
	return localBasicsDashboard
}

func RegisterBasicsDashboard(i IBasicsDashboard) {
	localBasicsDashboard = i
}

func BasicsAdminDept() IBasicsAdminDept {
	if localBasicsAdminDept == nil {
		panic("implement not found for interface IBasicsAdminDept, forgot register?")
	}
	return localBasicsAdminDept
}

func RegisterBasicsAdminDept(i IBasicsAdminDept) {
	localBasicsAdminDept = i
}

func BasicsDictData() IBasicsDictData {
	if localBasicsDictData == nil {
		panic("implement not found for interface IBasicsDictData, forgot register?")
	}
	return localBasicsDictData
}

func RegisterBasicsDictData(i IBasicsDictData) {
	localBasicsDictData = i
}

func BasicsDictType() IBasicsDictType {
	if localBasicsDictType == nil {
		panic("implement not found for interface IBasicsDictType, forgot register?")
	}
	return localBasicsDictType
}

func RegisterBasicsDictType(i IBasicsDictType) {
	localBasicsDictType = i
}

func BasicsEmsLog() IBasicsEmsLog {
	if localBasicsEmsLog == nil {
		panic("implement not found for interface IBasicsEmsLog, forgot register?")
	}
	return localBasicsEmsLog
}

func RegisterBasicsEmsLog(i IBasicsEmsLog) {
	localBasicsEmsLog = i
}

func BasicsFinance() IBasicsFinance {
	if localBasicsFinance == nil {
		panic("implement not found for interface IBasicsFinance, forgot register?")
	}
	return localBasicsFinance
}

func RegisterBasicsFinance(i IBasicsFinance) {
	localBasicsFinance = i
}

func BasicsGenCodes() IBasicsGenCodes {
	if localBasicsGenCodes == nil {
		panic("implement not found for interface IBasicsGenCodes, forgot register?")
	}
	return localBasicsGenCodes
}

func RegisterBasicsGenCodes(i IBasicsGenCodes) {
	localBasicsGenCodes = i
}

func BasicsHelpcenter() IBasicsHelpcenter {
	if localBasicsHelpcenter == nil {
		panic("implement not found for interface IBasicsHelpcenter, forgot register?")
	}
	return localBasicsHelpcenter
}

func RegisterBasicsHelpcenter(i IBasicsHelpcenter) {
	localBasicsHelpcenter = i
}

func BasicsHelpcenterCategory() IBasicsHelpcenterCategory {
	if localBasicsHelpcenterCategory == nil {
		panic("implement not found for interface IBasicsHelpcenterCategory, forgot register?")
	}
	return localBasicsHelpcenterCategory
}

func RegisterBasicsHelpcenterCategory(i IBasicsHelpcenterCategory) {
	localBasicsHelpcenterCategory = i
}

func BasicsLanguage() IBasicsLanguage {
	if localBasicsLanguage == nil {
		panic("implement not found for interface IBasicsLanguage, forgot register?")
	}
	return localBasicsLanguage
}

func RegisterBasicsLanguage(i IBasicsLanguage) {
	localBasicsLanguage = i
}

func BasicsLog() IBasicsLog {
	if localBasicsLog == nil {
		panic("implement not found for interface IBasicsLog, forgot register?")
	}
	return localBasicsLog
}

func RegisterBasicsLog(i IBasicsLog) {
	localBasicsLog = i
}

func BasicsLoginLog() IBasicsLoginLog {
	if localBasicsLoginLog == nil {
		panic("implement not found for interface IBasicsLoginLog, forgot register?")
	}
	return localBasicsLoginLog
}

func RegisterBasicsLoginLog(i IBasicsLoginLog) {
	localBasicsLoginLog = i
}

func BasicsPDF() IBasicsPDF {
	if localBasicsPDF == nil {
		panic("implement not found for interface IBasicsPDF, forgot register?")
	}
	return localBasicsPDF
}

func RegisterBasicsPDF(i IBasicsPDF) {
	localBasicsPDF = i
}

func BasicsPrinter() IBasicsPrinter {
	if localBasicsPrinter == nil {
		panic("implement not found for interface IBasicsPrinter, forgot register?")
	}
	return localBasicsPrinter
}

func RegisterBasicsPrinter(i IBasicsPrinter) {
	localBasicsPrinter = i
}

func BasicsProvinces() IBasicsProvinces {
	if localBasicsProvinces == nil {
		panic("implement not found for interface IBasicsProvinces, forgot register?")
	}
	return localBasicsProvinces
}

func RegisterBasicsProvinces(i IBasicsProvinces) {
	localBasicsProvinces = i
}

func BasicsServeLog() IBasicsServeLog {
	if localBasicsServeLog == nil {
		panic("implement not found for interface IBasicsServeLog, forgot register?")
	}
	return localBasicsServeLog
}

func RegisterBasicsServeLog(i IBasicsServeLog) {
	localBasicsServeLog = i
}

func BasicsSmsLog() IBasicsSmsLog {
	if localBasicsSmsLog == nil {
		panic("implement not found for interface IBasicsSmsLog, forgot register?")
	}
	return localBasicsSmsLog
}

func RegisterBasicsSmsLog(i IBasicsSmsLog) {
	localBasicsSmsLog = i
}

func BasicsSmsTemplate() IBasicsSmsTemplate {
	if localBasicsSmsTemplate == nil {
		panic("implement not found for interface IBasicsSmsTemplate, forgot register?")
	}
	return localBasicsSmsTemplate
}

func RegisterBasicsSmsTemplate(i IBasicsSmsTemplate) {
	localBasicsSmsTemplate = i
}

func BasicsStaff() IBasicsStaff {
	if localBasicsStaff == nil {
		panic("implement not found for interface IBasicsStaff, forgot register?")
	}
	return localBasicsStaff
}

func RegisterBasicsStaff(i IBasicsStaff) {
	localBasicsStaff = i
}

func BasicsSystemMessage() IBasicsSystemMessage {
	if localBasicsSystemMessage == nil {
		panic("implement not found for interface IBasicsSystemMessage, forgot register?")
	}
	return localBasicsSystemMessage
}

func RegisterBasicsSystemMessage(i IBasicsSystemMessage) {
	localBasicsSystemMessage = i
}

func BasicsTerminal() IBasicsTerminal {
	if localBasicsTerminal == nil {
		panic("implement not found for interface IBasicsTerminal, forgot register?")
	}
	return localBasicsTerminal
}

func RegisterBasicsTerminal(i IBasicsTerminal) {
	localBasicsTerminal = i
}

func BasicsUpload() IBasicsUpload {
	if localBasicsUpload == nil {
		panic("implement not found for interface IBasicsUpload, forgot register?")
	}
	return localBasicsUpload
}

func RegisterBasicsUpload(i IBasicsUpload) {
	localBasicsUpload = i
}

func BasicsWithdraw() IBasicsWithdraw {
	if localBasicsWithdraw == nil {
		panic("implement not found for interface IBasicsWithdraw, forgot register?")
	}
	return localBasicsWithdraw
}

func RegisterBasicsWithdraw(i IBasicsWithdraw) {
	localBasicsWithdraw = i
}

func BasicsWs() IBasicsWs {
	if localBasicsWs == nil {
		panic("implement not found for interface IBasicsWs, forgot register?")
	}
	return localBasicsWs
}

func RegisterBasicsWs(i IBasicsWs) {
	localBasicsWs = i
}
