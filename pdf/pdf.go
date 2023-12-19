package pdf

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/google/go-tika/tika"
	"github.com/ledongthuc/pdf"
	"github.com/signintech/gopdf"
)

var (
	resourcesPath string
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

// 页眉 页脚写入
func WriteHeaderAndFooter() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	if err := pdf.AddTTFFont("simkai", resourcesPath+"/font/simkai.ttf"); err != nil {
		log.Panic(err.Error())
		return
	}

	if err := pdf.SetFont("simkai", "", 8); err != nil {
		log.Panic(err.Error())
		return
	}

	pdf.AddHeader(func() {
		pdf.SetY(5)
		pdf.Cell(nil, "cad工程图")
	})

	pdf.AddFooter(func() {
		pdf.SetXY(570, 825)
		pdf.Cell(nil, "第1页")
	})

	pdf.AddPage()
	pdf.SetFontSize(20)
	pdf.SetXY(200, 400)
	pdf.Text("组装图")

	pdf.SetLineWidth(0.5)
	pdf.SetLineType("dashed")
	pdf.Line(0, 20, 585, 20)

	pdf.SetLineWidth(0.5)
	pdf.SetLineType("dotted")
	pdf.Line(0, 815, 585, 815)

	path := fmt.Sprintf("./pdf/create/%d.pdf", time.Now().Unix())
	pdf.WritePdf(path)
}

// 写入基础图形
func WriteBaseGraph() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	pdf.SetLineWidth(0.5)
	pdf.SetLineType("dashed")
	pdf.Line(0, 20, 585, 20)

	pdf.SetLineWidth(0.5)
	pdf.SetLineType("dotted")
	pdf.Line(0, 50, 585, 50)

	pdf.SetLineWidth(1) //设置线宽
	pdf.SetLineType("dashed")
	pdf.Oval(75, 75, 50, 50) //画圆  椭圆or圆

	pdf.SetStrokeColor(255, 0, 0) //设置边线颜色
	pdf.SetLineWidth(2)
	pdf.SetFillColor(0, 255, 0) //设置填充颜色  默认为黑色
	pdf.Polygon([]gopdf.Point{{X: 100, Y: 50}, {X: 100, Y: 100}, {X: 150, Y: 150}, {X: 200, Y: 150}, {X: 200, Y: 100}}, "DF")

	pdf.SetStrokeColor(255, 0, 0)
	pdf.SetLineWidth(2)
	pdf.SetFillColor(0, 255, 0)
	// style : "D" 只画外边框  "F":只填充内部 "DF" or "FD":内外都显示
	err := pdf.Rectangle(196.6, 336.8, 398.3, 379.3, "FD", 10, 20)

	if err != nil {
		return
	}

	pdf.SetStrokeColorCMYK(88, 49, 0, 0)
	pdf.SetLineWidth(2)
	pdf.SetFillColorCMYK(0, 5, 89, 0)
	err = pdf.Rectangle(196.6, 436.8, 398.3, 479.3, "DF", 10, 20)
	if err != nil {
		return
	}

	path := fmt.Sprintf("./pdf/create/%d.pdf", time.Now().Unix())
	pdf.WritePdf(path)
}

// 生成PDF 并且加密
func WritePasswordProtection() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4, //595.28, 841.89 = A4
		Protection: gopdf.PDFProtectionConfig{
			UseProtection: true,                                                                     //指定是否启用PDF文档的保护功能
			Permissions:   gopdf.PermissionsPrint | gopdf.PermissionsCopy | gopdf.PermissionsModify, //用于设置PDF文档的权限 打印、复制和修改
			OwnerPass:     []byte("123456"),                                                         //设置PDF文档的所有者密码为"123456"。
			UserPass:      []byte("123456789")},                                                     //设置PDF文档的用户密码为"123456789"
	})

	pdf.AddPage()
	pdf.AddTTFFont("simkai", resourcesPath+"/font/simkai.ttf")
	pdf.SetFont("simkai", "", 14) // 这里你可能还需要设置一下字体和字体大小
	pdf.Cell(nil, "goPDF")
	path := fmt.Sprintf("./pdf/create/%d.pdf", time.Now().Unix())
	pdf.WritePdf(path)
}

func DownloadPDF() {
	fileUrl := "https://tcpdf.org/files/examples/example_012.pdf"
	if err := downloadFile("example-pdf.pdf", fileUrl); err != nil {
		panic(err)
	}
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	pdf.AddPage()

	tpl1 := pdf.ImportPage("example-pdf.pdf", 1, "/MediaBox")
	//
	pdf.UseImportedTemplate(tpl1, 50, 100, 400, 0)
	//

	pdf.AddPage()
	tpl2 := pdf.ImportPage("example-pdf.pdf", 2, "/MediaBox")
	pdf.UseImportedTemplate(tpl2, 50, 100, 400, 0)

	pageNum := pdf.GetNumberOfPages()
	fmt.Println("已有pdf 总页数：", pageNum)

	path := fmt.Sprintf("./pdf/create/%d.pdf", time.Now().Unix())
	pdf.WritePdf(path)
}

