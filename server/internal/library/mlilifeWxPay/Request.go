package mlilifeWxPay

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"reflect"
	"sort"
	"strings"
)

// GetSign 生成签名
func GetSign(ctx context.Context, data interface{}, key string) (sign string, err error) {

	var (
		v              = reflect.ValueOf(data)
		t              = reflect.TypeOf(data)
		stringA        string
		stringSignTemp string
		hash           [16]byte
		params         []string
	)

	// 检查是否是指针，如果是指针则解引用
	if v.Kind() == reflect.Ptr {
		v = v.Elem() // 解引用获取实际值
		t = v.Type() // 获取解引用后的类型
	} else if v.Kind() == reflect.Struct {
		t = v.Type()
	} else {
		err = gerror.New("input data must be a struct or a pointer to a struct")
		return
	}

	// 遍历结构体字段
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		tag := fieldType.Tag.Get("json") // 获取 struct tag 中的 `sign` 标签

		// 忽略空值的字段
		if tag != "" && field.String() != "" && tag != "sign" {
			params = append(params, fmt.Sprintf("%s=%v", tag, field.Interface()))
		}
	}

	// 第一步：按ASCII码从小到大排序
	sort.Strings(params)
	stringA = strings.Join(params, "&")

	// 第二步：拼接API密钥
	stringSignTemp = stringA + "&key=" + key
	g.Log().Debug(ctx, stringSignTemp)
	// 第三步：MD5运算并转为大写
	hash = md5.Sum([]byte(stringSignTemp))
	sign = strings.ToUpper(hex.EncodeToString(hash[:]))
	g.Log().Debug(ctx, sign)
	return
}
