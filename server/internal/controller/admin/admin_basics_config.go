package admin

import (
	"APT/api/admin/basics"
	"APT/internal/consts"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"APT/internal/service"
	"context"
	"fmt"
)

func (c *ControllerBasics) ConfigGet(ctx context.Context, req *basics.ConfigGetReq) (res *basics.ConfigGetRes, err error) {
	res = new(basics.ConfigGetRes)
	res.GetConfigModel, err = service.BasicsConfig().GetConfigByGroup(ctx, &req.GetConfigInp)
	return
}
func (c *ControllerBasics) ConfigUpdate(ctx context.Context, req *basics.ConfigUpdateReq) (res *basics.ConfigUpdateRes, err error) {
	err = service.BasicsConfig().UpdateConfigByGroup(ctx, &req.UpdateConfigInp)
	return
}
func (c *ControllerBasics) ConfigGetCash(ctx context.Context, req *basics.ConfigGetCashReq) (res *basics.ConfigGetCashRes, err error) {
	res = new(basics.ConfigGetCashRes)
	res.GetConfigModel, err = service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "cash"})
	return
}
func (c *ControllerBasics) ConfigTypeSelect(ctx context.Context, req *basics.ConfigTypeSelectReq) (res *basics.ConfigTypeSelectRes, err error) {
	var (
		list []*input_form.Select
	)
	res = new(basics.ConfigTypeSelectRes)
	for _, v := range consts.ConfigTypes {
		list = append(list, &input_form.Select{
			Value: v,
			Name:  v,
			Label: v,
		})
		fmt.Println("--------------***********-------------------------")
		fmt.Println(list)
	}
	res.List = list
	//jsonData1, err := json.Marshal(list)
	//fmt.Println("--------------QAQ---------------------")
	//fmt.Println(string(jsonData1))
	//if err = gvar.New(list).Scan(&res); err != nil {
	//	return
	//}
	//jsonData, err := json.Marshal(res)
	//fmt.Println("--------------AAAAA--------------------")
	//fmt.Println(string(jsonData))
	return
}
func (c *ControllerBasics) SpaSmsConfigUpdate(ctx context.Context, req *basics.SpaSmsConfigUpdateReq) (res *basics.SpaSmsConfigUpdateRes, err error) {
	err = service.BasicsConfig().SpaSmsConfigUpdateReq(ctx, &req.UpdateSpaSmsConfigInp)
	return
}
func (c *ControllerBasics) ConfigUpdateOrderMember(ctx context.Context, req *basics.ConfigUpdateOrderMemberReq) (res *basics.ConfigUpdateOrderMemberRes, err error) {
	err = service.BasicsConfig().UpdateOrderMember(ctx)
	return
}
