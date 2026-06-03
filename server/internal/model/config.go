// Package model
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package model

// BasicConfig 基础配置
type BasicConfig struct {
	CaptchaSwitch  int    `json:"basicCaptchaSwitch"`
	CloseText      string `json:"basicCloseText"`
	Copyright      string `json:"basicCopyright"`
	IcpCode        string `json:"basicIcpCode"`
	Logo           string `json:"basicLogo"`
	Name           string `json:"basicName"`
	Domain         string `json:"basicDomain"`
	WsAddr         string `json:"basicWsAddr"`
	RegisterSwitch int    `json:"basicRegisterSwitch"`
	SystemOpen     bool   `json:"basicSystemOpen"`
}

// EmailTemplate 邮件模板
type EmailTemplate struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// EmailConfig 邮箱配置
type EmailConfig struct {
	User         string           `json:"smtpUser"`
	Password     string           `json:"smtpPass"`
	Addr         string           `json:"smtpAddr"`
	Host         string           `json:"smtpHost"`
	Port         int64            `json:"smtpPort"`
	SendName     string           `json:"smtpSendName"`
	AdminMailbox string           `json:"smtpAdminMailbox"`
	MinInterval  int              `json:"smtpMinInterval"`
	MaxIpLimit   int              `json:"smtpMaxIpLimit"`
	CodeExpire   int              `json:"smtpCodeExpire"`
	Template     []*EmailTemplate `json:"smtpTemplate"`
}

// UploadConfig 上传配置
type UploadConfig struct {
	// 通用配置
	Drive     string `json:"uploadDrive"`
	FileSize  int64  `json:"uploadFileSize"`
	FileType  string `json:"uploadFileType"`
	ImageSize int64  `json:"uploadImageSize"`
	ImageType string `json:"uploadImageType"`
	// 本地存储配置
	LocalPath string `json:"uploadLocalPath"`
	// UCloud对象存储配置
	UCloudBucketHost string `json:"uploadUCloudBucketHost"`
	UCloudBucketName string `json:"uploadUCloudBucketName"`
	UCloudEndpoint   string `json:"uploadUCloudEndpoint"`
	UCloudFileHost   string `json:"uploadUCloudFileHost"`
	UCloudPath       string `json:"uploadUCloudPath"`
	UCloudPrivateKey string `json:"uploadUCloudPrivateKey"`
	UCloudPublicKey  string `json:"uploadUCloudPublicKey"`
	// 腾讯云cos配置
	CosSecretId  string `json:"uploadCosSecretId"`
	CosSecretKey string `json:"uploadCosSecretKey"`
	CosBucketURL string `json:"uploadCosBucketURL"`
	CosPath      string `json:"uploadCosPath"`
	// 阿里云oss配置
	OssSecretId  string `json:"uploadOssSecretId"`
	OssSecretKey string `json:"uploadOssSecretKey"`
	OssEndpoint  string `json:"uploadOssEndpoint"`
	OssBucketURL string `json:"uploadOssBucketURL"`
	OssPath      string `json:"uploadOssPath"`
	OssBucket    string `json:"uploadOssBucket"`
	// 七牛云对象存储配置
	QiNiuAccessKey string `json:"uploadQiNiuAccessKey"`
	QiNiuSecretKey string `json:"uploadQiNiuSecretKey"`
	QiNiuDomain    string `json:"uploadQiNiuDomain"`
	QiNiuPath      string `json:"uploadQiNiuPath"`
	QiNiuBucket    string `json:"uploadQiNiuBucket"`
	// minio配置
	MinioAccessKey string `json:"uploadMinioAccessKey"`
	MinioSecretKey string `json:"uploadMinioSecretKey"`
	MinioEndpoint  string `json:"uploadMinioEndpoint"`
	MinioUseSSL    int    `json:"uploadMinioUseSSL"`
	MinioPath      string `json:"uploadMinioPath"`
	MinioBucket    string `json:"uploadMinioBucket"`
	MinioDomain    string `json:"uploadMinioDomain"`
	// google 云存储
	UploadGoogleCredentials string `json:"uploadGoogleCredentials"`
	UploadGoogleBucketName  string `json:"uploadGoogleBucketName"`
	UploadGoogleEndpoint    string `json:"uploadGoogleEndpoint"`
	UploadGooglePath        string `json:"uploadGooglePath"`
	// 缩放图设置
	UploadImageThumb string `json:"uploadImageThumb"`
}

