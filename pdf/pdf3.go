package pdf

import (
	"bytes"
	"fmt"

	"github.com/ledongthuc/pdf"
)

func ReadPdf7() {
	fmt.Println("ReadPdf7")
	content, err := readPdf7("example3.pdf")
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}

// 获取pdf文字内容
func readPdf7(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}

func ReadPdf8() {
	fmt.Println("ReadPdf8")
	content, err := readPdfGroup("example3.pdf")
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}

// 阅读按行分组的文本
func readPdfGroup(path string) (string, error) {
	f, r, err := pdf.Open(path)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, _ := p.GetTextByRow()
		for _, row := range rows {
			var content string
			for _, word := range row.Content {
				content += word.S

			}
			fmt.Println(content)
		}
	}
	return "", nil
}

func ReadPdf9() {
	fmt.Println("ReadPdf9")
	content, err := readPdfFormatAll("example-pdf.pdf")
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}

// PDF格式的所有文本
func readPdfFormatAll(path string) (string, error) {
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
