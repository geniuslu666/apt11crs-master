package admin

import (
	"APT/api/admin/pms"
	"APT/internal/library/translate"
	"context"
	"sync"
)

func (c *ControllerPms) TranslateText(ctx context.Context, req *pms.TranslateTextReq) (res *pms.TranslateTextRes, err error) {
	var (
		WaitGroup sync.WaitGroup
	)
	res = new(pms.TranslateTextRes)
	WaitGroup.Add(5)
	go func() {
		res.Zh, err = translate.TranslateText(ctx, req.Text, "zh")
		WaitGroup.Done()
	}()
	go func() {
		res.ZhCN, err = translate.TranslateText(ctx, req.Text, "zh_TW")
		WaitGroup.Done()
	}()
	go func() {
		res.Ko, err = translate.TranslateText(ctx, req.Text, "ko")
		WaitGroup.Done()
	}()
	go func() {
		res.Ja, err = translate.TranslateText(ctx, req.Text, "ja")
		WaitGroup.Done()
	}()
	go func() {
		res.En, err = translate.TranslateText(ctx, req.Text, "en")
		WaitGroup.Done()
	}()
	WaitGroup.Wait()
	return
}
