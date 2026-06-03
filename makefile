status:
	systemctl  status  APT11_APP.service  | grep Active -B 3
	systemctl  status  APT11_MQThCouponEffect.service  | grep Active -B 3
	systemctl  status  APT11_MQAvailabilities.service  | grep Active -B 3
	systemctl  status  APT11_MQExp.service  | grep Active -B 3
	systemctl  status  APT11_MQOrderExpire.service  | grep Active -B 3
	systemctl  status  APT11_MQOrderStay.service  | grep Active -B 3
	systemctl  status  APT11_MQPlaceOrder.service  | grep Active -B 3
	systemctl  status  APT11_MQRebate.service  | grep Active -B 3
	systemctl  status  APT11_MQOrderAward.service  | grep Active -B 3
	systemctl  status  APT11_SERVER.service  | grep Active -B 3
restart:
	systemctl restart APT11_APP.service
	systemctl restart APT11_MQThCouponEffect.service
	systemctl restart APT11_MQAvailabilities.service
	systemctl restart APT11_MQExp.service
	systemctl restart APT11_MQOrderExpire.service
	systemctl restart APT11_MQOrderStay.service
	systemctl restart APT11_MQPlaceOrder.service
	systemctl restart APT11_MQRebate.service
	systemctl restart APT11_MQOrderAward.service
	systemctl restart APT11_SERVER.service
stop:
	systemctl stop APT11_APP.service
	systemctl stop APT11_MQThCouponEffect.service
	systemctl stop APT11_MQAvailabilities.service
	systemctl stop APT11_MQExp.service
	systemctl stop APT11_MQOrderExpire.service
	systemctl stop APT11_MQOrderStay.service
	systemctl stop APT11_MQPlaceOrder.service
	systemctl stop APT11_MQRebate.service
	systemctl stop APT11_MQOrderAward.service
	systemctl stop APT11_SERVER.service
start:
	systemctl start APT11_APP.service
	systemctl start APT11_MQThCouponEffect.service
	systemctl start APT11_MQAvailabilities.service
	systemctl start APT11_MQExp.service
	systemctl start APT11_MQOrderExpire.service
	systemctl start APT11_MQOrderStay.service
	systemctl start APT11_MQPlaceOrder.service
	systemctl start APT11_MQRebate.service
	systemctl start APT11_MQOrderAward.service
	systemctl start APT11_SERVER.service