package cabinet

import (
	"APT/internal/library/cabinetApi"

	"github.com/gogf/gf/v2/frame/g"
)

type CabinetListReq struct {
	g.Meta `path:"/cabinet/cabinetList" method:"post" tags:"APP_CABINET" summary:"储物柜列表"`
}

type CabinetListRes struct {
	List []*cabinetApi.CabinetListResponseItem `json:"list"   dc:"数据列表"`
}
