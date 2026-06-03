// Package convert
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package convert

import (
	"github.com/gogf/gf/v2/os/gtime"
	"math"
	"time"
)

// UniqueSlice 切片去重
func UniqueSlice[K comparable](languages []K) []K {
	result := make([]K, 0, len(languages))
	temp := map[K]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func Remove(sl []interface{}, f func(v1 interface{}) bool) []interface{} {
	for k, v := range sl {
		if f(v) {
			sl[k] = sl[len(sl)-1]
			sl = sl[:len(sl)-1]
			return sl
		}
	}
	return sl
}

func RemoveSlice[K comparable](src []K, sub K) []K {
	for k, v := range src {
		if v == sub {
			copy(src[k:], src[k+1:])
			return src[:len(src)-1]
		}
	}
	return src
}

// DifferenceSlice 比较两个切片，返回他们的差集
// slice1 := []int{1, 2, 3, 4, 5}
// slice2 := []int{4, 5, 6, 7, 8}
// fmt.Println(Difference(slice1, slice2)) // Output: [1 2 3]
func DifferenceSlice[T comparable](s1, s2 []T) []T {
	m := make(map[T]bool)
	for _, item := range s1 {
		m[item] = true
	}

	var diff []T
	for _, item := range s2 {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return diff
}

func Diff(checkIn string, checkOut string, t string) int {
	t1 := gtime.NewFromStr(checkIn)
	t2 := gtime.NewFromStr(checkOut)
	diff := math.Abs(float64(t1.Timestamp() - t2.Timestamp()))

	switch t {
	case "year":
		if t1.After(t2) {
			t1, t2 = t2, t1 // 交换时间点，确保 t1 在 t2 之前
		}
		// 计算月份差
		monthsDiff := 0
		t1 = t1.AddDate(1, 0, 0) // 增加一个月
		for t1.Before(t2) || t1.Equal(t2) {
			t1 = t1.AddDate(1, 0, 0) // 增加一个月
			monthsDiff++
		}
		return monthsDiff
	case "month":
		if t1.After(t2) {
			t1, t2 = t2, t1 // 交换时间点，确保 t1 在 t2 之前
		}
		// 计算月份差
		monthsDiff := 0
		t1 = t1.AddDate(0, 1, 0) // 增加一个月
		for t1.Before(t2) || t1.Equal(t2) {
			t1 = t1.AddDate(0, 1, 0) // 增加一个月
			monthsDiff++
		}
		return monthsDiff
	case "days":
		return int(diff / 86400) // 3600 * 24
	case "hours":
		return int(diff / 3600 * 24)
	case "min":
		return int(diff / 60)
	case "minutes":
		return int(diff / 60)
	default: //默认秒
		return int(diff)
	}
}

func WeekStartDate() (startDate string) {
	// 获取当前时间
	now := time.Now()

	// 获取当前时间的星期几
	weekday := now.Weekday()

	// 计算距离周一的天数
	daysSinceMonday := int(weekday) - int(time.Sunday)
	if daysSinceMonday != 0 {
		// 如果今天不是周一，则计算本周的开始日期
		startOfWeek := now.AddDate(0, 0, -daysSinceMonday)
		startDate = startOfWeek.Format("2006-01-02")
	} else {
		// 如果今天是周一，直接输出当前日期
		startDate = now.Format("2006-01-02")
	}
	return
}

type DateRange struct {
	StartDate *gtime.Time
	EndDate   *gtime.Time
}

func GenerateDateRanges(startDateStr, endDateStr string, intervalDays int) ([]*DateRange, error) {
	var (
		ranges       []*DateRange
		startDate    *gtime.Time
		endDate      *gtime.Time
		currentStart *gtime.Time
		currentEnd   *gtime.Time
		err          error
	)

	// 解析开始和结束日期
	startDate, err = gtime.StrToTime(startDateStr)
	if err != nil {
		return nil, err
	}

	endDate, err = gtime.StrToTime(endDateStr)
	if err != nil {
		return nil, err
	}
	// 如果起始日期等于结束日期，直接创建一个区间并返回
	if startDate.Equal(endDate) {
		ranges = append(ranges, &DateRange{
			StartDate: startDate,
			EndDate:   endDate,
		})
		return ranges, nil
	}
	// 初始化 currentStart 为 startDate
	currentStart = startDate

	// 生成日期区间数组
	for currentStart.Before(endDate) {
		// 计算当前区间的 EndDate
		currentEnd = currentStart.AddDate(0, 0, intervalDays)

		// 如果当前计算的 EndDate 超过了最终的 endDate，则将其设置为 endDate
		if currentEnd.After(endDate) {
			currentEnd = endDate
		}

		// 将区间添加到数组
		ranges = append(ranges, &DateRange{
			StartDate: currentStart,
			EndDate:   currentEnd,
		})

		// 下一个区间的开始时间为当前的结束时间的下一天
		currentStart = currentEnd.AddDate(0, 0, 1)
	}

	return ranges, nil
}

// IntersectDays 接受两个日期区间的边界，返回它们的交集开始和结束日期之间的数组
func IntersectDays(start1, end1, start2, end2 string) []string {
	dateFunc := func(s string) time.Time {
		t, err := time.Parse("2006-01-02", s)
		if err != nil {
			panic(err)
		}
		return t
	}

	startTime1 := dateFunc(start1)
	endTime1 := dateFunc(end1)
	startTime2 := dateFunc(start2)
	endTime2 := dateFunc(end2)

	// 找出最晚的开始时间和最早的结束时间
	lateStart := startTime1
	if startTime2.After(lateStart) {
		lateStart = startTime2
	}

	earlyEnd := endTime1
	if endTime2.Before(earlyEnd) {
		earlyEnd = endTime2
	}

	if lateStart.After(earlyEnd) {
		return []string{} // 如果没有交集，返回空切片
	}

	// 创建包含交集区间内每一天的切片
	var days []string
	for curr := lateStart; curr.Before(earlyEnd.AddDate(0, 0, 1)); curr = curr.AddDate(0, 0, 1) {
		days = append(days, curr.Format("2006-01-02"))
	}
	return days
}
