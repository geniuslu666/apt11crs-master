// Package validate
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package validate

import (
	"context"
	"encoding/json"
	"net"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

// 是否判断

// IsDNSName 是否是域名地址
func IsDNSName(s string) bool {
	DNSName := `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`
	rxDNSName := regexp.MustCompile(DNSName)
	return s != "" && rxDNSName.MatchString(s)
}

// IsHTTPS 是否是https请求
func IsHTTPS(ctx context.Context) bool {
	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		g.Log().Info(ctx, "IsHTTPS ctx not request")
		return false
	}
	return r.TLS != nil || gstr.Equal(r.Header.Get("X-Forwarded-Proto"), "https")
}

// IsIp 是否为ipv4
func IsIp(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IsPublicIp 是否是公网IP
func IsPublicIp(ip string) bool {
	i := net.ParseIP(ip)

	if i.IsLoopback() || i.IsPrivate() || i.IsMulticast() || i.IsUnspecified() || i.IsLinkLocalUnicast() || i.IsLinkLocalMulticast() {
		return false
	}

	if ip4 := i.To4(); ip4 != nil {
		return !i.Equal(net.IPv4bcast)
	}
	return true
}

// IsLocalIPAddr 检测 IP 地址字符串是否是内网地址
func IsLocalIPAddr(ip string) bool {
	if "localhost" == ip {
		return true
	}
	return HasLocalIP(net.ParseIP(ip))
}

// HasLocalIP 检测 IP 地址是否是内网地址
func HasLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}

// IsMobile 是否为手机号码
func IsMobile(mobile string) bool {
	pattern := `^(1[2|3|4|5|6|7|8|9][0-9]\d{4,8})$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(mobile)
}

// IsEmail 是否为邮箱地址
func IsEmail(email string) bool {
	// pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z].){1,4}[a-z]{2,4}$` //匹配电子邮箱
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// IsURL 是否是url地址
func IsURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return false
	}
	URL, err := url.Parse(u)
	if err != nil || URL.Scheme == "" || URL.Host == "" {
		return false
	}
	return true
}

// IsIDCard 是否为身份证
func IsIDCard(idCard string) bool {
	sz := len(idCard)
	if sz != 18 {
		return false
	}
	weight := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	validate := []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
	sum := 0
	for i := 0; i < len(weight); i++ {
		sum += weight[i] * int(byte(idCard[i])-'0')
	}
	m := sum % 11
	return validate[m] == idCard[sz-1]
}

