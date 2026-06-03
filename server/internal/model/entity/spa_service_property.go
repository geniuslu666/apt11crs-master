// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SpaServiceProperty is the golang structure for table spa_service_property.
type SpaServiceProperty struct {
	ServiceId  int64 `json:"serviceId"  orm:"service_id"  description:"服务ID"`
	PropertyId int64 `json:"propertyId" orm:"property_id" description:"物业ID"`
}
