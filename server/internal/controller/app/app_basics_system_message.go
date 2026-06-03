package app

import (
	"context"

	"APT/api/app/basics"
	"APT/internal/library/contexts"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
)

func (c *ControllerBasics) AppMessageList(ctx context.Context, req *basics.AppMessageListReq) (res *basics.AppMessageListRes, err error) {
	var (
		MemberInfo = contexts.GetMemberUser(ctx)
	)
	res = new(basics.AppMessageListRes)
	if res.List, res.Count, err = service.BasicsSystemMessage().AppList(ctx, &input_basics.MessageAppListInp{
		PageReq:  req.PageReq,
		MemberId: MemberInfo.Id,
		Scene:    req.Scene,
		Type:     req.Type,
		UnRead:   false,
	}); err != nil {
		return
	}
	return
}
func (c *ControllerBasics) AppMessageRead(ctx context.Context, req *basics.AppMessageReadReq) (res *basics.AppMessageReadRes, err error) {
	var (
		MemberInfo = contexts.GetMemberUser(ctx)
	)
	err = service.BasicsSystemMessage().Read(ctx, &input_basics.MessageAppReadInp{
		MemberId: MemberInfo.Id,
		Id:       req.Id,
	})
	return
}
func (c *ControllerBasics) AppLatestMessage(ctx context.Context, req *basics.AppLatestMessageReq) (res *basics.AppLatestMessageRes, err error) {
	var (
		MemberInfo = contexts.GetMemberUser(ctx)
	)
	res = new(basics.AppLatestMessageRes)
	res.MessageAppLatestModel, err = service.BasicsSystemMessage().AppLatest(ctx, MemberInfo.Id)
	return
}