type ThumbSizeConfig struct {
	BigWidth    int `json:"bigWidth"`
	BigHeight   int `json:"bigHeight"`
	MidWidth    int `json:"midWidth"`
	MidHeight   int `json:"midHeight"`
	SmallWidth  int `json:"smallWidth"`
	SmallHeight int `json:"smallHeight"`
}

// GeoConfig 地理配置
type GeoConfig struct {
	GeoAmapWebKey string `json:"geoAmapWebKey"`
}

// SmsTemplate 短信模板
type SmsTemplate struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// SmsConfig 短信配置
type SmsConfig struct {
	// 基础
	SmsDrive             string `json:"smsDrive"`
	NoticeSmsDriveCN     string `json:"noticeSmsDriveCN"`
	NoticeSmsDriveAbroad string `json:"noticeSmsDriveAbroad"`
	SmsMinInterval       int    `json:"smsMinInterval"`
	SmsMaxIpLimit        int    `json:"smsMaxIpLimit"`
	SmsCodeExpire        int    `json:"smsCodeExpire"`
	// 阿里云
	AliYunAccessKeyID     string         `json:"smsAliYunAccessKeyID"`
	AliYunAccessKeySecret string         `json:"smsAliYunAccessKeySecret"`
	AliYunSign            string         `json:"smsAliYunSign"`
	AliYunTemplate        []*SmsTemplate `json:"smsAliYunTemplate"`
	// 腾讯云
	TencentSecretId  string         `json:"smsTencentSecretId"`
	TencentSecretKey string         `json:"smsTencentSecretKey"`
	TencentEndpoint  string         `json:"smsTencentEndpoint"`
	TencentRegion    string         `json:"smsTencentRegion"`
	TencentAppId     string         `json:"smsTencentAppId"`
	TencentSign      string         `json:"smsTencentSign"`
	TencentTemplate  []*SmsTemplate `json:"smsTencentTemplate"`
	// 一企通
	UmsSpCode    string         `json:"smsUmsSpCode"`    //一信通企业编号
	UmsAppKey    string         `json:"smsUmsAppKey"`    // 一信通AppKey
	UmsAppSecret string         `json:"smsUmsAppSecret"` //一信通AppSecret
	UmsTemplate  []*SmsTemplate `json:"smsUmsTemplate"`  //一信通短信模板
	// Twilio
	FromPhone      string         `json:"smsTwilioFromPhone"`
	Username       string         `json:"smsTwilioUsername"`
	Password       string         `json:"smsTwilioPassword"`
	TwilioTemplate []*SmsTemplate `json:"smsTwilioTemplate"`
}

type GeeTestConfig struct {
	Enabled     bool   `json:"geeTestEnabled"`
	CaptchaID   string `json:"geeTestCaptchaID"`
	CaptchaKey  string `json:"geeTestCaptchaKey"`
	ValidateURL string `json:"geeTestValidateURL"`
}

// OrderSmsConfig 短信配置
type OrderSmsConfig struct {
	SmsTemplate []*OrderSmsTemplate `json:"smsTemplate"`
}

// OrderSmsTemplate 短信模板
type OrderSmsTemplate struct {
	SmsDrive     string `json:"smsDrive"`
	Event        string `json:"event"`
	Name         string `json:"name"`
	Content      string `json:"content"`
	IsSendSms    int    `json:"isSendSms"`
	SendSmsPhone string `json:"sendSmsPhone"`
	TemplateCode string `json:"templateCode"`
}

