package basics

import (
	"APT/internal/model/input/input_basics"

	"github.com/gogf/gf/v2/frame/g"
)

type TestNavListReq struct {
	g.Meta `path:"/home/testNavList" method:"post" tags:"APP_BASICS" summary:"[测试]获取测试导航列表"`
}

type TestNavListRes struct {
	List []*input_basics.PmsTestNavApiListModel `json:"list"   dc:"数据列表"`
}