// 下载文件
func downloadFile(filePath string, url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(filePath)

	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}

func WriteTrimBox() {
	pdf := gopdf.GoPdf{}
	mm6ToPx := 50.00

	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4, //595.28, 841.89 = A4
		TrimBox:  gopdf.Box{Left: mm6ToPx, Top: mm6ToPx, Right: 595 - mm6ToPx, Bottom: 842 - mm6ToPx},
	})

	opt := gopdf.PageOption{
		PageSize: gopdf.PageSizeA4, //595.28, 841.89 = A4
		TrimBox:  &gopdf.Box{Left: mm6ToPx, Top: mm6ToPx, Right: 595 - mm6ToPx, Bottom: 842 - mm6ToPx},
	}
	pdf.AddPageWithOption(opt)

	pdf.AddTTFFont("simkai", resourcesPath+"/font/simkai.ttf")
	pdf.SetFont("simkai", "", 14) // 这里你可能还需要设置一下字体和字体大小
	pdf.Cell(nil, "goPDF")

	path := fmt.Sprintf("./pdf/create/%d.pdf", time.Now().Unix())
	pdf.WritePdf(path)
}

func WriteTable() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	//插入标题
	pdf.AddTTFFont("simkai", resourcesPath+"/font/simkai.ttf")
	pdf.SetFont("simkai", "", 20) // 这里你可能还需要设置一下字体和字体大小
	pdf.SetTextColor(0x00, 0x00, 0xff)
	pdf.SetXY(288, 20)
	pdf.Cell(nil, "你好")

	//插入图片
	pdf.Image(resourcesPath+"/pictures/background.jpg", 47, 50, &gopdf.Rect{W: 500, H: 300})

	//添加表格的标题
	pdf.SetXY(47, 370)
	pdf.SetFontSize(12)
	alignOption := gopdf.CellOption{Align: gopdf.Center | gopdf.Middle,
		Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top}

	pdf.CellWithOption(&gopdf.Rect{
		W: 500,
		H: 50,
	}, "旗舰店12月入库商品列表", alignOption)
	pdf.SetFontSize(10)
	pdf.SetTextColor(0xa8, 0xa8, 0xa8)

	//生成表格的表头
	pdf.SetX(47)
	pdf.SetY(420)
	pdf.CellWithOption(&gopdf.Rect{
		W: 50,
		H: 40,
	}, "商品ID", alignOption)

	pdf.SetX(97)
	pdf.SetY(420)

	pdf.CellWithOption(&gopdf.Rect{
		W: 400,
		H: 40,
	}, "商品名称", alignOption)

	pdf.SetX(497)
	pdf.SetY(420)

	pdf.CellWithOption(&gopdf.Rect{
		W: 50,
		H: 40,
	}, "商品库存", alignOption)

	//生成表格内容
	goods := getGoodList()

	pdf.SetTextColor(0x00, 0x00, 0x00)
	alignLeft := gopdf.CellOption{Align: gopdf.Left | gopdf.Middle,
		Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top}

	for i, v := range goods {
		curY := 460 + i*40
		pdf.SetXY(47, float64(curY))

		pdf.CellWithOption(&gopdf.Rect{
			W: 50,
			H: 40,
		}, " "+strconv.FormatInt(v.GoodsId, 10), alignLeft)

		pdf.SetX(97)
		pdf.SetY(float64(curY))

		pdf.CellWithOption(&gopdf.Rect{
			W: 400,
			H: 40,
		}, " "+v.GoodsName, alignLeft)

		pdf.SetX(497)
		pdf.SetY(float64(curY))

		pdf.CellWithOption(&gopdf.Rect{
			W: 50,
			H: 40,
		}, " "+strconv.FormatInt(v.Stock, 10), alignLeft)

	}

	path := fmt.Sprintf("./pdf/create/%d.pdf", time.Now().Unix())
	pdf.WritePdf(path)
}

// 获取商品列表
func getGoodList() (list []Good) {
	listGood := []Good{}
	listGood = append(listGood, Good{GoodsId: 1, GoodsName: "蜂蜜牛奶手工皂", Stock: 100})
	listGood = append(listGood, Good{GoodsId: 2, GoodsName: "乐传茶具", Stock: 35})
	listGood = append(listGood, Good{GoodsId: 3, GoodsName: "蓝牙音箱", Stock: 72})

	return listGood
}

type Good struct {
	GoodsId   int64  `json:"goodsId"`
	GoodsName string `json:"goodsName"`
	Stock     int64  `json:"stock"`
}

