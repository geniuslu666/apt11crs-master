package toretaApi

import (
	"context"
	"encoding/base64"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/util/gconv"

	"APT/internal/model"
)

var (
	baseUrl                  string                                            // 域名
	connections              = "/externals/public/connections.json"            // 店铺一览获取
	restaurantsEndpoint      = "/externals/public/restaurants.json"            // 店铺一览获取
	restaurantDetailEndpoint = "/externals/public/restaurants"                 // 店铺详细信息获取（需要拼接 /:key.json）
	coursesEndpoint          = "/externals/public/restaurants"                 // 课程一览获取（需要拼接 /:restaurant_id/courses.json）
	slotsEndpoint            = "/externals/public/web_reservations/slots.json" // 获取空席时间段
	createReservation        = "/externals/public/web_reservations.json"       // 创建预约
	cancelReservation        = "/externals/public/web_reservations"            // 取消预约
	notificationsEndpoint    = "/externals/public/notifications.json"          // 获取通知日志
	oauth                    = "/oauth2/token"                                 // OAuth相关
	Logger                   = g.Log().Path("logs/SDK/TORETA_API")
)

type ToretaClient struct {
	GHttpClient  *gclient.Client
	ApiToken     string
	AccessToken  string
	ClientId     string
	ClientSecret string
}

// RestaurantsParams 店铺一览获取参数
type RestaurantsParams struct {
	Page int `json:"page,omitempty" dc:"页码，未指定时默认第1页"`
}

// RestaurantsResponse 店铺一览获取响应
type RestaurantsResponse struct {
	Restaurants []RestaurantItem `json:"restaurants" dc:"店铺列表（每页最多100家）"`
	CurrentPage int              `json:"current_page" dc:"当前页码"`
	TotalPage   int              `json:"total_page" dc:"总页数"`
	HasNextPage bool             `json:"has_next_page" dc:"是否有下一页"`
}

// OARestaurantsResponse 店铺一览获取响应
type OARestaurantsResponse struct {
	Restaurants []OARestaurantItem `json:"restaurants" dc:"店铺列表"`
}

type RestaurantItem struct {
	Id        string `json:"id" dc:"店铺 key"`
	PublicKey string `json:"public_key" dc:"TORETA Web 预约时使用的 key"`
	Name      string `json:"name" dc:"店铺名称"`
}

type OARestaurantItem struct {
	Id   string `json:"id" dc:"店铺 key"`
	Name string `json:"name" dc:"店铺名称"`
}

// RestaurantDetailParams 店铺详细信息获取参数
type RestaurantDetailParams struct {
	Key string `json:"key" dc:"店铺 key"`
}

// RestaurantDetailResponse 店铺详细信息获取响应
type RestaurantDetailResponse struct {
	Key                      string `json:"key" dc:"店铺 key"`
	Name                     string `json:"name" dc:"店铺名称"`
	RestaurantSeatsMax       int    `json:"restaurant_seats_max" dc:"最大预约人数"`
	RestaurantSeatsMin       int    `json:"restaurant_seats_min" dc:"最小预约人数"`
	PublicKey                string `json:"public_key" dc:"PublicKey"`
	WebUrl                   string `json:"web_url" dc:"WebUrl"`
	AcceptChildren           bool   `json:"accept_children" dc:"是否接受儿童"`
	AcceptLimitDay           int    `json:"accept_limit_day" dc:"受理截止日期（单位：天）"`
	AcceptLimitHour          int    `json:"accept_limit_hour" dc:"受理截止时间（单位：秒）"`
	AllowedReservationDays   int    `json:"allowed_reservation_days" dc:"受理期间（单位：天）"`
	CourseRequired           bool   `json:"course_required" dc:"是否需要课程"`
	DefaultReservationPeriod int    `json:"default_reservation_period" dc:"默认预约时长（单位：秒）"`
	ShowHeaderMessage        bool   `json:"show_header_message" dc:"是否显示头部消息"`
	HeaderMessage            string `json:"header_message" dc:"头部消息"`
	NoteMessage              string `json:"note_message" dc:"注意消息"`
	CourseMessage            string `json:"course_message" dc:"课程消息"`
	MessageSuccess           string `json:"message_success" dc:"成功消息"`
	ShowConfirm              bool   `json:"show_confirm" dc:"是否显示确认消息"`
	SelectLanguageEnabled    bool   `json:"select_language_enabled" dc:"是否启用语言选择"`
	V3Visible                bool   `json:"v3_visible" dc:"是否启用 v3"`
	CancelEnabled            bool   `json:"cancel_enabled" dc:"是否启用取消"`
	CancelLimitDay           int    `json:"cancel_limit_day" dc:"取消预约天数（单位：天）"`
	CancelLimitHour          int    `json:"cancel_limit_hour" dc:"取消预约小时数（单位：秒）"`
	Timezone                 string `json:"timezone" dc:"时区"`
}

