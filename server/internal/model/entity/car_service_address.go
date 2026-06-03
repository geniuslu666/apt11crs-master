// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CarServiceAddress is the golang structure for table car_service_address.
type CarServiceAddress struct {
	ServiceId int64 `json:"serviceId" orm:"service_id" description:"服务ID"`
	AddressId int64 `json:"addressId" orm:"address_id" description:"地址ID"`
	Type      int   `json:"type"      orm:"type"       description:"1 出发地  2 目的地"`
}