type PayConfig struct {

	// PayCloud 支付
	PayCloudAppID           string `json:"payCloudAppID"`
	PayCloudSubAppId        string `json:"payCloudSubAppId"`
	PayCloudMerchantNo      string `json:"payCloudMerchantNo"`
	PayCloudStoreNo         string `json:"payCloudStoreNo"`
	PayCloudWxMiniStoreNo   string `json:"payCloudWxMiniStoreNo"`
	PayCloudWxMiniAppid     string `json:"payCloudWxMiniAppid"`
	PayCloudPrivateKey      string `json:"payCloudPrivateKey"`
	PayCloudEndpoint        string `json:"payCloudEndpoint"`
	PayCloudNotifyUrl       string `json:"payCloudNotifyUrl"`
	PayCloudRefundNotifyUrl string `json:"payCloudRefundNotifyUrl"`
	PayCloudReturnUrlWeb    string `json:"payCloudReturnUrlWeb"`
	PayCloudReturnUrlApp    string `json:"payCloudReturnUrlApp"`

	// Stripe 支付
	StripeKey        string `json:"payStripeKey"`
	StripeSignKey    string `json:"payStripeSignKey"`
	StripeSuccessURL string `json:"payStripeSuccessURL"`
	StripeCancelURL  string `json:"payStripeCancelURL"`
	StripeSuccessAPP string `json:"payStripeSuccessAPP"`
	StripeCancelAPP  string `json:"payStripeCancelAPP"`

	// Paypal 支付
	PaypalClientID     string `json:"payPaypalClientID"`
	PaypalSecret       string `json:"payPaypalSecret"`
	PaypalReturnUrlApp string `json:"payPaypalReturnUrlApp"`
	PaypalCleanUrlApp  string `json:"payPaypalCleanUrlApp"`
	PaypalReturnApp    string `json:"payPaypalReturnApp"`
	PaypalCleanApp     string `json:"payPaypalCleanApp"`
	PaypalIsProd       bool   `json:"payPaypalIsProd"`

	HotelStayExp int64 `json:"hotelStayExp"`

	PayMiniSignKey   string `json:"payMiniSignKey"`
	PayMiniAppID     string `json:"payMiniAppID"`
	PayMiniPid       string `json:"payMiniPid"`
	PayMiniNotifyUrl string `json:"payMiniNotifyUrl"`

	PaypayApiKey       string `json:"PaypayApiKey"`
	PaypayApiKeySecret string `json:"PaypayApiKeySecret"`
	PaypayMerchantID   string `json:"PaypayMerchantID"`
	PaypayNotifyUrl    string `json:"PaypayNotifyUrl"`

	WxminiPaychannel string `json:"wxminiPaychannel"`

	H5FxDomain string `json:"h5_fx_domain"`
}

type AppConfig struct {
	AppApply          bool     `json:"appApply"           dc:"是否开启APP审核模式"`
	AppOpen           bool     `json:"appOpen"            dc:"是否开启APP维护模式"`
	AppVersion        string   `json:"appversion"         dc:"当前最新的APP版本号"`
	Appforce          bool     `json:"appforce"           dc:"APP如果版本号过低是否强制升级APP"`
	AppVersionContent string   `json:"appversionContent"  dc:"当前最新的APP版本升级内容"`
	DownloadUrl       string   `json:"DownloadUrl"        dc:"ios下载地址"`
	IsUpdate          bool     `json:"isUpdate"           dc:"是否需要更新"`
	AppPrivacy        string   `json:"appPrivacy"         dc:"隐私政策"`
	PayModelList      []string `json:"payModelList"      dc:"支付方式列表"`
}

type AppVersionInfo struct {
	Name string `json:"name"`
	Info struct {
		AppVersion        string   `json:"app_version"`
		AppPrivacy        string   `json:"app_privacy"`
		AppDownload       string   `json:"app_download"`
		AppVersionContent string   `json:"app_version_content"`
		AppOpen           bool     `json:"app_open"`
		AppAudit          bool     `json:"app_audit"`
		AppForcedUpload   bool     `json:"app_forced_upload"`
		AppPayModel       []string `json:"app_pay_model"`
		AppPrivacyZh      string   `json:"app_privacy_zh"`
		AppPrivacyZhCN    string   `json:"app_privacy_zh_CN"`
		AppPrivacyEn      string   `json:"app_privacy_en"`
		AppPrivacyJa      string   `json:"app_privacy_ja"`
		AppPrivacyKo      string   `json:"app_privacy_ko"`
		MiniPayChannel    string   `json:"minipaychannel"`
	} `json:"info"`
}

