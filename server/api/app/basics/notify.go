package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type NotifyListReq struct {
	g.Meta `path:"/pmsNotify/list" method:"post" tags:"APP_BASICS" summary:"[站内信]获取列表"`
	input_basics.PmsNotifyListInp
}

type NotifyListRes struct {
	input_form.PageRes
	List []*input_basics.PmsNotifyListModel `json:"list"   dc:"数据列表"`
}

type NotifyViewReq struct {
	g.Meta `path:"/pmsNotify/view" method:"post" tags:"APP_BASICS" summary:"[站内信]详情"`
	input_basics.PmsNotifyViewInp
}

type NotifyViewRes struct {
	*input_basics.PmsNotifyViewModel
}

type NotifyEditReq struct {
	g.Meta `path:"/pmsNotify/edit" method:"post" tags:"APP_BASICS" summary:"[站内信]修改/新增"`
	input_basics.PmsNotifyEditInp
}

type NotifyEditRes struct{}

type NotifyDeleteReq struct {
	g.Meta `path:"/pmsNotify/delete" method:"post" tags:"APP_BASICS" summary:"[站内信]删除"`
	input_basics.PmsNotifyDeleteInp
}

type NotifyDeleteRes struct{}
