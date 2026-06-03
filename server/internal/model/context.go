// Package model
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package model

import (
	"APT/internal/model/entity"
	"APT/internal/model/model_gateway"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Context 请求上下文结构
type Context struct {
	Module          string               // 应用模块 admin｜api｜home｜websocket
	AddonName       string               // 插件名称 如果不是插件模块请求，可能为空
	User            *Identity            // 上下文用户信息
	MemberUser      *MemberIdentity      // 上下文用户信息
	TerminalUser    *TerminalIdentity    // 上下文用户信息
	TravelStaffUser *TravelStaffIdentity // 上下文用户信息
	Language        string               // 语言
	Response        *Response            // 请求响应
	GatewayCode     *model_gateway.GatewayCode
	Data            g.Map // 自定kv变量 业务模块根据需要设置，不固定
}

// Identity 通用身份模型
type Identity struct {
	Id          int64       `json:"id"              description:"用户ID"`
	Pid         int64       `json:"pid"             description:"上级ID"`
	DeptId      int64       `json:"deptId"          description:"部门ID"`
	DeptType    string      `json:"deptType"        description:"部门类型"`
	RoleId      int64       `json:"roleId"          description:"角色ID"`
	RoleKey     string      `json:"roleKey"         description:"角色唯一标识符"`
	Username    string      `json:"username"        description:"用户名"`
	RealName    string      `json:"realName"        description:"姓名"`
	Avatar      string      `json:"avatar"          description:"头像"`
	Email       string      `json:"email"           description:"邮箱"`
	Mobile      string      `json:"mobile"          description:"手机号码"`
	App         string      `json:"app"             description:"登录应用"`
	KefuAccount string      `json:"kefuAccount"     description:"客服账号"`
	LoginAt     *gtime.Time `json:"loginAt"         description:"登录时间"`
}

// MemberIdentity 通用身份模型
type MemberIdentity struct {
	*entity.PmsMember
	PmsMemberAuth *entity.PmsMemberAuth
	App           string      `json:"app"             description:"登录应用"`
	LoginAt       *gtime.Time `json:"loginAt"         description:"登录时间"`
	Channel       string      `json:"channel"         description:"渠道"`
	IsFx          bool        `json:"isFx"             description:"是否分销"`
}

// TerminalIdentity 通用身份模型
type TerminalIdentity struct {
	Type         string      `json:"type"              description:"账号类型"`
	Account      string      `json:"account"             description:"账号"`
	TerminalId   int64       `json:"terminalId"          description:"终端ID"`
	MchId        int64       `json:"mchId"          description:"商户ID"`
	StoreId      int64       `json:"storeId"          description:"门店ID"`
	RestaurantId int64       `json:"restaurantId"          description:"餐厅ID"`
	LoginAt      *gtime.Time `json:"loginAt"         description:"登录时间"`
}

// TravelStaffIdentity 通用身份模型
type TravelStaffIdentity struct {
	Id       uint64      `json:"id"            description:""`
	Name     string      `json:"name"          description:"姓名"`
	Mobile   string      `json:"mobile"        description:"电话"`
	Username string      `json:"username"      description:"登录账号"`
	Status   int         `json:"status"        description:"状态（1启用 2禁用）"`
	LoginAt  *gtime.Time `json:"loginAt"       description:"登录时间"`
}
