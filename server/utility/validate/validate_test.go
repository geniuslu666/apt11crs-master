// Package validate_test
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package validate_test

import (
	"APT/utility/validate"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func TestIsEmail(t *testing.T) {
	b := validate.IsEmail("QTT123456@163.com")
	gtest.Assert(true, b)
}

func TestIsPhoneNumber(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// 测试日本手机号
		testCases := []struct {
			phone    string
			expected bool
			region   validate.PhoneRegion
			format   string
		}{
			// 日本手机号测试
			{"+81-070-12345678", true, validate.PhoneRegionJapan, "+81-070-12345678"},
			{"+81-080-12345678", true, validate.PhoneRegionJapan, "+81-080-12345678"},
			{"+81-090-12345678", true, validate.PhoneRegionJapan, "+81-090-12345678"},
			{"81-070-12345678", true, validate.PhoneRegionJapan, "+81-070-12345678"},
			{"070-1234-5678", true, validate.PhoneRegionJapan, "+81-070-12345678"},
			{"07012345678", true, validate.PhoneRegionJapan, "+81-070-12345678"},
			{"08012345678", true, validate.PhoneRegionJapan, "+81-080-12345678"},
			{"09012345678", true, validate.PhoneRegionJapan, "+81-090-12345678"},
			
			// 中国大陆手机号测试
			{"+86-13812345678", true, validate.PhoneRegionChina, "+86-13812345678"},
			{"86-13812345678", true, validate.PhoneRegionChina, "+86-13812345678"},
			{"13812345678", true, validate.PhoneRegionChina, "+86-13812345678"},
			{"15912345678", true, validate.PhoneRegionChina, "+86-15912345678"},
			{"18812345678", true, validate.PhoneRegionChina, "+86-18812345678"},
			
			// 中国台湾手机号测试
			{"+886-912345678", true, validate.PhoneRegionTaiwan, "+886-912345678"},
			{"+886-0912345678", true, validate.PhoneRegionTaiwan, "+886-0912345678"},
			{"886-912345678", true, validate.PhoneRegionTaiwan, "+886-912345678"},
			{"912345678", true, validate.PhoneRegionTaiwan, "+886-912345678"},
			{"0912345678", true, validate.PhoneRegionTaiwan, "+886-0912345678"},
			
			// 中国香港手机号测试
			{"+852-51234567", true, validate.PhoneRegionHongKong, "+852-51234567"},
			{"+852-61234567", true, validate.PhoneRegionHongKong, "+852-61234567"},
			{"+852-91234567", true, validate.PhoneRegionHongKong, "+852-91234567"},
			{"852-51234567", true, validate.PhoneRegionHongKong, "+852-51234567"},
			{"51234567", true, validate.PhoneRegionHongKong, "+852-51234567"},
			{"61234567", true, validate.PhoneRegionHongKong, "+852-61234567"},
			{"91234567", true, validate.PhoneRegionHongKong, "+852-91234567"},
			
			// 中国澳门手机号测试
			{"+853-61234567", true, validate.PhoneRegionMacao, "+853-61234567"},
			{"853-61234567", true, validate.PhoneRegionMacao, "+853-61234567"},
			{"61234567", true, validate.PhoneRegionMacao, "+853-61234567"},
			
			// 无效手机号测试
			{"", false, "", ""},
			{"123", false, "", ""},
			{"12345678901", false, "", ""},
			{"+81-060-12345678", false, "", ""}, // 日本错误前缀
			{"+86-12345678901", false, "", ""}, // 中国大陆错误格式
			{"+852-41234567", false, "", ""}, // 香港错误前缀
			{"+853-51234567", false, "", ""}, // 澳门错误前缀
			{"abc123", false, "", ""},
		}

		for _, tc := range testCases {
			result := validate.IsPhoneNumber(tc.phone)
			t.AssertEQ(result.IsValid, tc.expected)
			
			if tc.expected {
				t.AssertEQ(result.Region, tc.region)
				t.AssertEQ(result.Format, tc.format)
			}
		}
	})
}