// CoursesParams 课程一览获取参数
type CoursesParams struct {
	RestaurantId string `json:"restaurant_id" dc:"店铺ID"`
}

// CoursesResponse 课程一览获取响应
type CoursesResponse struct {
	Courses []CourseItem `json:"courses" dc:"课程列表"`
}

type CourseItem struct {
	Id              string `json:"id" dc:"课程 key"`
	Name            string `json:"name" dc:"课程名称"`
	Price           int    `json:"price" dc:"价格"`
	StartDate       string `json:"start_date" dc:"刊载开始日期"`
	EndDate         string `json:"end_date" dc:"刊载结束日期"`
	ExcludeDateFrom string `json:"exclude_date_from" dc:"刊载停止开始日期"`
	ExcludeDateTo   string `json:"exclude_date_to" dc:"刊载停止结束日期"`
	CategoryName    string `json:"category_name" dc:"分类名"`
	EnableLunch     bool   `json:"enable_lunch" dc:"是否接受午餐预约"`
	EnableDinner    bool   `json:"enable_dinner" dc:"是否接受晚餐预约"`
}

// SlotsParams 获取空席时间段参数
type SlotsParams struct {
	Id        string `json:"id" dc:"店铺 key"`
	StartDate string `json:"start_date" dc:"期间开始日期，格式：YYYY-MM-DD"`
	EndDate   string `json:"end_date" dc:"期间结束日期，格式：YYYY-MM-DD"`
	Seats     int    `json:"seats" dc:"预约人数"`
}

// SlotsResponse 空席时间段响应
type SlotsResponse struct {
	Slots []SlotDate `json:"slots" dc:"空席的日期和时间段"`
}

type SlotDate struct {
	Date  string     `json:"date" dc:"存在空席的日期"`
	Slots []SlotTime `json:"slots" dc:"时间段"`
}

type SlotTime struct {
	StartTime int          `json:"start_time" dc:"预约开始时间（秒）"`
	Options   []SlotOption `json:"options" dc:"同一开始时间的不同结束时间和限制时间组合"`
}

type SlotOption struct {
	EndTime         int      `json:"end_time" dc:"预约结束时间（秒）"`
	HasLimitTime    bool     `json:"has_limit_time" dc:"是否有限制时间"`
	TableCategories []string `json:"table_categories" dc:"空席席位类型信息列表"`
}

// CreateReservationParams 创建预约参数
type CreateReservationParams struct {
	Id                 string `json:"id" dc:"店铺 key"`
	TargetDate         string `json:"target_date" dc:"指定日期，格式：YYYY-MM-DD"`
	StartTime          int    `json:"start_time" dc:"预约开始时间（秒）"`
	EndTime            int    `json:"end_time" dc:"预约结束时间（秒）"`
	Seats              int    `json:"seats" dc:"总人数"`
	LastName           string `json:"last_name" dc:"预约人姓氏"`
	FirstName          string `json:"first_name" dc:"预约人名字"`
	Email              string `json:"email" dc:"邮箱地址"`
	Phone              string `json:"phone" dc:"电话号码"`
	Note               string `json:"note,omitempty" dc:"备注"`
	SendMailToCustomer bool   `json:"send_mail_to_customer" dc:"是否发送邮件给客户"`
}

