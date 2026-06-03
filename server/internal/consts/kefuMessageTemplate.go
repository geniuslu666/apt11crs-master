package consts

import "github.com/gogf/gf/v2/frame/g"

// 客服订单模板
const (
	TemplateReservationZh = `html[
<p>物业名称：{{.PropertyDetail.Name}}</p>
<p>住客姓名：{{.GuestProfileDetail.FullName}}</p>
<p>入住时间：{{.CheckinTime}}</p>
<p>离开时间：{{.CheckoutTime}}</p>
<p>房间型号：{{.RoomTypeDetail.Name}}</p>
<p>订单总价：￥{{.BookingFee}} JPY</p>
<p>订单编号：{{.OrderSn}}</p>
<p>预订编号：{{.OutOrderSn}}</p>
<p>下单时间：{{.CreatedAt}}</p>]`

	TemplateReservationZhCN = `html[
<p>物業名稱：{{.PropertyDetail.Name}}</p>
<p>住客姓名：{{.GuestProfileDetail.FullName}}</p>
<p>入住時間：{{.CheckinTime}}</p>
<p>離開時間：{{.CheckoutTime}}</p>
<p>房間型號：{{.RoomTypeDetail.Name}}</p>
<p>訂單總價：￥{{.BookingFee}} JPY</p>
<p>訂單編號：{{.OrderSn}}</p>
<p>預訂編號：{{.OutOrderSn}}</p>
<p>下單時間：{{.CreatedAt}}</p>]`

	TemplateReservationEn = `html[
<p>Property Name: {{.PropertyDetail.Name}}</p>
<p>Guest Name: {{.GuestProfileDetail.FullName}}</p>
<p>Check-in Time: {{.CheckinTime}}</p>
<p>Check-out Time: {{.CheckoutTime}}</p>
<p>Room Type: {{.RoomTypeDetail.Name}}</p>
<p>Total Price: ¥{{.BookingFee}} JPY</p>
<p>Order No.: {{.OrderSn}}</p>
<p>Reservation No：{{.OutOrderSn}}</p>
<p>Order Date: {{.CreatedAt}}</p>]`

	TemplateReservationJa = `html[
<p>施設名：{{.PropertyDetail.Name}}</p>
<p>ゲスト名：{{.GuestProfileDetail.FullName}}</p>
<p>チェックイン時間：{{.CheckinTime}}</p>
<p>チェックアウト時間：{{.CheckoutTime}}</p>
<p>ルームタイプ：{{.RoomTypeDetail.Name}}</p>
<p>注文総額：￥{{.BookingFee}} JPY</p>
<p>注文番号：{{.OrderSn}}</p>
<p>予約番号：{{.OutOrderSn}}</p>
<p>注文日時：{{.CreatedAt}}</p>]`

	TemplateReservationKo = `html[
<p>물업 이름：{{.PropertyDetail.Name}}</p>
<p>게스트 이름：{{.GuestProfileDetail.FullName}}</p>
<p>체크인 시간：{{.CheckinTime}}</p>
<p>체크아웃 시간：{{.CheckoutTime}}</p>
<p>방 유형：{{.RoomTypeDetail.Name}}</p>
<p>주문 총액：￥{{.BookingFee}} JPY</p>
<p>주문 번호：{{.OrderSn}}</p>
<p>예약 번호：{{.OutOrderSn}}</p>
<p>주문 일시：{{.CreatedAt}}</p>]`
)

const (
	TemplateOrderZh = `html[
<p>物业名称：{{.PropertyDetail.Name}}</p>
<p>住客姓名：{{.GuestProfileDetail.FullName}}</p>
<p>入住时间：{{.CheckInDate}}</p>
<p>离开时间：{{.CheckOutDate}}</p>
<p>订单总价：￥{{.OrderAmount}} JPY</p>
<p>订单编号：{{.OrderSn}}</p>
<p>预订编号：{{.OutOrderSn}}</p>
<p>下单时间：{{.CreatedAt}}</p>]`

	TemplateOrderZhCN = `html[
<p>物業名稱：{{.PropertyDetail.Name}}</p>
<p>住客姓名：{{.GuestProfileDetail.FullName}}</p>
<p>入住時間：{{.CheckInDate}}</p>
<p>離開時間：{{.CheckOutDate}}</p>
<p>訂單總價：￥{{.OrderAmount}} JPY</p>
<p>訂單編號：{{.OrderSn}}</p>
<p>預訂編號：{{.OutOrderSn}}</p>
<p>下單時間：{{.CreatedAt}}</p>]`

	TemplateOrderEn = `html[
<p>Property Name: {{.PropertyDetail.Name}}</p>
<p>Guest Name: {{.GuestProfileDetail.FullName}}</p>
<p>Check-in Time: {{.CheckInDate}}</p>
<p>Check-out Time: {{.CheckOutDate}}</p>
<p>Total Price: ¥{{.OrderAmount}} JPY</p>
<p>Order No.: {{.OrderSn}}</p>
<p>Reservation No：{{.OutOrderSn}}</p>
<p>Order Date: {{.CreatedAt}}</p>]`

	TemplateOrderJa = `html[
<p>施設名：{{.PropertyDetail.Name}}</p>
<p>ゲスト名：{{.GuestProfileDetail.FullName}}</p>
<p>チェックイン時間：{{.CheckInDate}}</p>
<p>チェックアウト時間：{{.CheckOutDate}}</p>
<p>注文総額：￥{{.OrderAmount}} JPY</p>
<p>注文番号：{{.OrderSn}}</p>
<p>予約番号：{{.OutOrderSn}}</p>
<p>注文日時：{{.CreatedAt}}</p>]`

	TemplateOrderKo = `html[
<p>물업 이름：{{.PropertyDetail.Name}}</p>
<p>게스트 이름：{{.GuestProfileDetail.FullName}}</p>
<p>체크인 시간：{{.CheckInDate}}</p>
<p>체크아웃 시간：{{.CheckOutDate}}</p>
<p>주문 총액：￥{{.OrderAmount}} JPY</p>
<p>주문 번호：{{.OrderSn}}</p>
<p>예약 번호：{{.OutOrderSn}}</p>
<p>주문 일시：{{.CreatedAt}}</p>]`
)

