package logic_hotel

import (
	"APT/internal/dao"
	"APT/internal/library/airhousePublicApi"
	"APT/internal/library/hgorm"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_hotel"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"APT/utility/uuid"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// OrderChangeCheckIsExist 查询变更单是否存在
func (s *sHotelService) OrderChangeCheckIsExist(ctx context.Context, OrderId int, ChangeType string) (err error) {
	var (
		PmsAppReservationChange *entity.PmsAppReservationChange
	)

	if err = dao.PmsAppReservationChange.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppReservationChange.Columns().OrderId:      OrderId,
		dao.PmsAppReservationChange.Columns().ChangeType:   ChangeType,
		dao.PmsAppReservationChange.Columns().ChangeStatus: "DONE",
	}).Scan(&PmsAppReservationChange); err != nil {
		return
	}
	if !g.IsEmpty(PmsAppReservationChange) {
		// 已经进行过变更，无法继续变更
		err = gerror.New(gi18n.T(ctx, "the_change_has_already_been_made_and_cannot_continue"))
		return
	}
	return
}

// OrderChangeList 获取入住订单变更信息表列表
func (s *sHotelService) OrderChangeList(ctx context.Context, in *input_hotel.OrderChangeListInp) (list []*input_hotel.OrderChangeListModel, totalCount int, err error) {
	mod := dao.PmsAppReservationChange.Ctx(ctx)

	// 字段过滤
	mod = mod.Fields(input_hotel.OrderChangeListModel{})

	// 查询主键
	if in.Id > 0 {
		mod = mod.Where(dao.PmsAppReservationChange.Columns().Id, in.Id)
	}

	// 查询created_at
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsAppReservationChange.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.PmsAppReservationChange.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取入住订单变更信息表列表失败，请稍后重试！")
		return
	}
	return
}

// OrderChangeExport 导出入住订单变更信息表
func (s *sHotelService) OrderChangeExport(ctx context.Context, in *input_hotel.OrderChangeListInp) (err error) {
	list, _, err := s.OrderChangeList(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(input_hotel.OrderChangeExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出入住订单变更信息表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("入住订单变更信息表")
		exports   []input_hotel.OrderChangeExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// OrderChangeEdit 修改/新增入住订单变更信息表
func (s *sHotelService) OrderChangeEdit(ctx context.Context, in *input_hotel.OrderChangeEditInp) (err error) {
	// 验证'ChangeOrderSn'唯一
	if err = hgorm.IsUnique(ctx, &dao.PmsAppReservationChange, g.Map{dao.PmsAppReservationChange.Columns().ChangeOrderSn: in.ChangeOrderSn}, "变更订单号已存在", in.Id); err != nil {
		return
	}
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = dao.PmsAppReservationChange.Ctx(ctx).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改入住订单变更信息表失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = dao.PmsAppReservationChange.Ctx(ctx).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增入住订单变更信息表失败，请稍后重试！")
		}
		return
	})
}

// OrderChangeDelete 删除入住订单变更信息表
func (s *sHotelService) OrderChangeDelete(ctx context.Context, in *input_hotel.OrderChangeDeleteInp) (err error) {

	if _, err = dao.PmsAppReservationChange.Ctx(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除入住订单变更信息表失败，请稍后重试！")
		return
	}
	return
}

// OrderChangeView 获取入住订单变更信息表指定信息
func (s *sHotelService) OrderChangeView(ctx context.Context, in *input_hotel.OrderChangeViewInp) (res *input_hotel.OrderChangeViewModel, err error) {
	if err = dao.PmsAppReservationChange.Ctx(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取入住订单变更信息表信息，请稍后重试！")
		return
	}
	return
}

