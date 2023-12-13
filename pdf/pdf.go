package pdf

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/signintech/gopdf"
)

var (
	resourcesPath string
	pdf           *gopdf.GoPdf
)

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	resourcesPath = filepath.Join(cwd, "pdf")

}

// 写入文字 生成pdf
func BaseText() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	if err := pdf.AddTTFFont("simkai", resourcesPath+"/font/simkai.ttf"); err != nil {
		log.Panic(err.Error())
		return
	}

	if err := pdf.AddTTFFont("regular", resourcesPath+"/font/LiberationSerif-Regular.ttf"); err != nil {
		log.Panic(err.Error())
		return
	}

	if err := pdf.AddTTFFont("japanese", resourcesPath+"/font/Natsuzemi Maru Gothic Black.ttf"); err != nil {
		log.Panic(err.Error())
		return
	}

	if err := pdf.SetFont("simkai", "", 14); err != nil {
		log.Panic(err.Error())
		return
	}
	pdf.SetXY(20, 50)               //设置坐标点
	pdf.SetTextColor(156, 197, 140) //设置文字颜色
	pdf.Cell(nil, "欢迎使用goPDF ！！！")  //设置写入内容

	pdf.SetFont("regular", "", 32)
	pdf.SetXY(250, 200)                //设置坐标点
	pdf.SetTextColor(0, 0, 0)          //设置文字颜色
	pdf.Cell(nil, "gopher and gopher") //设置写入内容

	pdf.SetFont("japanese", "", 20)
	pdf.SetXY(150, 400)     //设置坐标点
	pdf.Cell(nil, "株式会社です") //设置写入内容

	path := fmt.Sprintf("./pdf/create/%d.pdf", time.Now().Unix())
	pdf.WritePdf(path)
}

// 写入图片生成pdf
func WriteImage() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	pdf.Image(resourcesPath+"/pictures/gopher01.jpg", 200, 50, nil)

	//image bytes
	b, err := os.ReadFile(resourcesPath + "/pictures/gopher02_color.png")
	if err != nil {
		log.Panic(err.Error())
	}

	imgH1, err := gopdf.ImageHolderByBytes(b)
	if err != nil {
		log.Panic(err.Error())
	}
	if err := pdf.ImageByHolder(imgH1, 200, 250, &gopdf.Rect{
		W: 120,
		H: 120,
	}); err != nil {
		log.Panic(err.Error())
	}

	//image io.Reader
	file, err := os.Open(resourcesPath + "/pictures/gopher02_color.png")
	if err != nil {
		log.Panic(err.Error())
	}

	imgH2, err := gopdf.ImageHolderByReader(file)
	if err != nil {
		log.Panic(err.Error())
	}

	maskHolder, err := gopdf.ImageHolderByPath(resourcesPath + "/pictures/mask.png")
	if err != nil {
		log.Panic(err.Error())
	}

	maskOpts := gopdf.MaskOptions{
		Holder: maskHolder,
		ImageOptions: gopdf.ImageOptions{
			X: 120,
			Y: 320,
			Rect: &gopdf.Rect{
				W: 300,
				H: 300,
			},
		},
	}

	transparency, err := gopdf.NewTransparency(0.5, "")
	if err != nil {
		log.Panic(err.Error())
	}

	imOpts := gopdf.ImageOptions{
		X:            120,
		Y:            400,
		Mask:         &maskOpts,
		Transparency: &transparency,
		Rect:         &gopdf.Rect{W: 300, H: 300},
	}
	pdf.Rotate(270.0, 270, 550)
	if err := pdf.ImageByHolderWithOptions(imgH2, imOpts); err != nil {
		log.Panic(err.Error())
	}

	path := fmt.Sprintf("./pdf/create/%d.pdf", time.Now().Unix())
	pdf.WritePdf(path)
}
