// Package validate
// 手机号验证使用示例
package validate

import (
	"fmt"
)

// ExamplePhoneValidation 演示如何使用手机号验证功能
func ExamplePhoneValidation() {
	// 测试各种格式的手机号
	testPhones := []string{
		// 日本手机号
		"+81-070-12345678",
		"08012345678",
		
		// 中国大陆手机号
		"+86-13812345678",
		"15912345678",
		
		// 中国台湾手机号
		"+886-912345678",
		"0912345678",
		
		// 中国香港手机号
		"+852-51234567",
		"61234567",
		
		// 中国澳门手机号
		"+853-61234567",
		"61234567",
		
		// 无效手机号
		"123456",
		"abc123",
	}

	fmt.Println("=== 手机号验证示例 ===")
	
	for _, phone := range testPhones {
		// 使用详细验证方法
		result := IsPhoneNumber(phone)
		
		fmt.Printf("手机号: %-20s | 有效: %-5t", phone, result.IsValid)
		
		if result.IsValid {
			fmt.Printf(" | 地区: %-10s | 标准格式: %s", result.Region, result.Format)
		}
		
		fmt.Println()
	}
	
	fmt.Println("\n=== 简化验证示例 ===")
	
	// 使用简化验证方法
	for _, phone := range testPhones[:5] { // 只测试前5个
		isValid := IsPhoneNumberSimple(phone)
		fmt.Printf("手机号: %-20s | 有效: %t\n", phone, isValid)
	}
	
	fmt.Println("\n=== 区号+手机号分开验证示例 ===")
	
	// 测试区号和手机号分开验证
	separateTestCases := []struct {
		countryCode string
		phoneNumber string
		description string
	}{
		{"+86", "13812345678", "中国大陆"},
		{"+81", "07012345678", "日本"},
		{"+886", "912345678", "中国台湾"},
		{"+852", "51234567", "中国香港"},
		{"+853", "61234567", "中国澳门"},
		{"+86", "12345678901", "无效-中国大陆错误格式"},
		{"+999", "13812345678", "无效-不支持的区号"},
	}
	
	for _, tc := range separateTestCases {
		result := IsPhoneNumberWithCountryCode(tc.countryCode, tc.phoneNumber)
		
		fmt.Printf("区号: %-5s | 手机号: %-12s | 有效: %-5t", tc.countryCode, tc.phoneNumber, result.IsValid)
		
		if result.IsValid {
			fmt.Printf(" | 地区: %-10s | 标准格式: %s", result.Region, result.Format)
		}
		
		fmt.Printf(" | 说明: %s\n", tc.description)
	}
	
	fmt.Println("\n=== 区号+手机号简化验证示例 ===")
	
	// 使用简化版本
	for _, tc := range separateTestCases[:5] { // 只测试前5个有效的
		isValid := IsPhoneNumberWithCountryCodeSimple(tc.countryCode, tc.phoneNumber)
		fmt.Printf("区号: %-5s | 手机号: %-12s | 有效: %t\n", tc.countryCode, tc.phoneNumber, isValid)
	}
}