// CreateReservationResponse 创建预约响应
type CreateReservationResponse struct {
	ReservationId string `json:"reservation_id" dc:"预约ID"`
	ReservationNo int    `json:"reservation_no" dc:"预约编号"`
}

// CancelReservationParams 取消预约参数
type CancelReservationParams struct {
	ReservationId string `json:"reservation_id" dc:"预约ID"`
	RestaurantId  string `json:"restaurant_id" dc:"店铺ID"`
}

// CancelReservationResponse 取消预约响应
type CancelReservationResponse struct {
	Status  string `json:"status" dc:"取消状态"`
	Message string `json:"message,omitempty" dc:"消息"`
}

// ReservationDetailParams 预约详情参数
type ReservationDetailParams struct {
	ReservationId string `json:"reservation_id" dc:"预约ID"`
}

// ReservationDetailResponse 预约详情响应
type ReservationDetailResponse struct {
	ReservationId string `json:"reservation_id" dc:"预约ID"`
	Seats         int    `json:"seats" dc:"总人数"`
	Status        string `json:"status" dc:"预约状态（0：未来店，1：到店，2：预约取消，3：未到，4：部分到店，5：网页取消，6：已用餐，7：已完成结账，8：已重置）"`
	HasTimeLimit  bool   `json:"has_time_limit" dc:"是否有限制时间"`
	StartAt       int    `json:"start_at" dc:"预约开始时间（秒）"`
	EndAt         int    `json:"end_at" dc:"预约结束时间（秒）"`
	Note          string `json:"note" dc:"备注"`
}

// NotificationsParams 通知日志获取参数
type NotificationsParams struct {
	StartAt int `json:"start_at" dc:"获取开始时间"`
	EndAt   int `json:"end_at" dc:"获取结束时间"`
}

// NotificationsResponse 通知日志获取响应
type NotificationsResponse []NotificationItem

type NotificationItem struct {
	ResourceKey    string `json:"resource_key" dc:"资源密钥"`
	ResourceType   string `json:"resource_type" dc:"资源类型"`
	ResourceAction string `json:"resource_action" dc:"资源操作"`
	RestaurantKey  string `json:"restaurant_key" dc:"店铺密钥"`
	StatusCode     int    `json:"status_code" dc:"通知日志的HTTP status_code"`
}

// OARestaurantBindParams 店铺绑定参数
type OARestaurantBindParams struct {
	Id string `json:"id" dc:"店铺 key"`
}

// OARestaurantBindResponse 店铺绑定响应
type OARestaurantBindResponse struct {
}

// RefreshOATokenParams 刷新Token参数
type RefreshOATokenParams struct {
	RefreshToken string `json:"refresh_token" dc:"刷新令牌"`
}

type RefreshOATokenApiParams struct {
	GrantType    string `json:"grant_type" dc:"授权类型"`
	ClientID     string `json:"client_id" dc:"客户端ID"`
	ClientSecret string `json:"client_secret" dc:"客户端密钥"`
	RefreshToken string `json:"refresh_token" dc:"刷新令牌"`
}

// RefreshOATokenResponse 刷新Token响应
type RefreshOATokenResponse struct {
	AccessToken string `json:"access_token" dc:"访问令牌"`
	TokenType   string `json:"token_type" dc:"令牌类型"`
	ExpiresIn   int    `json:"expires_in" dc:"有效期（秒）"`
}

// ToretaErrorResponse TORETA API 错误响应
type ToretaErrorResponse struct {
	Error ToretaError `json:"error" dc:"错误信息"`
}

type ToretaError struct {
	Status int    `json:"status" dc:"错误状态码"`
	Detail string `json:"detail" dc:"错误详情"`
}