// OrderChangeGuestEdit 订单变更用户信息
func (s *sHotelService) OrderChangeGuestEdit(ctx context.Context, in *input_hotel.OrderChangeGuestReq) (err error) {
	var (
		appReservation     *entity.PmsAppReservation
		appStay            *entity.PmsAppStay
		guestProfile       *entity.PmsGuestProfile
		ChangeOrderSn      = uuid.CreateOrderCode("HC")
		OldChangeGUestInfo = g.MapStrAny{}
		NewChangeGuestInfo = g.MapStrAny{}
		ChangeInsertId     int64
		ExpirationTime     int
		PayConfig          *model.PayConfig
	)
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 查询入住单
		if err = dao.PmsAppReservation.Ctx(ctx).TX(tx).Where(g.Map{
			dao.PmsAppReservation.Columns().Uuid: in.AirhostOrderUuid,
		}).Scan(&appReservation); err != nil {
			return
		}
		if g.IsEmpty(appReservation) {
			err = gerror.New("不存在入住单")
			return
		}
		// 查询主订单
		if err = dao.PmsAppStay.Ctx(ctx).TX(tx).
			Where(dao.PmsAppStay.Columns().OrderSn, appReservation.OrderSn).
			Scan(&appStay); err != nil {
			return
		}
		if g.IsEmpty(appStay) {
			err = gerror.New("不存在主订单")
			return
		}

		// 查询预定人信息
		if err = dao.PmsGuestProfile.Ctx(ctx).TX(tx).Where(g.Map{
			dao.PmsGuestProfile.Columns().Uid: appReservation.MainGuest,
		}).Scan(&guestProfile); err != nil {
			return
		}
		if g.IsEmpty(guestProfile) {
			err = gerror.New("预定人信息不存在")
			return
		}
		OldChangeGUestInfo = g.MapStrAny{
			"first_name": guestProfile.FirstName,
			"last_name":  guestProfile.LastName,
			"full_name":  guestProfile.FullName,
			"email":      guestProfile.Email,
			"phone":      guestProfile.Phone,
			"area_no":    guestProfile.AreaNo,
		}

		guestInfo := &entity.PmsGuestProfile{
			FirstName: in.FirstName,
			LastName:  in.LastName,
			FullName:  in.LastName + " " + in.FirstName,
			Email:     in.Email,
			Phone:     in.Phone,
			AreaNo:    in.AreaNo,
		}
		// 更新入住单入住人信息
		if _, err = dao.PmsGuestProfile.Ctx(ctx).TX(tx).Where(g.Map{
			dao.PmsGuestProfile.Columns().Uid: appReservation.MainGuest,
		}).OmitEmptyData().Data(guestInfo).Update(); err != nil {
			err = gerror.Wrap(err, "修改入住人信息失败")
			return
		}

		// 更新入住单入住人信息
		if _, err = dao.PmsGuestProfile.Ctx(ctx).TX(tx).Where(g.Map{
			dao.PmsGuestProfile.Columns().Uid: appStay.Booker,
		}).OmitEmptyData().Data(guestInfo).Update(); err != nil {
			err = gerror.Wrap(err, "修改入住人信息失败")
			return
		}

		NewChangeGuestInfo = g.MapStrAny{
			"first_name": guestProfile.FirstName,
			"last_name":  guestProfile.LastName,
			"full_name":  guestProfile.FullName,
			"email":      guestProfile.Email,
			"phone":      guestProfile.Phone,
			"area_no":    guestProfile.AreaNo,
		}
		if !g.IsEmpty(in.FirstName) {
			NewChangeGuestInfo["first_name"] = in.FirstName
		}
		if !g.IsEmpty(in.LastName) {
			NewChangeGuestInfo["last_name"] = in.LastName
		}
		NewChangeGuestInfo["full_name"] = fmt.Sprintf("%s %s", NewChangeGuestInfo["last_name"], NewChangeGuestInfo["first_name"])
		if !g.IsEmpty(in.Email) {
			NewChangeGuestInfo["email"] = in.Email
		}
		if !g.IsEmpty(in.Phone) {
			NewChangeGuestInfo["phone"] = in.Phone
		}
		if !g.IsEmpty(in.AreaNo) {
			NewChangeGuestInfo["area_no"] = in.AreaNo
		}

		if _, err = airhousePublicApi.PutGuestProfile(ctx, appStay.Booker, NewChangeGuestInfo); err != nil {
			return
		}

		if _, err = airhousePublicApi.PutGuestProfile(ctx, appReservation.MainGuest, NewChangeGuestInfo); err != nil {
			return
		}

		// 读取支付配置
		if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
			return
		}
		ExpirationTime = gvar.New(gtime.Now().Unix() + PayConfig.HotelStayExp).Int()

		// 插入变更信息
		if ChangeInsertId, err = dao.PmsAppReservationChange.Ctx(ctx).TX(tx).OmitEmptyData().Data(&entity.PmsAppReservationChange{
			ChangeOrderSn:  ChangeOrderSn,
			OrderId:        appReservation.Id,
			OrderSn:        appReservation.OrderSn,
			OutOrderSn:     appReservation.OutOrderSn,
			ChangeType:     "GUEST",
			OldMainGuest:   gjson.New(OldChangeGUestInfo),
			NewMainGuest:   gjson.New(NewChangeGuestInfo),
			ChangeStatus:   "DONE",
			ChangeAmount:   0,
			SubmitDate:     gtime.Now(),
			DoneDate:       gtime.Now(),
			ExpirationTime: ExpirationTime,
			IsFx:           appReservation.IsFx,
		}).InsertAndGetId(); err != nil {
			return
		}
		if g.IsEmpty(appReservation.IsFx) {
			_, _ = dao.PmsAppReservationChange.Ctx(ctx).TX(tx).WherePri(ChangeInsertId).Update(g.MapStrAny{
				dao.PmsAppReservationChange.Columns().IsFx: "N",
			})
		}
		// 回写变更成功
		if _, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).WherePri(in.OrderId).Data(g.MapStrAny{
			dao.PmsAppReservation.Columns().IsChangeGuest:   "Y",
			dao.PmsAppReservation.Columns().IsChangeGuestId: ChangeInsertId,
		}).Update(); err != nil {
			return
		}

		// 酒店订单日志
		if _, err = dao.PmsAppStayLog.Ctx(ctx).OmitEmptyData().Insert(&entity.PmsAppStayLog{
			OrderId:     appStay.Id,
			ActionWay:   "CHANGE_GUEST",
			Remark:      strconv.FormatInt(ChangeInsertId, 10),
			OperateType: "USER",
			OperateId:   appStay.MemberId,
		}); err != nil {
			return
		}

		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "订单变更成功",
			"en":    "Order change successful",
			"ja":    "注文の変更が完了しました",
			"ko":    "주문 변경 성공",
			"zh_CN": "訂單變更成功",
		}
		systemMessageContent := map[string]string{
			"zh":    "订单预定人信息变更成功",
			"en":    "Order reservation holder information successfully changed",
			"ja":    "注文予約者情報が正常に変更されました",
			"ko":    "주문 예약자 정보가 성공적으로 변경되었습니다.",
			"zh_CN": "訂單預定人資訊變更成功",
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(appStay.MemberId).Value()
		var memberLanguage string
		if phoneArea.String() == "+86" {
			memberLanguage = "zh"
		} else if phoneArea.String() == "+81" {
			memberLanguage = "ja"
		} else if phoneArea.String() == "+82" {
			memberLanguage = "ko"
		} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
			memberLanguage = "zh_CN"
		} else {
			memberLanguage = "en"
		}
		pushData := g.MapStrAny{
			"type":    0,
			"orderSn": appReservation.OrderSn,
		}
		pushDataJson, _ := json.Marshal(pushData)
		appPushData := g.MapStrStr{
			"type":  "2",
			"param": string(pushDataJson),
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "hotel",
			Type:                 "order",
			MemberId:             int(appStay.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/order/order-detail",
			WxLink:               fmt.Sprintf("/pages/orders/details?orderSn=%s", appReservation.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           appStay.MemberId,
			OperatorRole:         "MEMBER",
			OrderSn:              appReservation.OrderSn,
		})

		return
	}); err != nil {
		return
	}
	return
}