// 客服物业模板

const (
	TemplatePropertyZh = `html[
<p>物业名称：{{.Name}}</p>
<p>物业地址：{{.Address}}</p>
<p>物业电话：{{.Phone}}</p>
<p>物业邮箱：{{.ContactEmail}}</p>]`
	TemplatePropertyZhCN = `html[
<p>物業名稱：{{.Name}}</p>
<p>物業地址：{{.Address}}</p>
<p>物業電話：{{.Phone}}</p>
<p>物業信箱：{{.ContactEmail}}</p>]`
	TemplatePropertyEn = `html[
<p>Property Name: {{.Name}}</p>
<p>Property Address: {{.Address}}</p>
<p>Property Phone: {{.Phone}}</p>
<p>Property Email: {{.ContactEmail}}</p>]`
	TemplatePropertyJa = `html[
<p>施設名：{{.Name}}</p>
<p>施設住所：{{.Address}}</p>
<p>施設電話：{{.Phone}}</p>
<p>施設メール：{{.ContactEmail}}</p>]`
	TemplatePropertyKo = `html[
<p>물업 이름：{{.Name}}</p>
<p>물업 주소：{{.Address}}</p>
<p>물업 전화：{{.Phone}}</p>
<p>물업 이메일：{{.ContactEmail}}</p>]`
)

// 餐厅、按摩、接送机预定模板
const (
	TemplateSceneZh = `html[
<p>场景：{{.Scene}}</p>
<p>服务：{{.Service}}</p>
<p>预订人：{{.BookingName}}</p>
<p>服务时间：{{.BookDatetime}}</p>
<p>订单编号：{{.OrderSn}}</p>
<p>下单时间：{{.CreatedAt}}</p>]`

	TemplateSceneZhCN = `html[
<p>場景：{{.Scene}}</p>
<p>服務：{{.Service}}</p>
<p>預訂人：{{.BookingName}}</p>
<p>服務時間：{{.BookDatetime}}</p>
<p>訂單編號：{{.OrderSn}}</p>
<p>下單時間：{{.CreatedAt}}</p>]`

	TemplateSceneEn = `html[
<p>Scenario：{{.Scene}}</p>
<p>Serve：{{.Service}}</p>
<p>Booker：{{.BookingName}}</p>
<p>Service Hours：{{.BookDatetime}}</p>
<p>Order Number：{{.OrderSn}}</p>
<p>Order Time：{{.CreatedAt}}</p>]`

	TemplateSceneJa = `html[
<p>シナリオ：{{.Scene}}</p>
<p>仕える：{{.Service}}</p>
<p>ブッカー：{{.BookingName}}</p>
<p>営業時間：{{.BookDatetime}}</p>
<p>注文番号：{{.OrderSn}}</p>
<p>注文時間：{{.CreatedAt}}</p>]`

	TemplateSceneKo = `html[
<p>대본：{{.Scene}}</p>
<p>제공하다：{{.Service}}</p>
<p>부커：{{.BookingName}}</p>
<p>서비스 시간：{{.BookDatetime}}</p>
<p>주문 번호：{{.OrderSn}}</p>
<p>주문 시간：{{.CreatedAt}}</p>]`
)

var (
	BookingChangeTypeGuest = g.MapStrStr{
		"zh":    "入住人信息变更成功。",
		"zh_CN": "入住人資訊變更成功。",
		"en":    "The guest information has been successfully updated.",
		"ja":    "宿泊者の情報の変更に成功しました。",
		"ko":    "투숙자 정보 변경에 성공했습니다.",
	}
	BookingChangeTypePeople = g.MapStrStr{
		"zh":    "入住人数变更成功。",
		"zh_CN": "入住人數變更成功。",
		"en":    "The number of guests has been successfully updated.",
		"ja":    "宿泊者の人数の変更に成功しました。",
		"ko":    "투숙 인원 수 변경에 성공했습니다.",
	}
	BookingChangeTypeDate = g.MapStrStr{
		"zh":    "日期变更完成，请在新日期办理入住。",
		"zh_CN": "日期變更完成，請在新日期辦理入住。",
		"en":    "The date has been successfully changed. Please check in on the new date.",
		"ja":    "日付の変更が完了しました。新しい日付でチェックインしてください。",
		"ko":    "날짜 변경이 완료되었습니다. 새로운 날짜에 체크인해주세요.",
	}
)
