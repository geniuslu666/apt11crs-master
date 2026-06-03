// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SpaServiceLabel is the golang structure for table spa_service_label.
type SpaServiceLabel struct {
	ServiceId int64 `json:"serviceId" orm:"service_id" description:"服务ID"`
	LabelId   int64 `json:"labelId"   orm:"label_id"   description:"标签ID"`
}
