package logic_app_member

import "APT/internal/service"

type sAppMember struct{}

func NewAppMember() *sAppMember {
	return &sAppMember{}
}

func init() {
	service.RegisterAppMember(NewAppMember())
}
