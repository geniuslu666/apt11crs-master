package input_basics

import (
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsCollectDeleteInp 删除收藏夹
type PmsCollectDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsCollectDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsCollectDeleteModel struct{}

type PmsCollectAddInp struct {
	MemberId    int    `json:"member_id" dc:"会员ID"`
	CollectType string `json:"collectType" dc:"收藏类型   HOST 、 酒店  REPAST 餐饮"`
	CollectId   int    `json:"collectId" dc:"收藏数据ID"`
}

// PmsCollectListInp 获取收藏夹列表
type PmsCollectListInp struct {
	input_form.PageReq
	CollectType string `json:"collectType" dc:"收藏类型   HOST 、 酒店  REPAST 餐饮"`
	MemberId    int    `json:"memberId" dc:"会员ID"`
}

func (in *PmsCollectListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsCollectListModel struct {
	Id          int         `json:"id"          dc:"id"`
	MemberId    int         `json:"memberId"    dc:"member_id"`
	CollectType string      `json:"collectType" dc:"收藏类型   HOST 、 酒店  REPAST 餐饮"`
	CollectId   int         `json:"collectId"   dc:"收藏数据ID"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"更新时间"`
}
