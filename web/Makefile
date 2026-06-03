# 测试服务器
DEVBINPath = root@8.211.151.143:/www/wwwroot/CRSSERVER/resource/public
DEVPassword = yeebok88*
# 生产服务器
ProBINPath = root@8.211.150.169:/www/wwwroot/CRSSERVER/resource/public
ProPassword = NndqPGCtkd9QyIf4

run:
	yarn dev
build:
	yarn build
sync:
	./push.expect $(ProPassword) ../product/crs $(ProBINPath)
syncDev:
	./push.expect $(DEVPassword) ../product/crs $(DEVBINPath)
