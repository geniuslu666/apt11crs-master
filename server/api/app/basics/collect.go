package basics

import (
	"APT/api/app/hotel"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type CollectAddReq struct {
	g.Meta      `path:"/pmsCollect/add" method:"post" tags:"APP_BASICS" summary:"[收藏]加入收藏夹"`
	CollectType string `json:"collectType" v:"required#collect_type_unknown" dc:"收藏类型   HOST 、 酒店  REPAST 餐饮"`
	CollectId   int    `json:"collectId" v:"required#collection_data_id_unknown" dc:"收藏数据ID collectType=HOST  传递 酒店详情主键ID collectType=REPAST  传递 餐饮详情主键ID"`
}

type CollectAddRes struct{}

type CollectListReq struct {
	g.Meta `path:"/pmsCollect/list" method:"post" tags:"APP_BASICS" summary:"[收藏]获取收藏夹列表"`
	input_basics.PmsCollectListInp
}

type CollectListRes struct {
	input_form.PageRes
	List []struct {
		*hotel.PropertyOnlineSearchItem
		CollectId int `json:"collectId"`
	} `json:"list"   dc:"数据列表"`
}

type CollectDeleteReq struct {
	g.Meta `path:"/pmsCollect/delete" method:"post" tags:"APP_BASICS" summary:"[收藏]删除收藏夹"`
	input_basics.PmsCollectDeleteInp
}

type CollectDeleteRes struct{}
