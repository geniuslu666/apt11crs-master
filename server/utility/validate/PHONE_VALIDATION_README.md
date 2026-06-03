# 多国家手机号验证工具

## 概述

这个工具提供了一个通用的手机号验证方法，支持多个国家和地区的手机号格式验证，包括日本、中国大陆、中国台湾、中国香港和中国澳门。

## 支持的格式

### 日本手机号
- **国际区号**: +81
- **格式**: 固定以 070、080、090 开头，11位数字
- **示例**: 
  - `+81-070-12345678`
  - `+81-080-12345678` 
  - `+81-090-12345678`
  - `07012345678` (不带国际区号)

### 中国大陆手机号
- **国际区号**: +86
- **格式**: 11位数字，以1开头，第二位为3-9
- **示例**:
  - `+86-13812345678`
  - `13812345678` (不带国际区号)

### 中国台湾手机号
- **国际区号**: +886
- **格式**: 10位（09开头）或9位（去掉0，9开头）
- **示例**:
  - `+886-912345678` (9位)
  - `+886-0912345678` (10位)
  - `912345678` (不带国际区号)

### 中国香港手机号
- **国际区号**: +852
- **格式**: 8位数字，以5、6、9开头
- **示例**:
  - `+852-51234567`
  - `+852-61234567`
  - `+852-91234567`
  - `51234567` (不带国际区号)

### 中国澳门手机号
- **国际区号**: +853
- **格式**: 固定8位，6开头
- **示例**:
  - `+853-61234567`
  - `61234567` (不带国际区号)

## 使用方法

### 1. 完整手机号验证 (传统方式)

#### 详细验证 (推荐)

```go
import "APT/utility/validate"

// 验证手机号并获取详细信息
result := validate.IsPhoneNumber("+81-070-12345678")

if result.IsValid {
    fmt.Printf("手机号有效！\n")
    fmt.Printf("地区: %s\n", result.Region)
    fmt.Printf("标准格式: %s\n", result.Format)
} else {
    fmt.Printf("手机号无效！\n")
}
```

#### 简化验证

```go
import "APT/utility/validate"

// 只需要知道是否有效
isValid := validate.IsPhoneNumberSimple("13812345678")

if isValid {
    fmt.Println("手机号有效")
} else {
    fmt.Println("手机号无效")
}
```

### 2. 区号+手机号分开验证 (推荐新方式)

#### 详细验证

```go
import "APT/utility/validate"

// 区号和手机号分开传递，更精确的验证
result := validate.IsPhoneNumberWithCountryCode("+86", "13812345678")

if result.IsValid {
    fmt.Printf("手机号有效！\n")
    fmt.Printf("地区: %s\n", result.Region)
    fmt.Printf("标准格式: %s\n", result.Format)
} else {
    fmt.Printf("手机号无效！\n")
}
```

#### 简化验证

```go
import "APT/utility/validate"

// 只需要知道是否有效
isValid := validate.IsPhoneNumberWithCountryCodeSimple("+86", "13812345678")

if isValid {
    fmt.Println("手机号有效")
} else {
    fmt.Println("手机号无效")
}
```

### 优势对比

| 验证方式 | 优势 | 适用场景 |
|---------|------|----------|
| 完整手机号验证 | 使用简单，一个参数 | 用户输入完整手机号 |
| 区号+手机号分开验证 | 精确定位地区，避免误判 | 表单分别收集区号和手机号 |

## 返回结果

### PhoneValidationResult 结构体

```go
type PhoneValidationResult struct {
    IsValid bool        `json:"is_valid"` // 是否有效
    Region  PhoneRegion `json:"region"`   // 所属地区
    Format  string      `json:"format"`   // 标准格式
}
```

### PhoneRegion 枚举

```go
const (
    PhoneRegionJapan     PhoneRegion = "japan"      // 日本
    PhoneRegionChina     PhoneRegion = "china"      // 中国大陆
    PhoneRegionTaiwan    PhoneRegion = "taiwan"     // 中国台湾
    PhoneRegionHongKong  PhoneRegion = "hongkong"   // 中国香港
    PhoneRegionMacao     PhoneRegion = "macao"      // 中国澳门
)
```

## 在业务代码中的应用示例

### 1. API 参数验证

#### 方式一：完整手机号验证
```go
func CreateUser(ctx context.Context, req *CreateUserReq) error {
    // 验证手机号
    phoneResult := validate.IsPhoneNumber(req.Phone)
    if !phoneResult.IsValid {
        return errors.New("手机号格式不正确")
    }
    
    // 可以根据地区做不同处理
    switch phoneResult.Region {
    case validate.PhoneRegionChina:
        // 中国大陆手机号的特殊处理
    case validate.PhoneRegionJapan:
        // 日本手机号的特殊处理
    }
    
    // 保存标准化格式
    user.Phone = phoneResult.Format
    
    return nil
}
```

#### 方式二：区号+手机号分开验证 (推荐)
```go
func CreateUserWithSeparateFields(ctx context.Context, req *CreateUserReq) error {
    // 区号和手机号分开验证，更精确
    phoneResult := validate.IsPhoneNumberWithCountryCode(req.CountryCode, req.PhoneNumber)
    if !phoneResult.IsValid {
        return errors.New("手机号格式不正确")
    }
    
    // 直接知道是哪个地区的手机号
    switch phoneResult.Region {
    case validate.PhoneRegionChina:
        // 中国大陆手机号的特殊处理
    case validate.PhoneRegionJapan:
        // 日本手机号的特殊处理
    }
    
    // 保存标准化格式
    user.Phone = phoneResult.Format
    user.CountryCode = req.CountryCode
    
    return nil
}
```

### 2. 表单验证

```go
func ValidateUserForm(form *UserForm) []string {
    var errors []string
    
    if !validate.IsPhoneNumberSimple(form.Phone) {
        errors = append(errors, "手机号格式不正确")
    }
    
    return errors
}
```

## 特性

- ✅ 支持多种输入格式（带/不带国际区号、带/不带连字符）
- ✅ 自动清理输入（移除空格和连字符）
- ✅ 返回标准化格式
- ✅ 识别手机号所属地区
- ✅ 提供简化和详细两种验证方式
- ✅ 完整的测试覆盖

## 注意事项

1. 输入的手机号会自动清理空格和连字符
2. 返回的标准格式统一使用 `+区号-号码` 的格式
3. 验证严格按照各地区的官方规范进行
4. 建议在生产环境中使用详细验证方法，以便获取更多信息用于后续处理
