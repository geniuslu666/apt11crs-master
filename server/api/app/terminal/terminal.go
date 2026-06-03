package terminal

import (
	"APT/internal/model/input/input_terminal"
	"github.com/gogf/gf/v2/frame/g"
)

type VerifyLogReq struct {
	g.Meta `path:"/verifyLog/list" method:"post" tags:"APP_TERMINAL" summary:"核销记录"`
	input_terminal.VerifyListInp
}

type VerifyLogRes struct {
	List  []*input_terminal.VerifyListModel `json:"list"   dc:"数据列表"`
	Count int                               `json:"count"   dc:"数据总数"`
}

type VerifyLogViewReq struct {
	g.Meta `path:"/verifyLog/view" method:"post" tags:"APP_TERMINAL" summary:"核销记录详情"`
	input_terminal.VerifyLogViewInp
}

type VerifyLogViewRes struct {
	*input_terminal.VerifyLogViewModel
}

type CodeViewReq struct {
	g.Meta `path:"/code/view" method:"post" tags:"APP_TERMINAL" summary:"二维码详情"`
	input_terminal.CodeViewInp
}

type CodeViewRes struct {
	*input_terminal.CodeViewModel
}

type CodeVerifyReq struct {
	g.Meta `path:"/code/verify" method:"post" tags:"APP_TERMINAL" summary:"二维码核销"`
	input_terminal.CodeVerifyInp
}

type CodeVerifyRes struct {
	*input_terminal.CodeVerifyModel
}