// ToretaOAErrorResponse TORETA OAuth 错误响应
type ToretaOAErrorResponse struct {
	Error            string `json:"error" dc:"错误信息"`
	ErrorDescription string `json:"error_description" dc:"错误描述"`
}

// NewClient 创建新的Toreta客户端
func NewClient(ctx context.Context, config *model.ToretaApiConfig) *ToretaClient {

	baseUrl = config.ToretaDomain

	return &ToretaClient{
		GHttpClient:  g.Client(),
		ApiToken:     config.ApiToken,
		AccessToken:  config.AccessToken,
		ClientId:     config.ClientId,
		ClientSecret: config.ClientSecret,
	}
}

// DoRequest 执行HTTP请求
func (c *ToretaClient) DoRequest(ctx context.Context, method string, url string, needAccessToken bool, paramsReq interface{}) (response *gclient.Response, err error) {
	// 把原始参数转成 map
	paramMap := gconv.Map(paramsReq)
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	// 将 API Token 进行 base64 编码
	Logger.Info(ctx, "API Token：", c.ApiToken)
	encodedToken := base64.StdEncoding.EncodeToString([]byte(c.ApiToken))
	c.GHttpClient.SetHeader("TORETA-PUBLIC-API-TOKEN", encodedToken)

	// 'Authorization' => "Bearer {$accessToken}"
	if needAccessToken && c.AccessToken != "" {
		c.GHttpClient.SetHeader("Authorization", "Bearer "+c.AccessToken)
	}

	// 根据请求方法设置不同的 Content-Type
	if method == "GET" {
		// GET 请求不设置 Content-Type 为 JSON，让 GoFrame 自动将参数添加到 URL 查询参数
		// 这样 GoFrame 会自动处理：url + "?" + params
	} else {
		// POST/PUT/DELETE 等请求设置 JSON Content-Type，参数作为请求体发送
		c.GHttpClient.SetHeader("Content-Type", "application/json")
	}

	if response, err = c.GHttpClient.DoRequest(ctx, method, baseUrl+url, paramMap); err != nil {
		return
	}
	Logger.Info(ctx, "域名：", baseUrl)
	Logger.Info(ctx, "URL：", url)
	Logger.Info(ctx, response.Raw())
	return
}

// DoOARequest 执行OAuth请求
func (c *ToretaClient) DoOARequest(ctx context.Context, method string, url string, paramsReq interface{}) (response *gclient.Response, err error) {
	// 把原始参数转成 map
	paramMap := gconv.Map(paramsReq)
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	// 将 API Token 进行 base64 编码
	// Logger.Info(ctx, "API Token：", c.ApiToken)
	// encodedToken := base64.StdEncoding.EncodeToString([]byte(c.ApiToken))
	// c.GHttpClient.SetHeader("TORETA-PUBLIC-API-TOKEN", encodedToken)

	// 'Authorization' => "Bearer {$accessToken}"
	// if c.AccessToken != "" {
	// 	c.GHttpClient.SetHeader("Authorization", "Bearer "+c.AccessToken)
	// }

	// 根据请求方法设置不同的 Content-Type
	// if method == "GET" {
	// 	// GET 请求不设置 Content-Type 为 JSON，让 GoFrame 自动将参数添加到 URL 查询参数
	// 	// 这样 GoFrame 会自动处理：url + "?" + params
	// } else {
	// 	// POST/PUT/DELETE 等请求设置 JSON Content-Type，参数作为请求体发送
	// 	c.GHttpClient.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	// }
	c.GHttpClient.SetHeader("Content-Type", "application/x-www-form-urlencoded")

	if response, err = c.GHttpClient.DoRequest(ctx, method, baseUrl+url, paramMap); err != nil {
		return
	}
	Logger.Info(ctx, "域名：", baseUrl)
	Logger.Info(ctx, "URL：", url)
	Logger.Info(ctx, response.Raw())
	return
}
