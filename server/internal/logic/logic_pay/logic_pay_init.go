package logic_pay

import "APT/internal/service"

type sPayService struct{}

func NewPayService() *sPayService {
	return &sPayService{}
}

func init() {
	service.RegisterPayService(NewPayService())
}
