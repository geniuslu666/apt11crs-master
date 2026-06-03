package app

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/model"
	"APT/internal/service"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gtime"

	"APT/api/app/basics"
)

func (c *ControllerBasics) Dashboard(ctx context.Context, req *basics.DashboardReq) (res *basics.DashboardRes, err error) {
	var (
		TotalAPTHotelOrderNum int
		TodayAPTHotelOrderNum int
	)
	res = new(basics.DashboardRes)

	// 获取当前用户的数据看板权限
	res.Permissions = getDashboardPermissions(ctx)

	todayStart := gtime.Now().StartOfDay()
	todayEnd := gtime.Now().EndOfDay()

	lastOneDayStart := gtime.New(todayStart).Add(time.Duration(-24) * time.Hour)
	lastOneDayEnd := gtime.New(todayEnd).Add(time.Duration(-24) * time.Hour)

	// 会员总数
	res.PrivateDetail.MemberTotal, err = dao.PmsMember.Ctx(ctx).Count()
	// 今日增加会员
	res.PrivateDetail.TodayMemberAdd, err = dao.PmsMember.Ctx(ctx).WhereBetween(dao.PmsMember.Columns().CreatedAt, todayStart, todayEnd).Count()
	// 昨日新增会员数
	res.PrivateDetail.YesterdayMemberAdd, err = dao.PmsMember.Ctx(ctx).WhereBetween(dao.PmsMember.Columns().CreatedAt, lastOneDayStart, lastOneDayEnd).Count()

	// 今日民宿下单数
	TodayAPTHotelOrderNum, err = dao.PmsAppStay.Ctx(ctx).
		WhereBetween(dao.PmsAppStay.Columns().CreatedAt, todayStart, todayEnd).
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		Where(dao.PmsAppStay.Columns().Source, "APP").
		Count()
	res.PrivateDetail.TodayHotelOrderNum = TodayAPTHotelOrderNum

	// 昨日民宿下单数
	res.PrivateDetail.YesterdayHotelOrderNum, err = dao.PmsAppStay.Ctx(ctx).
		WhereBetween(dao.PmsAppStay.Columns().CreatedAt, lastOneDayStart, lastOneDayEnd).
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		Where(dao.PmsAppStay.Columns().Source, "APP").
		Count()

	// 民宿下单总数
	TotalAPTHotelOrderNum, err = dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		Where(dao.PmsAppStay.Columns().Source, "APP").
		Count()
	res.PrivateDetail.TotalHotelOrderNum = TotalAPTHotelOrderNum

	// IOS 来源
	res.RegSourceDetail.IOSNum, _ = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Source, "IOS").Count()

	// android 来源
	res.RegSourceDetail.AndroidNum, _ = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Source, "Android").Count()

	// Weixin 来源
	res.RegSourceDetail.WeixinNum, _ = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Source, "WX_MINI").Count()

	// H5 来源
	res.RegSourceDetail.H5Num, _ = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Source, "H5").Count()

	// 民宿下单总数
	res.TotalDetail.TotalHotelOrderNum, err = dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		Count()
	// 今日民宿下单总数
	res.TotalDetail.TodayHotelOrderNum, err = dao.PmsAppStay.Ctx(ctx).
		WhereBetween(dao.PmsAppStay.Columns().CreatedAt, todayStart, todayEnd).
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		Count()

	// APT民宿下单总数
	res.TotalDetail.TotalAPTHotelOrderNum = TotalAPTHotelOrderNum

	// 今日APT民宿下单总数
	res.TotalDetail.TodayAPTHotelOrderNum = TodayAPTHotelOrderNum

	// OTA民宿下单总数
	res.TotalDetail.TotalOTAHotelOrderNum, err = dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		Where(dao.PmsAppStay.Columns().Source, "AIRHOST").
		Count()

	// 今日OTA民宿下单总数
	res.TotalDetail.TodayOTAHotelOrderNum, err = dao.PmsAppStay.Ctx(ctx).
		WhereBetween(dao.PmsAppStay.Columns().CreatedAt, todayStart, todayEnd).
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		Where(dao.PmsAppStay.Columns().Source, "AIRHOST").
		Count()

	// 车队数据
	// 车队总订单数（已支付+已退款）
	res.CarData.TotalCarOrderNum, err = dao.CarOrder.Ctx(ctx).
		Where(dao.CarOrder.Columns().PayStatus, "HAVE_PAID").
		Count()
	// 今日预约数
	res.CarData.TodayCarOrderNum, err = dao.CarOrder.Ctx(ctx).
		Where(dao.CarOrder.Columns().PayStatus, "HAVE_PAID").
		Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
		Count()
	// 今日预约接机数
	res.CarData.TodayPickUpOrderNum, err = dao.CarOrder.Ctx(ctx).
		Where(dao.CarOrder.Columns().ServiceType, "PICKUP").
		Where(dao.CarOrder.Columns().PayStatus, "HAVE_PAID").
		Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
		Count()
	// 今日预约送机数
	res.CarData.TodayDeliveryOrderNum, err = dao.CarOrder.Ctx(ctx).
		Where(dao.CarOrder.Columns().ServiceType, "DELIVERY").
		Where(dao.CarOrder.Columns().PayStatus, "HAVE_PAID").
		Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
		Count()
	// 累计营业额(已支付)
	res.CarData.TotalOrderMoney, err = dao.CarOrder.Ctx(ctx).
		WhereNot(dao.SpaOrder.Columns().OrderStatus, "CANCEL").
		Where(dao.CarOrder.Columns().PayStatus, "HAVE_PAID").
		Sum(dao.CarOrder.Columns().OrderAmount)
	// 今日营业额(已支付)
	res.CarData.TodayOrderMoney, err = dao.CarOrder.Ctx(ctx).
		WhereNot(dao.SpaOrder.Columns().OrderStatus, "CANCEL").
		Where(dao.CarOrder.Columns().PayStatus, "HAVE_PAID").
		Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
		Sum(dao.CarOrder.Columns().OrderAmount)

	// 拒单列表
	if err = dao.CarOrder.Ctx(ctx).
		Where(dao.CarOrder.Columns().OrderStatus, "CANCEL").
		WhereNotNull(dao.CarOrder.Columns().PayTime).
		WhereNot(dao.CarOrder.Columns().ConfirmRefuseReason, "").
		Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
		Fields(dao.CarOrder.Columns().Id, dao.CarOrder.Columns().OrderSn, dao.CarOrder.Columns().ServiceType, dao.CarOrder.Columns().BookStartTime, dao.CarOrder.Columns().ConfirmRefuseReason).
		OrderAsc(dao.CarOrder.Columns().Id).Scan(&res.CarData.RefuseOrderList); err != nil {
		return
	}
	// 为拒单列表设置拒绝类型
	for _, item := range res.CarData.RefuseOrderList {
		item.ConfirmRefuseReason = "系统拒单"
	}

	// 按摩数据
	// 总订单数
	res.SpaData.TotalSpaOrderNum, err = dao.SpaOrder.Ctx(ctx).
		Where(dao.SpaOrder.Columns().PayStatus, "HAVE_PAID").
		Count()
	// 今日预约数
	res.SpaData.TodaySpaOrderNum, err = dao.SpaOrder.Ctx(ctx).
		Where(dao.SpaOrder.Columns().PayStatus, "HAVE_PAID").
		Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
		Count()
	// 今日上门预约数
	res.SpaData.TodayDoorToDoorOrderNum, err = dao.SpaOrder.Ctx(ctx).
		Where(dao.SpaOrder.Columns().ServiceType, 2).
		Where(dao.SpaOrder.Columns().PayStatus, "HAVE_PAID").
		Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
		Count()
	// 今日到店预约数
	res.SpaData.TodayInStoreOrderNum, err = dao.SpaOrder.Ctx(ctx).
		Where(dao.SpaOrder.Columns().ServiceType, 1).
		Where(dao.SpaOrder.Columns().PayStatus, "HAVE_PAID").
		Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
		Count()
	// 累计营业额(不包含退款)
	res.SpaData.TotalOrderMoney, err = dao.SpaOrder.Ctx(ctx).
		WhereNot(dao.SpaOrder.Columns().OrderStatus, "CANCEL").
		Where(dao.SpaOrder.Columns().PayStatus, "HAVE_PAID").
		Sum(dao.SpaOrder.Columns().OrderAmount)
	// 今日营业额(不包含退款)
	res.SpaData.TodayOrderMoney, err = dao.SpaOrder.Ctx(ctx).
		WhereNot(dao.SpaOrder.Columns().OrderStatus, "CANCEL").
		Where(dao.SpaOrder.Columns().PayStatus, "HAVE_PAID").
		Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
		Sum(dao.SpaOrder.Columns().OrderAmount)
	// 拒单列表
	if err = dao.SpaOrder.Ctx(ctx).
		Where(dao.SpaOrder.Columns().OrderStatus, "CANCEL").
		WhereNotNull(dao.SpaOrder.Columns().PayTime).
		WhereNot(dao.SpaOrder.Columns().ConfirmRefuseReason, "").
		Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
		Fields(dao.SpaOrder.Columns().Id, dao.SpaOrder.Columns().OrderSn, dao.CarOrder.Columns().ServiceType, dao.SpaOrder.Columns().BookStartTime, dao.SpaOrder.Columns().ConfirmRefuseReason).
		OrderAsc(dao.SpaOrder.Columns().Id).Scan(&res.SpaData.RefuseOrderList); err != nil {
		return
	}
	// 为拒单列表设置拒绝类型
	for _, item := range res.SpaData.RefuseOrderList {
		item.ConfirmRefuseReason = "系统拒单"
	}
	res.UpdateTime = gtime.Now()
	return

}

