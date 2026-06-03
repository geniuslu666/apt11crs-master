package main

import (
	_ "APT/internal/logic"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		EmsParams *input_basics.SendEmsInp
		ctx       = gctx.GetInitCtx()
		err       error
	)
	FilePath := "https://apt11-dev-storage.oss-accelerate.aliyuncs.com/hotgo/attachment/2025-06-18/dapihttlsg9kacxqm1.pdf"
	EmsParams = new(input_basics.SendEmsInp)
	EmsParams.Content = `
<p>お客様へ</p>
<p style="text-indent:2em">ご注文の領収書を添付いたしましたので、ご確認ください。</p>
<p style="text-indent:2em">Dear Customer, please find attached the receipt for your order. Kindly check it.</p>
<p style="text-indent:2em">親愛的顧客您好，已隨信附上您的訂單收據，敬請查收。</p>
<a href="` + FilePath + `">領収書</a>`
	EmsParams.Event = "text"
	EmsParams.Email = "1256005331@qq.com"
	if err = service.BasicsEmsLog().Send(ctx, EmsParams); err != nil {
		panic(err)
	}
}
