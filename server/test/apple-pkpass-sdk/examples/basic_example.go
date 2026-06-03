package main

import (
	"APT/test/apple-pkpass-sdk/pkpass"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	// 创建一个新的通行证
	pass := pkpass.NewPass()

	// 设置基本信息
	pass.SetField("passTypeIdentifier", "pass.com.yeebook.hotel")
	pass.SetField("serialNumber", "2023COFFEE001")
	pass.SetField("teamIdentifier", "QU4VS5KTSM")
	pass.SetField("organizationName", "Example Organization")
	pass.SetField("description", "Example Pass")
	pass.SetField("formatVersion", 1)

	// 设置过期时间
	pass.SetField("expirationDate", time.Now().AddDate(1, 0, 0).Format(time.RFC3339))

	// 设置颜色
	pass.SetField("backgroundColor", "rgb(60, 65, 76)")
	pass.SetField("foregroundColor", "rgb(255, 255, 255)")
	pass.SetField("labelColor", "rgb(255, 255, 255)")

	// 设置为活动票类型
	//pass.SetPassType(pkpass.PassTypeEventTicket)

	// 添加主字段
	pass.AddPrimaryField(pkpass.Field{
		Key:   "event",
		Label: "Event",
		Value: "Example Concert",
	})

	// 添加次要字段
	pass.AddSecondaryField(pkpass.Field{
		Key:   "date",
		Label: "Date",
		Value: time.Now().AddDate(0, 1, 0).Format("Jan 2, 2006"),
	})

	// 添加辅助字段
	pass.AddAuxiliaryField(pkpass.Field{
		Key:   "time",
		Label: "Time",
		Value: time.Now().AddDate(0, 1, 0).Format("3:04 PM"),
	})

	// 添加背面字段
	pass.AddBackField(pkpass.Field{
		Key:   "terms",
		Label: "Terms & Conditions",
		Value: "This is a sample pass for demonstration purposes only.",
	})

	// 设置条形码
	pass.SetBarcode(pkpass.Barcode{
		Format:          "PKBarcodeFormatQR",
		Message:         "1234567890",
		MessageEncoding: "iso-8859-1",
		AltText:         "1234567890",
	})

	// 添加位置信息
	pass.AddLocation(pkpass.Location{
		Latitude:     37.33182,
		Longitude:    -122.03118,
		RelevantText: "Apple Park",
	})

	// 加载签名者
	signer, err := pkpass.LoadSigner(
		"/Users/gaoyang/go/src/APT11_CRS/certificates/passcertificate.pem",
		"/Users/gaoyang/go/src/APT11_CRS/certificates/passkey.pem",
		"/Users/gaoyang/go/src/APT11_CRS/certificates/WWDR.pem",
	)
	if err != nil {
		log.Fatalf("加载签名者失败: %v", err)
	}

	// 设置签名者
	pass.SetSigner(signer)

	// 添加图像
	iconData, err := ioutil.ReadFile("/Users/gaoyang/go/src/APT11_CRS/certificates/icon.png")
	if err != nil {
		log.Fatalf("读取图标失败: %v", err)
	}
	pass.AddImage("icon.png", iconData)

	logoData, err := ioutil.ReadFile("/Users/gaoyang/go/src/APT11_CRS/certificates/logo.png")
	if err != nil {
		log.Fatalf("读取logo失败: %v", err)
	}
	pass.AddImage("logo.png", logoData)

	// 保存通行证
	err = pass.Save("example_pass.pkpass")
	if err != nil {
		log.Fatalf("保存通行证失败: %v", err)
	}

	fmt.Println("通行证已成功创建: example_pass.pkpass")
}