func TestIsPhoneNumberSimple(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// 测试简化版本
		validPhones := []string{
			"+81-070-12345678",
			"13812345678",
			"+886-912345678",
			"51234567",
			"+853-61234567",
		}
		
		invalidPhones := []string{
			"",
			"123",
			"abc123",
			"+81-060-12345678",
			"41234567",
		}
		
		for _, phone := range validPhones {
			t.AssertEQ(validate.IsPhoneNumberSimple(phone), true)
		}
		
		for _, phone := range invalidPhones {
			t.AssertEQ(validate.IsPhoneNumberSimple(phone), false)
		}
	})
}

func TestIsPhoneNumberWithCountryCode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// 测试区号+手机号分开验证
		testCases := []struct {
			countryCode string
			phoneNumber string
			expected    bool
			region      validate.PhoneRegion
			format      string
		}{
			// 日本手机号测试
			{"+81", "07012345678", true, validate.PhoneRegionJapan, "+81-070-12345678"},
			{"81", "08012345678", true, validate.PhoneRegionJapan, "+81-080-12345678"},
			{"+81", "09012345678", true, validate.PhoneRegionJapan, "+81-090-12345678"},
			
			// 中国大陆手机号测试
			{"+86", "13812345678", true, validate.PhoneRegionChina, "+86-13812345678"},
			{"86", "15912345678", true, validate.PhoneRegionChina, "+86-15912345678"},
			{"+86", "18812345678", true, validate.PhoneRegionChina, "+86-18812345678"},
			
			// 中国台湾手机号测试
			{"+886", "912345678", true, validate.PhoneRegionTaiwan, "+886-912345678"},
			{"886", "0912345678", true, validate.PhoneRegionTaiwan, "+886-0912345678"},
			
			// 中国香港手机号测试
			{"+852", "51234567", true, validate.PhoneRegionHongKong, "+852-51234567"},
			{"852", "61234567", true, validate.PhoneRegionHongKong, "+852-61234567"},
			{"+852", "91234567", true, validate.PhoneRegionHongKong, "+852-91234567"},
			
			// 中国澳门手机号测试
			{"+853", "61234567", true, validate.PhoneRegionMacao, "+853-61234567"},
			{"853", "61234567", true, validate.PhoneRegionMacao, "+853-61234567"},
			
			// 无效测试
			{"", "13812345678", false, "", ""},
			{"+86", "", false, "", ""},
			{"+86", "12345678901", false, "", ""}, // 错误格式
			{"+81", "06012345678", false, "", ""}, // 日本错误前缀
			{"+852", "41234567", false, "", ""}, // 香港错误前缀
			{"+999", "13812345678", false, "", ""}, // 不支持的区号
		}

		for _, tc := range testCases {
			result := validate.IsPhoneNumberWithCountryCode(tc.countryCode, tc.phoneNumber)
			t.AssertEQ(result.IsValid, tc.expected)
			
			if tc.expected {
				t.AssertEQ(result.Region, tc.region)
				t.AssertEQ(result.Format, tc.format)
			}
		}
	})
}

func TestIsPhoneNumberWithCountryCodeSimple(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// 测试简化版本
		validCases := []struct {
			countryCode string
			phoneNumber string
		}{
			{"+81", "07012345678"},
			{"+86", "13812345678"},
			{"+886", "912345678"},
			{"+852", "51234567"},
			{"+853", "61234567"},
		}
		
		invalidCases := []struct {
			countryCode string
			phoneNumber string
		}{
			{"", "13812345678"},
			{"+86", ""},
			{"+86", "12345678901"},
			{"+999", "13812345678"},
		}
		
		for _, tc := range validCases {
			t.AssertEQ(validate.IsPhoneNumberWithCountryCodeSimple(tc.countryCode, tc.phoneNumber), true)
		}
		
		for _, tc := range invalidCases {
			t.AssertEQ(validate.IsPhoneNumberWithCountryCodeSimple(tc.countryCode, tc.phoneNumber), false)
		}
	})
}