// 获取pdf文件总页数
func GetPageCount() {
	//fileUrl := "https://tcpdf.org/files/examples/example_012.pdf"

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	if err := pdf.AddTTFFont("simkai", resourcesPath+"/font/simkai.ttf"); err != nil {
		log.Panic(err.Error())
		return
	}
	pdf.SetFont("simkai", "", 14)
	pdf.Cell(nil, "欢迎使用goPDF ！！！") //设置写入内容
	pdf.Br(20)
	pageNum := pdf.GetNumberOfPages()
	fmt.Println("已有pdf 总页数：", pageNum)

	path := fmt.Sprintf("./pdf/create/%d.pdf", time.Now().Unix())
	pdf.WritePdf(path)
}

func WriteLinks() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4}) //595.28, 841.89 = A4
	pdf.AddPage()

	if err := pdf.AddTTFFont("simkai", resourcesPath+"/font/simkai.ttf"); err != nil {
		log.Panic(err.Error())
		return
	}
	pdf.SetFont("simkai", "", 14)

	pdf.SetXY(30, 40)
	pdf.SetTextColor(156, 197, 140) //设置文字颜色
	pdf.Text("Link to example.com")
	pdf.AddExternalLink("http://example.com/", 27.5, 28, 125, 15)

	pdf.SetXY(30, 70)
	pdf.Text("Link to second page")
	pdf.AddInternalLink("anchor", 27.5, 58, 120, 15)

	pdf.AddPage()
	pdf.SetXY(30, 100)
	pdf.SetAnchor("anchor")
	pdf.Text("Anchor position")

	path := fmt.Sprintf("./pdf/create/%d.pdf", time.Now().Unix())
	pdf.WritePdf(path)
}

// 读取pdf文件
func ReadPDF() {
	// pdf := gopdf.GoPdf{}
	fileUrl := "https://tcpdf.org/files/examples/example_012.pdf"
	if err := downloadFile("example-pdf.pdf", fileUrl); err != nil {
		panic(err)
	}

	//pdfReader, err := pdf.Open(fileUrl)

	pdf.DebugOn = true

	content, err := readPdf("111.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)

}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	totalPage := r.NumPage()
	fmt.Println("已有pdf 总页数：", totalPage)
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	fmt.Println("====", b, err)
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}

func ReadPdf2() {

	content, err := readPdf2("111.pdf")
	if err != nil {
		panic(err)
	}

	fmt.Println(content)
}

func readPdf2(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		// var lastTextStyle pdf.Text
		// texts := p.Content().Text
		// for _, text := range texts {
		// 	fmt.Printf("Font: %s, Font-size: %f, x: %f, y: %f, content: %s \n", lastTextStyle.Font, lastTextStyle.FontSize, lastTextStyle.X, lastTextStyle.Y, lastTextStyle.S)
		// 	lastTextStyle = text
		// }

		// columns, err := p.GetTextByColumn()
		// for _, row := range columns {
		// 	println(">>>> column: ", row.Position, row.Content)
		// 	for _, word := range row.Content {
		// 		fmt.Println(word.S)
		// 	}
		// }
		// b, err := r.GetPlainText()
		// fmt.Println("====", b, err)
		// if err != nil {
		// 	return "", err
		// }
		rows, _ := p.GetTextByRow()
		for _, row := range rows {
			// println(">>>> row: ", row.Position)
			var content string
			for _, word := range row.Content {
				content += word.S

			}
			fmt.Println(content)
		}
	}
	return "", nil
}

func readPdf3(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		var lastTextStyle pdf.Text
		texts := p.Content().Text
		for _, text := range texts {
			fmt.Printf("Font: %s, Font-size: %f, x: %f, y: %f, content: %s \n", lastTextStyle.Font, lastTextStyle.FontSize, lastTextStyle.X, lastTextStyle.Y, lastTextStyle.S)
			lastTextStyle = text
		}
	}
	return "", nil
}

func ReadPdf4() {
	// var filePath string
	// flag.StringVar(&filePath, "fp", "", "111.pdf")
	// flag.Parse()
	// if filePath == "" {
	// 	panic("file path must be provided")
	// }
	content, err := readPdf4("example3.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}

func readPdf4(path string) (string, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}

	client := tika.NewClient(nil, "http://127.0.0.1:9998")
	return client.Parse(context.TODO(), f)
}

func ReadPdf6() {
	// file, err := os.Open("111.pdf")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// reader := tika.Open(file)

	// result, err := reader.ReadAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 打印PDF文件的内容
	// fmt.Println(string(result))
}

func readPdf6(path string) string {
	file, reader, err := pdf.Open(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	content := ""
	for i := 1; i <= reader.NumPage(); i++ {
		text := reader.Page(i).Content().Text
		for _, item := range text {
			content += item.S
		}
	}
	return content
}
