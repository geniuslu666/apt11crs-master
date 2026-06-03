package consts

var (
	RabbitMQExchangeName                    = "PMS_ORDER"
	RabbitMQExchangeDelayedName             = "PMS_ORDER_DELAYED"
	RabbitMQQueueNameOrderExpire            = "MQ_ORDER_EXPIRE"              // 【队列】订单过期
	RabbitMQQueueNameOrderStay              = "MQ_SYNC_ORDER_STAY"           // 【队列】同步住宿订单
	RabbitMQQueueNameAvailabilities         = "MQ_SYNC_ORDER_AVAILABILITIES" // 【队列】同步库存价格
	RabbitMQQueueNameRebate                 = "MQ_CHECKOUT_REBATE"           // 【队列】计算佣金
	RabbitMQQueueNameExp                    = "MQ_CHECKOUT_EXP"              // 【队列】计算经验
	RabbitMQQueueNamePlaceOrder             = "MQ_PLACE_ORDER"               // 【队列】支付后向AIRHOST下单
	RabbitMQQueueNameThCouponEffect         = "MQ_TH_COUPON_EFFECT"          // 【队列】礼品券生效
	RabbitMQQueueNameOrderAward             = "MQ_ORDER_AWARD"               // 【队列】下单奖励
	RabbitMQQueueNameOrderRemind            = "MQ_ORDER_REMIND"              // 【队列】 订单场景推送消息
	RabbitMQQueueNameOrderExport            = "MQ_ORDER_EXPORT"              // 【队列】 订单导出
	RabbitMQQueueNameEmployeeActivityEffect = "MQ_EMPLOYEE_ACTIVITY_EFFECT"  // 【队列】员工活动生效
	RabbitMQQueueNameSystemMessage          = "MQ_SYSTEM_MESSAGE"            // 【队列】 系统消息
	RabbitMQQueueNameFxChangeOrderPush      = "MQ_FX_CHANGE_ORDER_PUSH"      // 【队列】 分销订单状态变更通知

	RabbitMQGB = "GB"
)
