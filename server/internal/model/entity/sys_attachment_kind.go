// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysAttachmentKind is the golang structure for table sys_attachment_kind.
type SysAttachmentKind struct {
	Id    int    `json:"id"    orm:"id"    description:""`
	Label string `json:"label" orm:"label" description:"附件分类名称"`
	Key   string `json:"key"   orm:"key"   description:"分类key"`
	Value string `json:"value" orm:"value" description:"分类value"`
	Icon  string `json:"icon"  orm:"icon"  description:"分类图标"`
	Tag   string `json:"tag"   orm:"tag"   description:""`
}
