package cmd

var (
//	Aladdin = &gcmd.Command{
//		Name:  "aladdin",
//		Brief: "aladdin OfferInfo test",
//		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
//			var (
//				OfferInfoResponse     *aladdinApi.OfferInfoResponse
//				BookingBusResponse    *aladdinApi.BookingBusResponse
//				CancelOrderResponse   *aladdinApi.CancelOrderResponse
//				UserOrderInfoResponse *aladdinApi.UserOrderInfoResponse
//				cmd                   string
//			)
//			g.Log().Warning(ctx, "查询报价")
//			if OfferInfoResponse, err = aladdinApi.NewClient().OfferInfo(ctx, &aladdinApi.OfferInfoParams{
//				Type:               1,
//				ArrivalAirportCode: "KIX",
//				DateTime:           "2024-12-01 12:00:00",
//			}); err != nil {
//				return
//			}
//			g.Log().Warning(ctx, OfferInfoResponse)
//			g.Log().Warning(ctx, "下单")
//			if _, err = fmt.Scan(&cmd); err != nil {
//				return
//			}
//			if cmd == "N" || cmd == "n" || cmd == "no" || cmd == "NO" {
//				return
//			}
//			cmd = ""
//			if BookingBusResponse, err = aladdinApi.NewClient().BookingBus(ctx, &aladdinApi.BookingBusParams{
//				Currency:             "JPY",
//				Type:                 1,
//				Session:              OfferInfoResponse.Data.Session,
//				DepartureAirportCode: "KIX",
//				DepartureTerminal:    "T1",
//				ArrivingTerminal:     "T1",
//				ArrivalAirportCode:   "KIX",
//				Contacts:             "高杨",
//				CountryCode:          "86",
//				ContactsPhone:        "15605284028",
//				ContactsEmail:        "1256005331@qq.com",
//				DateTime:             "2024-12-01 08:00:00",
//				PersonPrice:          220000,
//				PersonNum:            1,
//				OrderPrice:           220000,
//				Hotel:                "住一新今宮2号店",
//				Address:              "3-chōme-8-7 Ebisunishi, Naniwa Ward, Osaka, 556-0003日本",
//			}); err != nil {
//				return
//			}
//			g.Log().Warning(ctx, BookingBusResponse)
//			g.Log().Warning(ctx, "查询订单信息")
//			if _, err = fmt.Scan(&cmd); err != nil {
//				return
//			}
//			if cmd == "N" || cmd == "n" || cmd == "no" || cmd == "NO" {
//				return
//			}
//			cmd = ""
//
//			if UserOrderInfoResponse, err = aladdinApi.NewClient().UserOrderInfo(ctx, &aladdinApi.UserOrderInfoParams{
//				OrderId: BookingBusResponse.Data.OrderId,
//			}); err != nil {
//				return
//			}
//			g.Log().Warning(ctx, UserOrderInfoResponse)
//			g.Log().Warning(ctx, "查询取消预约")
//			if _, err = fmt.Scan(&cmd); err != nil {
//				return
//			}
//			if cmd == "N" || cmd == "n" || cmd == "no" || cmd == "NO" {
//				return
//			}
//			cmd = ""
//			if CancelOrderResponse, err = aladdinApi.NewClient().CancelOrder(ctx, &aladdinApi.CancelOrderParams{
//				OrderId: BookingBusResponse.Data.OrderId,
//			}); err != nil {
//				return
//			}
//
//			g.Log().Warning(ctx, CancelOrderResponse)
//			return
//		},
//	}
)