// getDashboardPermissions 获取当前用户的数据看板权限
func getDashboardPermissions(ctx context.Context) []string {
	UserInfo := contexts.GetMemberUser(ctx)
	if UserInfo == nil {
		return nil
	}

	// 获取配置
	var AppDataBoardViewConfig *model.AppDataBoardViewConfig
	var err error
	if AppDataBoardViewConfig, err = service.BasicsConfig().GetAppDataBoardViewConfig(ctx); err != nil {
		return nil
	}

	if AppDataBoardViewConfig.TestType == 1 {
		// 会员分组 - 分组用户拥有全部权限
		pmsMemberGroupIds := AppDataBoardViewConfig.CanTestGroupIds
		pmsMemberGroupIdsList := strings.Split(pmsMemberGroupIds, ",")
		// 获取用户的分组ID
		memberGroupId, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, UserInfo.Id).Value(dao.PmsMember.Columns().GroupId)
		for _, v := range pmsMemberGroupIdsList {
			if gvar.New(v).Int() == memberGroupId.Int() {
				return []string{"member", "hotel", "spa", "car"}
			}
		}
	} else {
		// 指定会员 - 根据配置的权限返回
		pmsMemberIds := AppDataBoardViewConfig.CanTestMemberIds
		pmsMemberIdsArr := strings.Split(pmsMemberIds, ",")
		for _, v := range pmsMemberIdsArr {
			if gvar.New(v).Int() == UserInfo.Id {
				// 解析会员权限配置
				permissions := parseDashboardMemberPermissions(AppDataBoardViewConfig.MemberPermissions, UserInfo.Id)
				if len(permissions) == 0 {
					// 如果没有配置权限，默认全部权限
					return []string{"member", "hotel", "spa", "car"}
				}
				return permissions
			}
		}
	}
	return nil
}

// parseDashboardMemberPermissions 解析会员权限配置
// 格式: memberId:permission1,permission2|memberId:permission1
func parseDashboardMemberPermissions(memberPermissions string, memberId int) []string {
	if memberPermissions == "" {
		return nil
	}
	memberIdStr := fmt.Sprintf("%d", memberId)
	// 按 | 分割每个会员的权限配置
	memberPerms := strings.Split(memberPermissions, "|")
	for _, mp := range memberPerms {
		// 按 : 分割会员ID和权限
		parts := strings.Split(mp, ":")
		if len(parts) != 2 {
			continue
		}
		if parts[0] == memberIdStr {
			// 找到该会员的权限配置
			if parts[1] == "" {
				return nil
			}
			return strings.Split(parts[1], ",")
		}
	}
	return nil
}
