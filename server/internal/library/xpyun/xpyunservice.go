package xpyun

import (
	model2 "APT/internal/library/xpyun/model"
	"APT/internal/service"
	"context"
)

// BASE_URL 国内
//const BASE_URL = "https://open.xpyun.net/api/openapi"

// 国外
const BASE_URL = "https://sg.open.xpyun.net/api/openapi"

// getBaseURL 从配置表中获取 xpyun 域名
func getBaseURL(ctx context.Context) string {
	xpyunConfig, err := service.BasicsConfig().GetXpyunApi(ctx)
	if err != nil || xpyunConfig == nil || xpyunConfig.XpyunDomain == "" {
		// 如果获取配置失败或配置为空，使用默认值
		return BASE_URL
	}
	return xpyunConfig.XpyunDomain
}

func xpyunPostJson(url string, request interface{}) *model2.XPYunResp {
	result := HttpPostJson(url, request)

	return result
}

/**
 * 1.设置打印机语音类型
 * @param ctx - 上下文
 * @param request
 * @return
 */
func XpYunAddPrinters(ctx context.Context, request *model2.AddPrinterRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/addPrinters"
	return xpyunPostJson(url, request)
}

/**
 * 2.设置打印机语音类型
 * @param ctx - 上下文
 * @param request
 * @return
 */
func XpYunSetVoiceType(ctx context.Context, request model2.SetVoiceTypeRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/setVoiceType"
	return xpyunPostJson(url, request)
}

/**
 * 3.打印小票订单
 * @param ctx - 上下文
 * @param restRequest - 打印订单信息
 * @return
 */
func XpYunPrint(ctx context.Context, request *model2.PrintRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/print"

	// 规范化打印内容，替换不兼容的特殊字符
	request.NormalizeContent()

	return xpyunPostJson(url, request)
}

/**
 * 4.打印标签订单
 * @param ctx - 上下文
 * @param restRequest - 打印订单信息
 * @return
 */
func XpYunPrintLabel(ctx context.Context, request *model2.PrintRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/printLabel"

	// 规范化打印内容，替换不兼容的特殊字符
	request.NormalizeContent()

	return xpyunPostJson(url, request)
}

/**
 * 5.批量删除打印机
 * @param ctx - 上下文
 * @param request
 * @return
 */
func XpYunDelPrinters(ctx context.Context, request *model2.DelPrinterRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/delPrinters"

	return xpyunPostJson(url, request)
}

/**
 * 6.修改打印机信息
 * @param ctx - 上下文
 * @param request
 * @return
 */
func XpYunUpdatePrinter(ctx context.Context, request *model2.UpdPrinterRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/updPrinter"

	return xpyunPostJson(url, request)
}

/**
 * 7.清空待打印队列
 * @param ctx - 上下文
 * @param request
 * @return
 */
func XpYunDelPrinterQueue(ctx context.Context, request *model2.ClearPrintOrderRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/delPrinterQueue"

	return xpyunPostJson(url, request)
}

/**
 * 8.查询订单是否打印成功
 * @param ctx - 上下文
 * @param request
 * @return
 */
func XpYunQueryOrderState(ctx context.Context, request *model2.QueryOrderStateRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/queryOrderState"

	return xpyunPostJson(url, request)
}

/**
 * 9.查询打印机某天的订单统计数
 * @param ctx - 上下文
 * @param request
 * @return
 */
func XpYunQueryOrderStatis(ctx context.Context, request *model2.QueryOrderStatisRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/queryOrderStatis"

	return xpyunPostJson(url, request)
}

/**
* 10.查询打印机状态
* 0、离线 1、在线正常 2、在线不正常
 * 备注：异常一般是无纸，离线的判断是打印机与服务器失去联系超过30秒
* @param ctx - 上下文
* @param request
* @return
*/
func XpYunQueryPrinterStatus(ctx context.Context, request *model2.PrinterRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/queryPrinterStatus"

	return xpyunPostJson(url, request)
}

/**
* 11.批量获取指定打印机状态
* 0、离线 1、在线正常 2、在线不正常
 * 备注：异常一般是无纸，离线的判断是打印机与服务器失去联系超过30秒
* @param ctx - 上下文
* @param request
* @return
*/
func XpYunQueryPrintersStatus(ctx context.Context, request *model2.PrinterRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/queryPrintersStatus"

	return xpyunPostJson(url, request)
}

/**
* 12.金额播报
* 发送用户需要播报的语音内容给支持金额播报的芯烨云打印机
* @param ctx - 上下文
* @param request
* @return
 */
func XpYunPlayVoice(ctx context.Context, request *model2.VoiceRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/playVoice"

	return xpyunPostJson(url, request)
}

/**
 * 13.POS指令
 * @param ctx - 上下文
 * @param restRequest
 * @return
 */
func XpYunPos(ctx context.Context, request *model2.PrintRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/pos"

	return xpyunPostJson(url, request)
}

/**
 * 14.钱箱控制
 * @param ctx - 上下文
 * @param restRequest
 * @return
 */
func XpYunControlBox(ctx context.Context, request *model2.PrinterRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/controlBox"

	return xpyunPostJson(url, request)
}

/**
 * 15.扩展语音播报
 * @param ctx - 上下文
 * @param restRequest
 * @return
 */
func XpYunPlayVoiceExt(ctx context.Context, request *model2.VoicePlayMsgRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/playVoiceExt"

	return xpyunPostJson(url, request)
}

/**
 * 16.自定义语音播报
 * @param ctx - 上下文
 * @param restRequest
 * @return
 */
func XpYunPlayCustomVoice(ctx context.Context, request *model2.VoicePlayMsgRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/playCustomVoice"

	return xpyunPostJson(url, request)
}

/**
 * 17.店铺 LOGO 上传
 * @param ctx - 上下文
 * @param restRequest
 * @return
 */
func XpYunUploadLogo(ctx context.Context, request *model2.UploadLogoRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/uploadLogo"

	return xpyunPostJson(url, request)
}

/**
 * 18.店铺 LOGO 删除
 * @param ctx - 上下文
 * @param restRequest
 * @return
 */
func XpYunDelUploadLogo(ctx context.Context, request *model2.PrinterRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/delUploadLogo"

	return xpyunPostJson(url, request)
}

/**
 * 19.获取打印机基本信息
 * @param ctx - 上下文
 * @param restRequest
 * @return
 */
func XpYunPrinterInfo(ctx context.Context, request *model2.PrinterRequest) *model2.XPYunResp {
	baseURL := getBaseURL(ctx)
	url := baseURL + "/xprinter/printerInfo"

	return xpyunPostJson(url, request)
}
