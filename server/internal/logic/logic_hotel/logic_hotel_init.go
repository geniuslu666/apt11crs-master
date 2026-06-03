package logic_hotel

import "APT/internal/service"

type sHotelService struct{}

func NewHotelService() *sHotelService {
	return &sHotelService{}
}

func init() {
	service.RegisterHotelService(NewHotelService())
}
