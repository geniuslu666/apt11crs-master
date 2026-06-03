package main

import (
	"APT/internal/global"
	"APT/internal/library/storager"
	_ "APT/internal/logic"
	"APT/internal/model"
	"APT/internal/service"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/signintech/gopdf"
)

func main() {
	var (
		pdf        = gopdf.GoPdf{}
		err        error
		name       = "インボイス番号高杨안녕하세요"
		family     = "NotoSansCJK"
		familyPath = "/Users/gaoyang/go/src/APT11_CRS/server/ttf/NotoSansSC-6.ttf"
		texts      = []string{
			"インボイス番号：T5120101066551",
			"株式会社HIWIN",
			"〒598-0091",
			"大阪府泉南郡田尻町嘉祥寺５８８",
		}
		filePath     = guid.S() + ".pdf"
		UploadConfig *model.UploadConfig
		ctx          = gctx.GetInitCtx()
	)
	global.Init(ctx)
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 600, H: 400}})
	pdf.AddPage()
	if err = pdf.AddTTFFont(family, familyPath); err != nil {
		panic(err)
	}
	if err = pdf.SetFont(family, "", 14); err != nil {
		panic(err)
	}

	// 设置标题
	pdf.SetX(10)
	pdf.SetY(30)
	if err = pdf.Cell(nil, "領収書"); err != nil {
		panic(err)
	}

	// 日期和客户
	pdf.SetX(400)
	pdf.SetY(50)
	if err = pdf.Cell(nil, "2025年06月○○日"); err != nil {
		panic(err)
	}

	pdf.SetX(300)
	pdf.SetY(70)

	if err = pdf.Cell(nil, "｜"); err != nil {
		panic(err)
	}

	pdf.SetX(400)
	pdf.SetY(90)

	if err = pdf.Cell(nil, name+"  様"); err != nil {
		panic(err)
	}

	// 金额框
	pdf.RectFromUpperLeftWithStyle(10, 110, 500, 30, "D")
	pdf.SetX(30)
	pdf.SetY(115)
	if err = pdf.Cell(nil, "¥〇〇〇,〇〇〇ー"); err != nil {
		panic(err)
	}
	pdf.SetX(480)
	if err = pdf.Cell(nil, "JPY"); err != nil {
		panic(err)
	}

	// 明细说明
	pdf.SetY(150)
	pdf.SetX(70)
	if err = pdf.MultiCell(&gopdf.Rect{
		W: 420,
		H: 30,
	}, "但し、住— ○○○○号店 R374208502 2025.5.25〜2025.5.27宿泊代として"); err != nil {
		panic(err)
	}
	if err = pdf.SetFont(family, "", 10); err != nil {
		panic(err)
	}

	// 明细说明
	pdf.SetY(190)
	pdf.SetX(70)
	if err = pdf.Cell(nil, "上記正に領収致しました。"); err != nil {
		panic(err)
	}
	if err = pdf.SetFont(family, "", 14); err != nil {
		panic(err)
	}

	for i := 0; i < len(texts); i++ {
		pdf.SetY(gconv.Float64(240) + gconv.Float64(i*20))
		pdf.SetX(70)
		if err = pdf.Cell(nil, texts[i]); err != nil {
			panic(err)
		}
	}

	// 印章图
	if err = pdf.Image("/Users/gaoyang/go/src/APT11_CRS/server/ttf/img.png", 400, 240, &gopdf.Rect{
		W: 100,
		H: 100,
	}); err != nil {
		panic(err)
	}

	// 电子发票标识图
	//if err = pdf.Image("./electronic_invoice.png", 50, 250, nil);err != nil{
	//	panic(err)
	//}

	// 输出 PDF
	if err = pdf.WritePdf(filePath); err != nil {
		panic(err)
	}

	// 上传文件
	if UploadConfig, err = service.BasicsConfig().GetUpload(ctx); err != nil {
		return
	}
	g.Log().Debug(ctx, "上传文件：", filePath)
	// 上传到驱动
	if filePath, err = storager.New(UploadConfig.Drive).UploadFile(ctx, filePath); err != nil {
		panic(err)
		return
	}
	g.Log().Debug(ctx, "云存储：", filePath)
	filePath = storager.LastUrl(ctx, filePath, UploadConfig.Drive)
	g.Log().Debug(ctx, "云存储：", filePath)

}
