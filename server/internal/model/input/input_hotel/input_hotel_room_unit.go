// Package sysin

package input_hotel

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/util/gmeta"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsRoomUnitUpdateFields 修改房间字段过滤
type PmsRoomUnitUpdateFields struct {
	Uid    string `json:"uid"    dc:"第三方系统的ID"`
	RtUid  string `json:"rtUid"  dc:"房型ID"`
	RoomNo string `json:"roomNo" dc:"房间号"`
}

// PmsRoomUnitInsertFields 新增房间字段过滤
type PmsRoomUnitInsertFields struct {
	Uid    string `json:"uid"    dc:"第三方系统的ID"`
	RtUid  string `json:"rtUid"  dc:"房型ID"`
	RoomNo string `json:"roomNo" dc:"房间号"`
}

// PmsRoomUnitEditInp 修改/新增房间
type PmsRoomUnitEditInp struct {
	entity.PmsRoomUnit
}

func (in *PmsRoomUnitEditInp) Filter(ctx context.Context) (err error) {
	// 验证第三方系统的ID
	if err := g.Validator().Rules("required").Data(in.Uid).Messages("第三方系统的ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证房型ID
	if err := g.Validator().Rules("required").Data(in.RtUid).Messages("房型ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type PmsRoomUnitEditModel struct{}

// PmsRoomUnitDeleteInp 删除房间
type PmsRoomUnitDeleteInp struct {
	Id interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *PmsRoomUnitDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsRoomUnitDeleteModel struct{}

// PmsRoomUnitViewInp 获取指定房间信息
type PmsRoomUnitViewInp struct {
	Id int `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *PmsRoomUnitViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsRoomUnitViewModel struct {
	entity.PmsRoomUnit
}

// PmsRoomUnitListInp 获取房间列表
type PmsRoomUnitListInp struct {
	input_form.PageReq
	RoomNo          string        `json:"roomNo"          dc:"房间号"`
	CreatedAt       []*gtime.Time `json:"createdAt"       dc:"创建时间"`
	PmsRoomTypeName string        `json:"pmsRoomTypeName" dc:"房型名称"`
	RtUid           string        `json:"rt_uid" dc:"房型ID"`
	Puid            string        `json:"puid" dc:"物业ID"`
}

func (in *PmsRoomUnitListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsRoomUnitListModel struct {
	Id             int         `json:"id"              dc:"主键"`
	Uid            string      `json:"uid"              dc:"主键"`
	RtUid          string      `json:"rtUid"           dc:"房型ID"`
	RoomNo         string      `json:"roomNo"          dc:"房间号"`
	CreatedAt      *gtime.Time `json:"createdAt"       dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"       dc:"更新时间"`
	RoomTypeDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_room_type"`
		*entity.PmsRoomType
	} `json:"roomTypeDetail" orm:"with:uid=rt_uid"  dc:"房型信息"`
}

// PmsRoomUnitExportModel 导出房间
type PmsRoomUnitExportModel struct {
	Uid             string      `json:"uid"             dc:"第三方系统的ID"`
	Id              int         `json:"id"              dc:"主键"`
	RtUid           string      `json:"rtUid"           dc:"房型ID"`
	RoomNo          string      `json:"roomNo"          dc:"房间号"`
	CreatedAt       *gtime.Time `json:"createdAt"       dc:"创建时间"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       dc:"更新时间"`
	PmsRoomTypeName string      `json:"pmsRoomTypeName" dc:"房型名称"`
}
