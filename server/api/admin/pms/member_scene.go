package pms

import (
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type MemberSceneListReq struct {
	g.Meta `path:"/pmsMemberScene/list" method:"get" tags:"ADMIN_PMS" summary:"会员等级场景_列表"`
	input_app_member.PmsMemberSceneListInp
}

type MemberSceneListRes struct {
	input_form.PageRes
	List []*input_app_member.PmsMemberSceneListModel `json:"list"   dc:"数据列表"`
}

type MemberSceneViewReq struct {
	g.Meta `path:"/pmsMemberScene/view" method:"get" tags:"ADMIN_PMS" summary:"会员等级场景_详情"`
	input_app_member.PmsMemberSceneViewInp
}

type MemberSceneViewRes struct {
	*input_app_member.PmsMemberSceneViewModel
}

type MemberSceneEditReq struct {
	g.Meta `path:"/pmsMemberScene/edit" method:"post" tags:"ADMIN_PMS" summary:"会员等级场景_修改/新增"`
	input_app_member.PmsMemberSceneEditInp
}

type MemberSceneEditRes struct{}

type MemberSceneDeleteReq struct {
	g.Meta `path:"/pmsMemberScene/delete" method:"post" tags:"ADMIN_PMS" summary:"会员等级场景_删除"`
	input_app_member.PmsMemberSceneDeleteInp
}

type MemberSceneDeleteRes struct{}

type MemberSceneSwitchReq struct {
	g.Meta `path:"/pmsMemberScene/switch" method:"post" tags:"ADMIN_PMS" summary:"会员等级场景_更新"`
	input_app_member.PmsMemberSceneSwitchInp
}

type MemberSceneSwitchRes struct{}
