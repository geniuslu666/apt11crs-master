# 测试服务器
DEVBINPath = root@8.211.151.143:/www/wwwroot/CRSSERVER
DEVPassword = yeebok88*
# 生产服务器
BINPath = root@8.211.150.169:/www/wwwroot/CRSSERVER
Password = NndqPGCtkd9QyIf4

sync:
	./push.expect $(Password) ../product/bin/APT11 $(BINPath)
syncDev:
	./push.expect $(DEVPassword) ../product/bin/APT11 $(DEVBINPath)
# 一键编译。编译web前端，并将编译后的包移动到服务端对应静态资源路径下，最后编译服务端
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux  GOARCH=amd64  go build -trimpath -ldflags "-s -w" -o ../product/bin/APT11 ./main.go
status:
	systemctl  status  APT11_APP.service  | grep Active -B 3
	systemctl  status  APT11_MQOrderRemind.service  | grep Active -B 3
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
	systemctl restart APT11_MQOrderRemind.service
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
	systemctl stop APT11_MQOrderRemind.service
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
	systemctl start APT11_MQOrderRemind.service
	systemctl start APT11_MQThCouponEffect.service
	systemctl start APT11_MQAvailabilities.service
	systemctl start APT11_MQExp.service
	systemctl start APT11_MQOrderExpire.service
	systemctl start APT11_MQOrderStay.service
	systemctl start APT11_MQPlaceOrder.service
	systemctl start APT11_MQRebate.service
	systemctl start APT11_MQOrderAward.service
	systemctl start APT11_SERVER.service