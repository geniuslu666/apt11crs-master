package app

import (
	"APT/internal/library/cabinetApi"
	"APT/internal/library/contexts"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/app/cabinet"
	"APT/internal/model"
	"APT/internal/service"
)

func (c *ControllerCabinet) CabinetList(ctx context.Context, req *cabinet.CabinetListReq) (res *cabinet.CabinetListRes, err error) {
	var (
		cabinetResponse  *cabinetApi.CabinetListResponse
		CabinetApiConfig *model.CabinetApiConfig
	)
	if CabinetApiConfig, err = service.BasicsConfig().GetCabinetApi(ctx); err != nil {
		return
	}
	if cabinetResponse, err = cabinetApi.NewClient(ctx, CabinetApiConfig).CabinetList(ctx); err != nil {
		return
	}
	if cabinetResponse.Code != 0 {
		err = gerror.New(cabinetResponse.Msg)
		return
	}
	res = new(cabinet.CabinetListRes)
	res.List = cabinetResponse.Data
	var Language = contexts.GetLanguage(ctx)
	for _, v := range res.List {
		switch Language {
		case "zh":
			v.Name = v.NameZh
			v.MchBranchName = v.MchBranchNameZh
			v.MchBranchAddress = v.MchBranchAddressZh
		case "en":
			v.Name = v.NameEn
			v.MchBranchName = v.MchBranchNameEn
			v.MchBranchAddress = v.MchBranchAddressEn
		case "ja":
			v.Name = v.NameJa
			v.MchBranchName = v.MchBranchNameJa
			v.MchBranchAddress = v.MchBranchAddressJa
		case "ko":
			v.Name = v.NameKo
			v.MchBranchName = v.MchBranchNameKo
			v.MchBranchAddress = v.MchBranchAddressKo
		case "zh_CN":
			v.Name = v.NameTw
			v.MchBranchName = v.MchBranchNameTw
			v.MchBranchAddress = v.MchBranchAddressTw
		}
	}
	return
}
