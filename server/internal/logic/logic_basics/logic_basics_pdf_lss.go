package logic_basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/signintech/gopdf"
)

type sBasicsPDF struct{}

func NewBasicsPDF() *sBasicsPDF {
	return &sBasicsPDF{}
}

func init() {
	service.RegisterBasicsPDF(NewBasicsPDF())
}

func (s *sBasicsPDF) OrderReceipt(ctx context.Context, ipt *input_basics.PdfLssIpt) (pdfFailPath string, err error) {
	var (
		pdf        = gopdf.GoPdf{}
		family     = "NotoSansSC"
		familyPath = "ttf/NotoSansSC-6.ttf"
		texts      = []string{
			"インボイス番号：T5120101066551",
			"株式会社HIWIN",
			"〒598-0091",
			"大阪府泉南郡田尻町嘉祥寺５８８",
		}
	)
	pdfFailPath = "resource/public/pdf/" + guid.S() + ".pdf"
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 600, H: 400}})
	pdf.AddPage()
	if err = pdf.AddTTFFont(family, familyPath); err != nil {
		return
	}
	if err = pdf.SetFont(family, "", 14); err != nil {
		return
	}
	
	// 设置标题
	pdf.SetX(10)
	pdf.SetY(30)
	if err = pdf.Cell(nil, "領収書"); err != nil {
		return
	}
	
	// 日期和客户
	pdf.SetX(400)
	pdf.SetY(30)
	if err = pdf.Cell(nil, gtime.Now().Format("Y年m月d日")); err != nil {
		return
	}
	
	pdf.SetX(10)
	pdf.SetY(70)
	
	if err = pdf.Cell(nil, ipt.UserName); err != nil {
		return
	}
	pdf.SetX(500)
	pdf.SetY(70)
	if err = pdf.Cell(nil, "様"); err != nil {return}
	pdf.Line(10, 90, 510, 90)
	// 金额框
	pdf.RectFromUpperLeftWithStyle(10, 110, 500, 30, "D")
	pdf.SetX(30)
	pdf.SetY(115)
	if err = pdf.Cell(nil, fmt.Sprintf("¥%dー", ipt.OrderAmount)); err != nil {
		return
	}
	pdf.SetX(480)
	if err = pdf.Cell(nil, "JPY"); err != nil {
		return
	}
	
	// 明细说明
	pdf.SetY(150)
	pdf.SetX(70)
	
	if err = pdf.MultiCell(&gopdf.Rect{
		W: 420,
		H: 30,
	}, fmt.Sprintf("但し、住— %s %s %s〜%s宿泊代として", ipt.PropertyName, ipt.OrderSn, ipt.StartTime, ipt.EndTime)); err != nil {
		return
	}
	if err = pdf.SetFont(family, "", 10); err != nil {
		return
	}
	
	// 明细说明
	pdf.SetY(190)
	pdf.SetX(70)
	if err = pdf.Cell(nil, "上記正に領収致しました。"); err != nil {
		return
	}
	if err = pdf.SetFont(family, "", 14); err != nil {
		return
	}
	
	for i := 0; i < len(texts); i++ {
		pdf.SetY(gconv.Float64(240) + gconv.Float64(i*20))
		pdf.SetX(70)
		if err = pdf.Cell(nil, texts[i]); err != nil {
			return
		}
	}
	
	// 印章图
	if err = pdf.Image("ttf/img.png", 400, 240, &gopdf.Rect{
		W: 100,
		H: 100,
	}); err != nil {
		return
	}
	
	// 输出 PDF
	if err = pdf.WritePdf(pdfFailPath); err != nil {
		return
	}
	return
}
