package imagetool

import (
	"bytes"
	"fmt"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

// ImageProcessor 提供图片处理功能
type ImageProcessor struct {
	backgroundPath string
	background     image.Image
	fontPath       string
}

type ImageQRCodeImageLocal struct {
	X             int
	Y             int
	X1            int
	Y1            int
	QrCodePx      int
	QrCodeContent string
}

// NewImageProcessor 初始化一个新的图片处理器
func NewImageProcessor(bgPath, fontPath string) (*ImageProcessor, error) {
	bg, err := gg.LoadImage(bgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load background image: %w", err)
	}
	return &ImageProcessor{
		backgroundPath: bgPath,
		background:     bg,
		fontPath:       fontPath,
	}, nil
}

// DrawText 在背景图片上写入文字
func (ip *ImageProcessor) DrawText(text string, x, y float64, color color.RGBA, fontSize float64) error {
	dc := gg.NewContextForImage(ip.background)
	if err := dc.LoadFontFace(ip.fontPath, fontSize); err != nil {
		return fmt.Errorf("failed to load font: %w", err)
	}
	dc.SetColor(color)
	dc.DrawStringAnchored(text, x, y, 0.5, 0.5)
	ip.background = dc.Image()
	return nil
}

// DrawUrlImage 读取外部图片绘制进背景图片中
func (ip *ImageProcessor) DrawUrlImage(imgURL string) (err error) {

	resp, err := http.Get(imgURL)
	if err != nil {
		fmt.Println("Error downloading image:", err)
		return
	}
	defer resp.Body.Close()

	imageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading image bytes:", err)
		return
	}

	tempImage, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	// 缩小图片，例如调整为原图的50%
	resizedImage := resize.Resize(0, 0, tempImage, resize.Lanczos3)
	dc := gg.NewContext(int(float64(resizedImage.Bounds().Dx())*0.73), int(float64(resizedImage.Bounds().Dy())*0.74))
	dc = gg.NewContextForImage(dc.Image())
	// 添加圆角效果
	radius := 8 // 圆角半径，可以根据需要调整
	dc.Push()   // 保存当前状态
	dc.SetRGBA(1, 1, 1, 0)
	dc.Clear()
	// 绘制四个圆角矩形覆盖在图片的四个角上
	dc.DrawRoundedRectangle(0, 0, float64(dc.Width()), float64(dc.Height()), float64(radius))
	dc.StrokePreserve()
	dc.Clip()
	dc.DrawImage(resizedImage, 0, 0) // 居中对齐
	// 弹出之前保存的状态，恢复到之前的画布状态
	dc.Pop()
	resizedImage = dc.Image()
	imgWidth := resizedImage.Bounds().Dx()
	imgHeight := resizedImage.Bounds().Dy()
	x := float64(ip.background.Bounds().Dx()-imgWidth) / 2
	y := float64(ip.background.Bounds().Dy()-imgHeight) / 3.2
	dc = gg.NewContextForImage(ip.background)
	// 重新绘制图片，此时会被圆角遮罩限制
	dc.DrawImage(resizedImage, int(x), int(y))
	ip.background = dc.Image()
	return
}

// PasteImage 在背景图片上粘贴另一张图片
func (ip *ImageProcessor) PasteImage(foregroundPath string, x, y int) error {
	foreground, err := gg.LoadImage(foregroundPath)
	if err != nil {
		return fmt.Errorf("failed to load foreground image: %w", err)
	}
	dc := gg.NewContextForImage(ip.background)
	dc.DrawImage(foreground, x, y)
	ip.background = dc.Image()
	return nil
}

func (ip *ImageProcessor) QRCodeImage(imageQRCodeImageLocal ImageQRCodeImageLocal) {
	// 设置二维码内容和大小
	qrCode, err := qrcode.New(imageQRCodeImageLocal.QrCodeContent, qrcode.Medium)
	if err != nil {
		panic(err)
	}
	qrCode.DisableBorder = true
	// 定义二维码在图片中的位置和大小
	qrCodeRect := image.Rect(imageQRCodeImageLocal.X, imageQRCodeImageLocal.Y, imageQRCodeImageLocal.X1, imageQRCodeImageLocal.Y1) // 例如：左上角坐标(100, 100)，宽高为200x200的矩形区域
	// 创建一个新的图像，用于绘制原始图片和二维码
	dc := gg.NewContext(ip.background.Bounds().Dx(), ip.background.Bounds().Dy())
	dc.DrawImage(ip.background, 0, 0)

	// 将二维码绘制到指定区域
	dc.SetColor(color.White)
	dc.DrawRectangle(float64(qrCodeRect.Min.X), float64(qrCodeRect.Min.Y), float64(qrCodeRect.Max.X-qrCodeRect.Min.X), float64(qrCodeRect.Max.Y-qrCodeRect.Min.Y))
	dc.Fill()
	dc.SetColor(color.Black)

	// 注意：此处需要将二维码转换为image.Image类型
	qrCodeImg := qrCode.Image(imageQRCodeImageLocal.QrCodePx) // 假设二维码大小为200x200像素
	dc.DrawImage(qrCodeImg, qrCodeRect.Min.X, qrCodeRect.Min.Y)
	ip.background = dc.Image()
}

// SaveToFile 保存处理后的图片到指定路径
func (ip *ImageProcessor) SaveToFile(outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()
	return jpeg.Encode(file, ip.background, nil)
}

func (ip *ImageProcessor) CosUploadImage() (tmpFile *os.File, err error) {
	// 创建一个Buffer来存储图片数据
	var (
		buf    bytes.Buffer
		imgBuf []byte
	)
	dc := gg.NewContextForImage(ip.background)
	// 将图片内容写入Buffer
	if err = dc.EncodePNG(&buf); err != nil {
		return
	}
	imgBuf = buf.Bytes()
	tmpFile, err = os.CreateTemp("", "temp_image_*.png")
	if err != nil {
		log.Fatalf("Error creating temp file: %v", err)
	}
	if _, err = tmpFile.Write(imgBuf); err != nil {
		return
	}
	tmpFile, err = os.Open(tmpFile.Name())
	return
}
