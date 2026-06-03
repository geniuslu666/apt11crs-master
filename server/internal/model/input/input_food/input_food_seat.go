// Package sysin

package input_food

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodSeatUpdateFields 修改订餐-座位表字段过滤
type FoodSeatUpdateFields struct {
	SeatName string `json:"seatName" dc:"座位名称"`
	Status   int    `json:"status"   dc:"状态"`
}

// FoodSeatInsertFields 新增订餐-座位表字段过滤
type FoodSeatInsertFields struct {
	SeatName string `json:"seatName" dc:"座位名称"`
	Status   int    `json:"status"   dc:"状态"`
}

// FoodSeatEditInp 修改/新增订餐-座位表
type FoodSeatEditInp struct {
	entity.FoodSeat
}

func (in *FoodSeatEditInp) Filter(ctx context.Context) (err error) {
	// 验证座位名称
	if err := g.Validator().Rules("required").Data(in.SeatName).Messages("座位名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type FoodSeatEditModel struct{}

// FoodSeatDeleteInp 删除订餐-座位表
type FoodSeatDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodSeatDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodSeatDeleteModel struct{}

// FoodSeatViewInp 获取指定订餐-座位表信息
type FoodSeatViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodSeatViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodSeatViewModel struct {
	entity.FoodSeat
}

// FoodSeatListInp 获取订餐-座位表列表
type FoodSeatListInp struct {
	input_form.PageReq
	SeatName string `json:"seatName" dc:"座位名称"`
	Status   int    `json:"status" dc:"状态"`
}

func (in *FoodSeatListInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodSeatListModel struct {
	Id       int         `json:"id"       dc:"id"`
	SeatName string      `json:"seatName" dc:"座位名称"`
	Status   int         `json:"status"   dc:"状态"`
	CreateAt *gtime.Time `json:"createAt" dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" dc:"更新时间"`
}

// FoodSeatStatusInp 更新订餐-座位表状态
type FoodSeatStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *FoodSeatStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type FoodSeatStatusModel struct{}
