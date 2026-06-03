package logic_cabinet

import "APT/internal/service"

type sCabinetService struct{}

func NewCabinetService() *sCabinetService {
	return &sCabinetService{}
}

func init() {
	service.RegisterCabinetService(NewCabinetService())
}
