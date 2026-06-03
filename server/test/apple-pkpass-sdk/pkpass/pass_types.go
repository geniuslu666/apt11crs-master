package pkpass

// PassType 定义通行证类型
type PassType string

const (
	PassTypeBoardingPass PassType = "boardingPass"
	PassTypeCoupon       PassType = "coupon"
	PassTypeEventTicket  PassType = "eventTicket"
	PassTypeGeneric      PassType = "generic"
	PassTypeStoreCard    PassType = "storeCard"
)

// Location 定义位置信息
type Location struct {
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Altitude     float64 `json:"altitude,omitempty"`
	RelevantText string  `json:"relevantText,omitempty"`
	BeaconUUID   string  `json:"beaconUUID,omitempty"`
	Major        int     `json:"major,omitempty"`
	Minor        int     `json:"minor,omitempty"`
}

// AddLocation 添加位置信息
func (p *Pass) AddLocation(location Location) {
	if p.Data["locations"] == nil {
		p.Data["locations"] = []Location{}
	}

	locations := p.Data["locations"].([]Location)
	locations = append(locations, location)
	p.Data["locations"] = locations
}

// SetPassType 设置通行证类型
func (p *Pass) SetPassType(passType PassType) {
	// 确保存在对应的通行证类型结构
	if p.Data[string(passType)] == nil {
		p.Data[string(passType)] = map[string]interface{}{}
	}

	// 设置通行证类型
	p.Data["passType"] = string(passType)
}

// Barcode 定义条形码信息
type Barcode struct {
	Format          string `json:"format"`
	Message         string `json:"message"`
	MessageEncoding string `json:"messageEncoding"`
	AltText         string `json:"altText,omitempty"`
}

// SetBarcode 设置条形码
func (p *Pass) SetBarcode(barcode Barcode) {
	p.Data["barcode"] = barcode
}

// SetBarcodes 设置多个条形码
func (p *Pass) SetBarcodes(barcodes []Barcode) {
	p.Data["barcodes"] = barcodes
}

// WebService 定义 WebService 信息
type WebService struct {
	AuthenticationToken string `json:"authenticationToken"`
	WebServiceURL       string `json:"webServiceURL"`
	RelevantDate        string `json:"relevantDate,omitempty"`
}

// SetWebService 设置 WebService
func (p *Pass) SetWebService(webService WebService) {
	p.Data["webServiceURL"] = webService.WebServiceURL
	p.Data["authenticationToken"] = webService.AuthenticationToken
}
