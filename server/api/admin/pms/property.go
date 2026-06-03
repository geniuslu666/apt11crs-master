package pms

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"

	"github.com/gogf/gf/v2/frame/g"
)

type PropertyLanguageParams struct {
	Name         input_language.LanguageModel `json:"name"          v:"required-if:type,Name"         dc:"多语言物业名称"`
	Description  input_language.LanguageModel `json:"description"   v:"required-if:type,Description"  dc:"多语言物业描述"`
	Address      input_language.LanguageModel `json:"address"       v:"required-if:type,Address"      dc:"多语言地址信息"`
	TagList      input_language.LanguageModel `json:"tag_list"      v:"required-if:type,TagList"      dc:"多语言物业标签"`
	RoomDes      input_language.LanguageModel `json:"room_des"      v:"required-if:type,RoomDes"      dc:"多语言房间说明"`
	Surroundings input_language.LanguageModel `json:"surroundings"  v:"required-if:type,Surroundings" dc:"多语言周边环境"`
	BusStation   input_language.LanguageModel `json:"bus_station"   v:"required-if:type,BusStation"   dc:"多语言公交站点"`
	Subway       input_language.LanguageModel `json:"subway"        v:"required-if:type,Subway"       dc:"多语言地铁站"`
	CancelPolicy input_language.LanguageModel `json:"cancel_policy" v:"required-if:type,CancelPolicy" dc:"多语言取消政策"`
	RequiredBook input_language.LanguageModel `json:"required_book" v:"required-if:type,RequiredBook" dc:"多语言订房必读"`
	CheckInGuide input_language.LanguageModel `json:"check_in_guide" v:"required-if:type,CheckInGuide" dc:"多语言入住指南"`
}

type PropertyLanguageReq struct {
	g.Meta   `path:"/pms/editPropertyLanguage" method:"post" tags:"ADMIN_PMS" summary:"PMS物业_多语言编辑"`
	Type     string                 `json:"type" v:"required|in:Name,Description,Address,TagList,RoomDes,Surroundings,BusStation,Subway,CancelPolicy,RequiredBook,CheckInGuide#请输入类型|类型错误" dc:"类型"`
	Id       int                    `json:"id" v:"required#请输入物业id" dc:"物业id"`
	Data     PropertyLanguageParams `json:"data" dc:"多语言属性信息"`
	RegionId int                    `json:"region_id" dc:"区域id"`
}
type PropertyLanguageRes struct{}

type PropertyListReq struct {
	g.Meta `path:"/pmsProperty/list" method:"get" tags:"ADMIN_PMS" summary:"PMS物业_列表"`
	input_hotel.PmsPropertyListInp
}

type PropertyListRes struct {
	input_form.PageRes
	List []*input_hotel.PmsPropertyListModel `json:"list"   dc:"数据列表"`
}

type PropertyAllReq struct {
	g.Meta `path:"/pmsProperty/all" method:"get" tags:"ADMIN_PMS" summary:"PMS物业_所有"`
	input_hotel.PmsPropertyAllInp
}

type PropertyAllRes struct {
	input_form.PageRes
	List []*input_hotel.PmsPropertyAllModel `json:"list"   dc:"数据列表"`
}

type PropertyExportReq struct {
	g.Meta `path:"/pmsProperty/export" method:"get" tags:"ADMIN_PMS" summary:"PMS物业_导出"`
	input_hotel.PmsPropertyListInp
}

type PropertyExportRes struct{}

type PropertyViewReq struct {
	g.Meta `path:"/pmsProperty/view" method:"get" tags:"ADMIN_PMS" summary:"PMS物业_详情"`
	input_hotel.PmsPropertyViewInp
}

type PropertyViewRes struct {
	*input_hotel.PmsPropertyViewLannguageModel
}

type PropertyEditReq struct {
	g.Meta `path:"/pmsProperty/edit" method:"post" tags:"ADMIN_PMS" summary:"PMS物业_编辑"`
	input_hotel.PmsPropertyEditInp
}

type PropertyEditRes struct{}

type PropertyEditGroupReq struct {
	g.Meta `path:"/pmsProperty/editGroup" method:"post" tags:"ADMIN_PMS" summary:"PMS物业_绑定会员分组"`
	input_hotel.PmsPropertyEditGroupInp
}

type PropertyEditGroupRes struct{}

type PropertyDeleteReq struct {
	g.Meta `path:"/pmsProperty/delete" method:"post" tags:"ADMIN_PMS" summary:"PMS物业_删除"`
	input_hotel.PmsPropertyDeleteInp
}

type PropertyDeleteRes struct{}

type PropertyRecycleReq struct {
	g.Meta `path:"/pmsProperty/recycle" method:"post" tags:"ADMIN_PMS" summary:"PMS物业_恢复"`
	input_hotel.PmsPropertyRecycleInp
}

type PropertyRecycleRes struct{}

type PropertySetGalleryCoverReq struct {
	g.Meta        `path:"/pmsProperty/setGalleryCover" method:"post" tags:"ADMIN_PMS" summary:"PMS物业_设定相册封面"`
	Uid           string `json:"uid" v:"required#请输入物业id" dc:"物业id"`
	GalleryCover  string `json:"gallery_cover" v:"required-without-all:GalleryImages|url#请输入相册封面|相册封面格式错误" dc:"相册封面"`
	GalleryImages string `json:"gallery_images"  v:"required-without-all:GalleryCover|url#请输入相册图片|相册图片格式错误"     description:"相册图片"`
}
type PropertySetGalleryCoverRes struct{}

type PropertyGetGalleryReq struct {
	g.Meta `path:"/pmsProperty/getGallery" method:"post" tags:"ADMIN_PMS" summary:"PMS物业_获取相册封面"`
	Uid    string `json:"uid" v:"required#请输入物业id" dc:"物业id"`
}

type PropertyGetGalleryRes struct {
	GalleryImages string `json:"gallery_images"        orm:"gallery_images"          description:"画廊图片"`
	GalleryCover  string `json:"gallery_cover"         orm:"gallery_cover"           description:"画廊封面"`
}

type PropertySortReq struct {
	g.Meta `path:"/pmsProperty/sortUpdate" method:"post" tags:"ADMIN_PMS" summary:"PMS物业_排序"`
	input_hotel.PmsPropertySortInp
}

type PropertySortRes struct{}

type PropertySwitchSpaCanOrderReq struct {
	g.Meta `path:"/pmsProperty/switchSpaCanOrder" method:"post" tags:"ADMIN_PMS" summary:"PMS物业_更新按摩预定状态"`
	input_hotel.PropertySwitchSpaCanOrderInp
}

type PropertySwitchSpaCanOrderRes struct{}

type PropertyUpdateAccessPassReq struct {
	g.Meta `path:"/pmsProperty/updateAccessPass" method:"post" tags:"ADMIN_PMS" summary:"PMS物业_更新门禁密码"`
	input_hotel.PropertyUpdateAccessPassInp
}

type PropertyUpdateAccessPassRes struct{}