// IsSameDay 是否为同一天
func IsSameDay(t1, t2 int64) bool {
	y1, m1, d1 := time.Unix(t1, 0).Date()
	y2, m2, d2 := time.Unix(t2, 0).Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// IsSameMinute 是否为同一分钟
func IsSameMinute(t1, t2 int64) bool {
	d1 := time.Unix(t1, 0).Format("2006-01-02 15:04")
	d2 := time.Unix(t2, 0).Format("2006-01-02 15:04")
	return d1 == d2
}

// IsMobileVisit 是否为移动端访问
func IsMobileVisit(userAgent string) bool {
	if len(userAgent) == 0 {
		return false
	}

	is := false
	mobileKeywords := []string{"Mobile", "Android", "Silk/", "Kindle", "BlackBerry", "Opera Mini", "Opera Mobi"}
	for i := 0; i < len(mobileKeywords); i++ {
		if strings.Contains(userAgent, mobileKeywords[i]) {
			is = true
			break
		}
	}
	return is
}

// IsWxBrowserVisit 是否为微信访问
func IsWxBrowserVisit(userAgent string) bool {
	if len(userAgent) == 0 {
		return false
	}

	is := false
	userAgent = strings.ToLower(userAgent)
	mobileKeywords := []string{"MicroMessenger"}
	for i := 0; i < len(mobileKeywords); i++ {
		if strings.Contains(userAgent, strings.ToLower(mobileKeywords[i])) {
			is = true
			break
		}
	}
	return is
}

// IsWxMiniProgramVisit 是否为微信小程序访问
func IsWxMiniProgramVisit(userAgent string) bool {
	if len(userAgent) == 0 {
		return false
	}

	is := false
	userAgent = strings.ToLower(userAgent)
	mobileKeywords := []string{"miniProgram"}
	for i := 0; i < len(mobileKeywords); i++ {
		if strings.Contains(userAgent, strings.ToLower(mobileKeywords[i])) {
			is = true
			break
		}
	}
	return is
}

func IsValidJSON(jsonStr string) bool {
	var emptyInterface interface{}
	err := json.Unmarshal([]byte(jsonStr), &emptyInterface)
	return err == nil
}

// PhoneRegion 手机号地区类型
type PhoneRegion string

const (
	PhoneRegionJapan    PhoneRegion = "japan"    // 日本
	PhoneRegionChina    PhoneRegion = "china"    // 中国大陆
	PhoneRegionTaiwan   PhoneRegion = "taiwan"   // 中国台湾
	PhoneRegionHongKong PhoneRegion = "hongkong" // 中国香港
	PhoneRegionMacao    PhoneRegion = "macao"    // 中国澳门
)

// PhoneValidationResult 手机号验证结果
type PhoneValidationResult struct {
	IsValid bool        `json:"is_valid"` // 是否有效
	Region  PhoneRegion `json:"region"`   // 所属地区
	Format  string      `json:"format"`   // 标准格式
}

// IsPhoneNumber 验证手机号是否有效（支持多国家/地区）
// 支持的格式：
// - 日本: +81-070-12345678, +81-080-12345678, +81-090-12345678 或不带国际区号的 070/080/090开头11位
// - 中国大陆: +86-13812345678 或不带国际区号的 1开头11位
// - 中国台湾: +886-912345678, +886-0912345678 或不带国际区号的 09/9开头
// - 中国香港: +852-51234567, +852-61234567, +852-91234567 或不带国际区号的 5/6/9开头8位
// - 中国澳门: +853-61234567 或不带国际区号的 6开头8位
func IsPhoneNumber(phone string) *PhoneValidationResult {
	if phone == "" {
		return &PhoneValidationResult{IsValid: false}
	}

	// 清理输入，移除空格和连字符
	cleanPhone := strings.ReplaceAll(strings.ReplaceAll(phone, " ", ""), "-", "")

	// 检查各个地区的格式
	if result := checkJapanPhone(cleanPhone); result.IsValid {
		return result
	}
	if result := checkChinaPhone(cleanPhone); result.IsValid {
		return result
	}
	if result := checkTaiwanPhone(cleanPhone); result.IsValid {
		return result
	}
	if result := checkHongKongPhone(cleanPhone); result.IsValid {
		return result
	}
	if result := checkMacaoPhone(cleanPhone); result.IsValid {
		return result
	}

	return &PhoneValidationResult{IsValid: false}
}

// checkJapanPhone 检查日本手机号
// 格式: +81-070-12345678 或 07012345678
func checkJapanPhone(phone string) *PhoneValidationResult {
	// 带国际区号的格式
	japanWithCode := `^(\+81|81)(70|80|90)\d{8}$`
	if matched, _ := regexp.MatchString(japanWithCode, phone); matched {
		// 标准化格式
		re := regexp.MustCompile(`^(\+81|81)((70|80|90)\d{8})$`)
		matches := re.FindStringSubmatch(phone)
		if len(matches) >= 3 {
			standardFormat := "+81-" + matches[2][:3] + "-" + matches[2][3:]
			return &PhoneValidationResult{
				IsValid: true,
				Region:  PhoneRegionJapan,
				Format:  standardFormat,
			}
		}
	}

	// 不带国际区号的格式
	japanLocal := `^0(70|80|90)\d{8}$`
	if matched, _ := regexp.MatchString(japanLocal, phone); matched {
		// 标准化格式
		standardFormat := "+81-" + phone[1:4] + "-" + phone[4:]
		return &PhoneValidationResult{
			IsValid: true,
			Region:  PhoneRegionJapan,
			Format:  standardFormat,
		}
	}

	return &PhoneValidationResult{IsValid: false}
}

// checkChinaPhone 检查中国大陆手机号
// 格式: +86-13812345678 或 13812345678
func checkChinaPhone(phone string) *PhoneValidationResult {
	// 带国际区号的格式
	chinaWithCode := `^(\+86|86)1[3-9]\d{9}$`
	if matched, _ := regexp.MatchString(chinaWithCode, phone); matched {
		// 标准化格式
		re := regexp.MustCompile(`^(\+86|86)(1[3-9]\d{9})$`)
		matches := re.FindStringSubmatch(phone)
		if len(matches) >= 3 {
			standardFormat := "+86-" + matches[2]
			return &PhoneValidationResult{
				IsValid: true,
				Region:  PhoneRegionChina,
				Format:  standardFormat,
			}
		}
	}

	// 不带国际区号的格式
	chinaLocal := `^1[3-9]\d{9}$`
	if matched, _ := regexp.MatchString(chinaLocal, phone); matched {
		standardFormat := "+86-" + phone
		return &PhoneValidationResult{
			IsValid: true,
			Region:  PhoneRegionChina,
			Format:  standardFormat,
		}
	}

	return &PhoneValidationResult{IsValid: false}
}

// checkTaiwanPhone 检查中国台湾手机号
// 格式: +886-912345678 (9位) 或 +886-0912345678 (10位)
func checkTaiwanPhone(phone string) *PhoneValidationResult {
	// 带国际区号的格式 - 9位 (去掉0)
	taiwanWithCode9 := `^(\+886|886)9\d{8}$`
	if matched, _ := regexp.MatchString(taiwanWithCode9, phone); matched {
		re := regexp.MustCompile(`^(\+886|886)(9\d{8})$`)
		matches := re.FindStringSubmatch(phone)
		if len(matches) >= 3 {
			standardFormat := "+886-" + matches[2]
			return &PhoneValidationResult{
				IsValid: true,
				Region:  PhoneRegionTaiwan,
				Format:  standardFormat,
			}
		}
	}

	// 带国际区号的格式 - 10位 (保留0)
	taiwanWithCode10 := `^(\+886|886)09\d{8}$`
	if matched, _ := regexp.MatchString(taiwanWithCode10, phone); matched {
		re := regexp.MustCompile(`^(\+886|886)(09\d{8})$`)
		matches := re.FindStringSubmatch(phone)
		if len(matches) >= 3 {
			standardFormat := "+886-" + matches[2]
			return &PhoneValidationResult{
				IsValid: true,
				Region:  PhoneRegionTaiwan,
				Format:  standardFormat,
			}
		}
	}

	// 不带国际区号的格式
	taiwanLocal := `^(09\d{8}|9\d{8})$`
	if matched, _ := regexp.MatchString(taiwanLocal, phone); matched {
		standardFormat := "+886-" + phone
		return &PhoneValidationResult{
			IsValid: true,
			Region:  PhoneRegionTaiwan,
			Format:  standardFormat,
		}
	}

	return &PhoneValidationResult{IsValid: false}
}

// checkHongKongPhone 检查中国香港手机号
// 格式: +852-51234567 (以5、6、9开头的8位数字)
func checkHongKongPhone(phone string) *PhoneValidationResult {
	// 带国际区号的格式
	hkWithCode := `^(\+852|852)[569]\d{7}$`
	if matched, _ := regexp.MatchString(hkWithCode, phone); matched {
		re := regexp.MustCompile(`^(\+852|852)([569]\d{7})$`)
		matches := re.FindStringSubmatch(phone)
		if len(matches) >= 3 {
			standardFormat := "+852-" + matches[2]
			return &PhoneValidationResult{
				IsValid: true,
				Region:  PhoneRegionHongKong,
				Format:  standardFormat,
			}
		}
	}

	// 不带国际区号的格式
	hkLocal := `^[569]\d{7}$`
	if matched, _ := regexp.MatchString(hkLocal, phone); matched {
		standardFormat := "+852-" + phone
		return &PhoneValidationResult{
			IsValid: true,
			Region:  PhoneRegionHongKong,
			Format:  standardFormat,
		}
	}

	return &PhoneValidationResult{IsValid: false}
}

// checkMacaoPhone 检查中国澳门手机号
// 格式: +853-61234567 (6开头的8位数字)
func checkMacaoPhone(phone string) *PhoneValidationResult {
	// 带国际区号的格式
	macaoWithCode := `^(\+853|853)6\d{7}$`
	if matched, _ := regexp.MatchString(macaoWithCode, phone); matched {
		re := regexp.MustCompile(`^(\+853|853)(6\d{7})$`)
		matches := re.FindStringSubmatch(phone)
		if len(matches) >= 3 {
			standardFormat := "+853-" + matches[2]
			return &PhoneValidationResult{
				IsValid: true,
				Region:  PhoneRegionMacao,
				Format:  standardFormat,
			}
		}
	}

	// 不带国际区号的格式
	macaoLocal := `^6\d{7}$`
	if matched, _ := regexp.MatchString(macaoLocal, phone); matched {
		standardFormat := "+853-" + phone
		return &PhoneValidationResult{
			IsValid: true,
			Region:  PhoneRegionMacao,
			Format:  standardFormat,
		}
	}

	return &PhoneValidationResult{IsValid: false}
}

// IsPhoneNumberSimple 简化版手机号验证，只返回是否有效
func IsPhoneNumberSimple(phone string) bool {
	result := IsPhoneNumber(phone)
	return result.IsValid
}

// IsPhoneNumberWithCountryCode 根据区号验证对应地区的手机号
// countryCode: 国际区号，如 "+86", "86", "+81", "81" 等
// phoneNumber: 纯手机号码，如 "13812345678"
func IsPhoneNumberWithCountryCode(countryCode, phoneNumber string) *PhoneValidationResult {
	if countryCode == "" || phoneNumber == "" {
		return &PhoneValidationResult{IsValid: false}
	}

	// 清理输入
	cleanCountryCode := strings.ReplaceAll(strings.ReplaceAll(countryCode, " ", ""), "-", "")
	cleanPhoneNumber := strings.ReplaceAll(strings.ReplaceAll(phoneNumber, " ", ""), "-", "")

	// 标准化区号格式（移除+号进行比较）
	normalizedCode := strings.TrimPrefix(cleanCountryCode, "+")

	// 根据区号选择对应的验证方法
	switch normalizedCode {
	case "81":
		return checkJapanPhoneByNumber(cleanPhoneNumber)
	case "86":
		return checkChinaPhoneByNumber(cleanPhoneNumber)
	case "886":
		return checkTaiwanPhoneByNumber(cleanPhoneNumber)
	case "852":
		return checkHongKongPhoneByNumber(cleanPhoneNumber)
	case "853":
		return checkMacaoPhoneByNumber(cleanPhoneNumber)
	default:
		return &PhoneValidationResult{IsValid: true}
	}
}

// IsPhoneNumberWithCountryCodeSimple 简化版区号+手机号验证
func IsPhoneNumberWithCountryCodeSimple(countryCode, phoneNumber string) bool {
	result := IsPhoneNumberWithCountryCode(countryCode, phoneNumber)
	return result.IsValid
}

// checkJapanPhoneByNumber 根据纯手机号验证日本手机号
func checkJapanPhoneByNumber(phoneNumber string) *PhoneValidationResult {
	// 日本手机号格式：070/080/090开头的11位数字
	// 支持格式：07012345678, 070-1234-5678, 7012345678, 70-1234-5678
	japanPattern := `^(0)?(70|80|90)(?:\d{8}|\d{4}-\d{4})$`
	if matched, _ := regexp.MatchString(japanPattern, phoneNumber); matched {
		// 移除连字符进行标准化处理
		cleanNumber := strings.ReplaceAll(phoneNumber, "-", "")

		// 如果没有前导0，需要添加
		if !strings.HasPrefix(cleanNumber, "0") {
			cleanNumber = "0" + cleanNumber
		}

		standardFormat := "+81-" + cleanNumber[1:4] + "-" + cleanNumber[4:]
		return &PhoneValidationResult{
			IsValid: true,
			Region:  PhoneRegionJapan,
			Format:  standardFormat,
		}
	}
	return &PhoneValidationResult{IsValid: false}
}

// checkChinaPhoneByNumber 根据纯手机号验证中国大陆手机号
func checkChinaPhoneByNumber(phoneNumber string) *PhoneValidationResult {
	// 中国大陆手机号格式：1开头，第二位3-9，共11位
	chinaPattern := `^1[3-9]\d{9}$`
	if matched, _ := regexp.MatchString(chinaPattern, phoneNumber); matched {
		standardFormat := "+86-" + phoneNumber
		return &PhoneValidationResult{
			IsValid: true,
			Region:  PhoneRegionChina,
			Format:  standardFormat,
		}
	}
	return &PhoneValidationResult{IsValid: false}
}

// checkTaiwanPhoneByNumber 根据纯手机号验证中国台湾手机号
func checkTaiwanPhoneByNumber(phoneNumber string) *PhoneValidationResult {
	// 台湾手机号格式：09开头10位 或 9开头9位
	taiwanPattern := `^(09\d{8}|9\d{8})$`
	if matched, _ := regexp.MatchString(taiwanPattern, phoneNumber); matched {
		standardFormat := "+886-" + phoneNumber
		return &PhoneValidationResult{
			IsValid: true,
			Region:  PhoneRegionTaiwan,
			Format:  standardFormat,
		}
	}
	return &PhoneValidationResult{IsValid: false}
}

// checkHongKongPhoneByNumber 根据纯手机号验证中国香港手机号
func checkHongKongPhoneByNumber(phoneNumber string) *PhoneValidationResult {
	// 香港手机号格式：5/6/9开头的8位数字
	hkPattern := `^[569]\d{7}$`
	if matched, _ := regexp.MatchString(hkPattern, phoneNumber); matched {
		standardFormat := "+852-" + phoneNumber
		return &PhoneValidationResult{
			IsValid: true,
			Region:  PhoneRegionHongKong,
			Format:  standardFormat,
		}
	}
	return &PhoneValidationResult{IsValid: false}
}

// checkMacaoPhoneByNumber 根据纯手机号验证中国澳门手机号
func checkMacaoPhoneByNumber(phoneNumber string) *PhoneValidationResult {
	// 澳门手机号格式：6开头的8位数字
	macaoPattern := `^6\d{7}$`
	if matched, _ := regexp.MatchString(macaoPattern, phoneNumber); matched {
		standardFormat := "+853-" + phoneNumber
		return &PhoneValidationResult{
			IsValid: true,
			Region:  PhoneRegionMacao,
			Format:  standardFormat,
		}
	}
	return &PhoneValidationResult{IsValid: false}
}
