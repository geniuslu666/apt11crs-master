package pms

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_language"
	"github.com/gogf/gf/v2/frame/g"
)

// LanguageListReq 查询语言字典列表
type LanguageListReq struct {
	g.Meta `path:"/pmsLanguage/list" method:"get" tags:"ADMIN_PMS" summary:"语言字典_列表"`
	input_language.PmsLanguageListInp
}

type LanguageListRes struct {
	input_form.PageRes
	List []*input_language.PmsLanguageListModel `json:"list"   dc:"数据列表"`
}

// LanguageExportReq 导出语言字典列表
type LanguageExportReq struct {
	g.Meta `path:"/pmsLanguage/export" method:"get" tags:"ADMIN_PMS" summary:"语言字典_导出"`
	input_language.PmsLanguageListInp
}

type LanguageExportRes struct{}

// LanguageViewReq 获取语言字典指定信息
type LanguageViewReq struct {
	g.Meta `path:"/pmsLanguage/view" method:"get" tags:"ADMIN_PMS" summary:"语言字典_详情"`
	input_language.PmsLanguageViewInp
}

type LanguageViewRes struct {
	*input_language.PmsLanguageViewModel
}

// LanguageEditReq 修改/新增语言字典
type LanguageEditReq struct {
	g.Meta `path:"/pmsLanguage/edit" method:"post" tags:"ADMIN_PMS" summary:"语言字典_编辑"`
	input_language.PmsLanguageEditInp
}

type LanguageEditRes struct{}

// LanguageDeleteReq 删除语言字典
type LanguageDeleteReq struct {
	g.Meta `path:"/pmsLanguage/delete" method:"post" tags:"ADMIN_PMS" summary:"语言字典_删除"`
	input_language.PmsLanguageDeleteInp
}

type LanguageDeleteRes struct{}