type YYConfig struct {
	RecommendModel       string  `json:"yyRecommendModel"     dc:"推荐模式"`
	ExchangeRate         float64 `json:"yyExchangeRate"       dc:"积分汇率"`
	MemberRegisterAward  float64 `json:"memberRegisterAward"  dc:"会员邀请注册奖励"`
	ChannelRegisterAward float64 `json:"channelRegisterAward" dc:"渠道邀请注册奖励"`
	MemberBrokerageRate  float64 `json:"memberBrokerageRate"  dc:"会员返佣比例"`
}

type WXShareConfig struct {
	Icon        string `json:"icon"        dc:"推荐模式"`
	ShareUrl    string `json:"shareUrl"    dc:"分享链接"`
	ShareDomain string `json:"shareDomain" dc:"分享主域名"`
	TitleZh     string `json:"titleZh"     dc:"会员邀请注册奖励"`
	TitleKo     string `json:"titleKo"     dc:"会员邀请注册奖励"`
	TitleZhCn   string `json:"titleZhCn"   dc:"会员邀请注册奖励"`
	TitleJa     string `json:"titleJa"     dc:"会员邀请注册奖励"`
	TitleEn     string `json:"titleEn"     dc:"会员邀请注册奖励"`
	ContentZh   string `json:"contentZh"   dc:"会员邀请注册奖励"`
	ContentZhCn string `json:"contentZhCn" dc:"会员邀请注册奖励"`
	ContentEn   string `json:"contentEn"   dc:"会员邀请注册奖励"`
	ContentJa   string `json:"contentJa"   dc:"会员邀请注册奖励"`
	ContentKo   string `json:"contentKo"   dc:"会员邀请注册奖励"`
}

type MemberRegRewardConfig struct {
	IsOpen        int     `json:"isOpen"         dc:"是否开启新人注册奖励"`
	RewardType    string  `json:"rewardType"     dc:"奖励类型"`
	RewardBalance float64 `json:"rewardBalance"  dc:"奖励积分"`
	CouponTypeIds string  `json:"couponTypeIds"  dc:"送优惠券"`
	ThCouponIds   string  `json:"thCouponIds"    dc:"送礼品券"`
}

type InviteNewRewardConfig struct {
	IsInviteRegOpen      int     `json:"isInviteRegOpen"      dc:"是否开启邀请注册奖励"`
	RewardType           string  `json:"rewardType"           dc:"奖励类型"`
	RewardBalance        float64 `json:"rewardBalance"        dc:"奖励积分"`
	CouponTypeIds        string  `json:"couponTypeIds"        dc:"送优惠券"`
	IsInviteRewardOpen   int     `json:"isInviteRewardOpen"   dc:"是否开启邀请人奖励"`
	MemberRewardBalance  float64 `json:"memberRewardBalance"  dc:"普通会员奖励积分"`
	StaffRewardBalance   float64 `json:"staffRewardBalance"   dc:"员工奖励积分"`
	ChannelRewardBalance float64 `json:"channelRewardBalance" dc:"渠道奖励积分"`
}

type MemberIntentionConfig struct {
	LinkPhone     string `json:"linkPhone"           dc:"联系电话"`
	WechatNo      string `json:"wechatNo"            dc:"微信号"`
	WechatCodeImg string `json:"wechatCodeImg"       dc:"微信二维码"`
	LineId        string `json:"lineId"              dc:"Line Id"`
	LineImg       string `json:"lineImg"             dc:"Line图片"`
	IsSendSms     int    `json:"isSendSms"           dc:"是否发送短信"`
	SendSmsPhone  string `json:"sendSmsPhone"        dc:"发送短信的联系人电话"`
}

// LoginConfig 登录配置
type LoginConfig struct {
	RegisterSwitch int     `json:"loginRegisterSwitch"`
	CaptchaSwitch  int     `json:"loginCaptchaSwitch"`
	Avatar         string  `json:"loginAvatar"`
	RoleId         int64   `json:"loginRoleId"`
	DeptId         int64   `json:"loginDeptId"`
	PostIds        []int64 `json:"loginPostIds"`
	Protocol       string  `json:"loginProtocol"`
	Policy         string  `json:"loginPolicy"`
	AutoOpenId     int     `json:"loginAutoOpenId"`
	ForceInvite    int     `json:"loginForceInvite"`
}

