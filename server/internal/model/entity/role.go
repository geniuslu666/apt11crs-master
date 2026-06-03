// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Role is the golang structure for table role.
type Role struct {
	Id      uint   `json:"id"      orm:"id"       description:""`
	Name    string `json:"name"    orm:"name"     description:"角色名称"`
	Method  string `json:"method"  orm:"method"   description:"允许的HTTP方法，*代表全部"`
	Path    string `json:"path"    orm:"path"     description:"允许的访客路径，*代表全部"`
	IsSuper uint   `json:"isSuper" orm:"is_super" description:"是否为超管，未启用"`
}
