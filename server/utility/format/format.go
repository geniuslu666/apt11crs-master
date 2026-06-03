// Package format
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package format

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"regexp"
	"strconv"
	"time"
)

// Round2String 四舍五入保留小数，默认2位
func Round2String(value float64, args ...interface{}) string {
	var places = 2
	if len(args) > 0 {
		places = gconv.Int(args[0])
	}
	return fmt.Sprintf("%0."+strconv.Itoa(places)+"f", value)
}

// Round2Float64 四舍五入保留小数，默认2位
func Round2Float64(value float64, args ...interface{}) float64 {
	return gconv.Float64(Round2String(value, args...))
}

// FileSize 字节的单位转换 保留两位小数
func FileSize(data int64) string {
	var factor float64 = 1024
	res := float64(data)
	for _, unit := range []string{"", "K", "M", "G", "T", "P"} {
		if res < factor {
			return fmt.Sprintf("%.2f%sB", res, unit)
		}
		res /= factor
	}
	return fmt.Sprintf("%.2f%sB", res, "P")
}

// AgoTime 多久以前
func AgoTime(gt *gtime.Time) string {
	if gt == nil {
		return ""
	}
	n := gtime.Now().Timestamp()
	t := gt.Timestamp()

	var ys int64 = 31536000
	var ds int64 = 86400
	var hs int64 = 3600
	var ms int64 = 60
	var ss int64 = 1
	var rs string

	d := n - t
	switch {
	case d > ys:
		rs = fmt.Sprintf("%d年前", int(d/ys))
	case d > ds:
		rs = fmt.Sprintf("%d天前", int(d/ds))
	case d > hs:
		rs = fmt.Sprintf("%d小时前", int(d/hs))
	case d > ms:
		rs = fmt.Sprintf("%d分钟前", int(d/ms))
	case d > ss:
		rs = fmt.Sprintf("%d秒前", int(d/ss))
	default:
		rs = "刚刚"
	}
	return rs
}

func ChangeTime(timeStr string) (timeChange string) {
	// 解析12小时制时间
	t, err := time.Parse("03:04 PM", timeStr)
	if err != nil {
		return
	}
	// 格式化为 H:i:s
	return t.Format("15:04")
}

func ChangeGMP(inputDate string) (resstr string) {
	// 解析输入日期
	parsedDate, err := time.Parse("2006-01-02", inputDate)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	// 获取当前时间的时分秒
	currentTime := time.Now()
	parsedDate = time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second(), currentTime.Nanosecond(), time.Local)

	// 获取当前时区
	loc, err := time.LoadLocation("Asia/Shanghai") // 你可以根据需要更改时区
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}
	parsedDate = parsedDate.In(loc)

	// 格式化日期
	return parsedDate.Format("Mon Jan 02 2006 15:04:05 GMT -0700 (MST)")
}

func GetHtmlImgSrc(html string) (src string) {
	// 正则表达式匹配 img 标签的 src 属性
	re := regexp.MustCompile(`<img\s+src="([^"]+)"`)
	// 查找匹配的字符串
	matches := re.FindStringSubmatch(html)
	// 提取图片链接
	if len(matches) > 1 {
		src = matches[1]
	} else {
		src = html

	}
	return
}