// PrinterSettingConfig 打印机设置
type PrinterSettingConfig struct {
	IsOpen     int    `json:"isOpen"      dc:"是否开启打印 1-开启 2-关闭"`
	PrinterIds string `json:"printerIds"  dc:"绑定的打印机"`
	PrinterNum int    `json:"printerNum"  dc:"打印联数"`
}

type PmsPriceConfig struct {
	PricePercent int64 `json:"price_percent" dc:"价格百分比"`
}

type AppDataBoardViewConfig struct {
	TestType          int    `json:"testType"      dc:"允许访问人员"`
	CanTestGroupIds   string `json:"canTestGroupIds"           dc:"允许访问的分组"`
	CanTestMemberIds  string `json:"canTestMemberIds"        dc:"允许访问的会员"`
	MemberPermissions string `json:"memberPermissions"       dc:"会员权限配置，格式: memberId:permission1,permission2|memberId:permission1"`
}

type LanguagePackSetting struct {
	LanguagePackVersion  string `json:"languagePackVersion" dc:"语言包版本"`
	LanguagePackFilePath string `json:"languagePackFilePath" dc:"语言包文件路径"`
}

// CabinetApiConfig 储物柜Api配置
type CabinetApiConfig struct {
	CabinetDomain    string `json:"cabinetDomain"`
	CabinetAppid     string `json:"cabinetAppid"`
	CabinetAppkey    string `json:"cabinetAppkey"`
	CabinetApiSecret string `json:"cabinetApiSecret"`
}

// ToretaApiConfig Toreta Api配置
type ToretaApiConfig struct {
	ToretaDomain string `json:"toretaDomain"`
	ApiToken     string `json:"apiToken"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// XpyunConfig Xpyun配置
type XpyunConfig struct {
	XpyunDomain string `json:"xpyunDomain"`
}

// SystemMessageConfig 系统消息配置
type SystemMessageConfig struct {
	IsMqPushOpen int `json:"isMqPushOpen"`
}

// HotelSettingConfig 酒店配置
type HotelSettingConfig struct {
	HotelOrderCheckoutDay  int `json:"hotelOrderCheckoutDay"`
	HandleCheckoutOrderNum int `json:"handleCheckoutOrderNum"`
}

// ContactCenterConfig 客服中心配置
type ContactCenterConfig struct {
	Japan *ContactInfo `json:"japan" dc:"日本客服中心"`
	China *ContactInfo `json:"china" dc:"中国客服中心"`
}

// ContactInfo 客服信息
type ContactInfo struct {
	Phone    string `json:"phone"    dc:"电话"`
	Email    string `json:"email"    dc:"邮箱"`
	WorkTime string `json:"workTime" dc:"工作时间"`
}

type ProfileButtonConfig struct {
	Data *ProfileButtonInfo `json:"data" dc:"数据"`
}

// ProfileButtonInfo 按钮信息
type ProfileButtonInfo struct {
	Qrcode    int `json:"qrcode"    dc:"会员码开关（1-开启  0-关闭）"`
	Coupon    int `json:"coupon"    dc:"优惠券开关（1-开启  0-关闭）"`
	ThCoupon  int `json:"thCoupon"    dc:"礼品券开关（1-开启  0-关闭）"`
	DataBoard int `json:"dataBoard"    dc:"数据看板开关（1-开启  0-关闭）"`
	Employee  int `json:"employee"    dc:"员工专区开关（1-开启  0-关闭）"`
	Workbench int `json:"workbench"    dc:"工作台开关（1-开启  0-关闭）"`
	Fx        int `json:"fx"    dc:"渠道分销开关（1-开启  0-关闭）"`
	Help      int `json:"help"    dc:"帮助中心开关（1-开启  0-关闭）"`
	Contact   int `json:"contact"    dc:"联系我们开关（1-开启  0-关闭）"`
	Agreement int `json:"agreement"    dc:"用户协议开关（1-开启  0-关闭）"`
	Privacy   int `json:"privacy"    dc:"隐私政策开关（1-开启  0-关闭）"`
}
