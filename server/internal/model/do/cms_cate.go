// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsCate is the golang structure of table cms_cate for DAO operations like Where/Data.
type CmsCate struct {
	g.Meta    `orm:"table:cms_cate, do:true"`
	Id        interface{} //
	CatName   interface{} // 分类名称
	CreatedAt *gtime.Time // 创建时间
}
