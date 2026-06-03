package notify

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/model/entity"
	"context"
	"database/sql"
	"errors"

	"APT/api/notify/kefu"
)

func (c *ControllerKefu) Welcome(ctx context.Context, req *kefu.WelcomeReq) (res *kefu.WelcomeRes, err error) {
	var (
		CsFastContents []*entity.CsFastContent
		Language       = contexts.GetLanguage(ctx)
	)
	res = new(kefu.WelcomeRes)

	if err = dao.CsFastContent.Ctx(ctx).
		Where(&entity.CsFastContent{
			Type:   req.Type,
			Status: "Y",
		}).
		OrderDesc(dao.CsFastContent.Columns().Sort).
		OmitEmptyWhere().Scan(&CsFastContents); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}

	for _, v := range CsFastContents {
		Content := ""
		switch Language {
		case "zh":
			Content = v.ZhContent
			break
		case "zh_CN":
			Content = v.ZhCnContent
			break
		case "en":
			Content = v.EnContent
			break
		case "ja":
			Content = v.JaContent
			break
		case "ko":
			Content = v.KoContent
			break
		default:
			Content = v.EnContent
		}
		if v.Type == "WELCOME" {
			res.Welcome = Content
		} else if v.Type == "OPTION" {
			res.Options = append(res.Options, Content)
		}
	}
	return
}
