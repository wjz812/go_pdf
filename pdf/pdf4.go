package pdf

import (
	"fmt"
	"os"

	"github.com/unidoc/unipdf/v3/extractor"
	pdf "github.com/unidoc/unipdf/v3/model"
)

func ReadPdf10() {

	// f, _, err := model.NewPdfReaderFromFile("example.pdf", nil)
	// if err != nil {
	// 	panic(err)
	// }
	// // Get the number of pages in the PDF file.
	// numPages, err := f.GetNumPages()
	// fmt.Println("pdf 总页数", numPages)
	// if err != nil {
	// 	panic(err)
	// }
	// page, err := f.GetPage(1)
	// if err != nil {
	// 	panic(err)
	// }
	// content, err := page.GetAllContentStreams()
	// fmt.Printf("Page %d content:\n%s\n", 1, content)
	// Loop through each page and extract text content.
	// for i := 1; i <= numPages; i++ {
	// 	page, err := f.GetPage(i)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	content, err := page.GetAllContentStreams()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	// fmt.Println(content)
	// 	fmt.Printf("Page %d content:\n%s\n", i, content)
	// }

	inputPath := "example.pdf"

	text, err := getPdfTextContent(inputPath, 1)
	if err != nil {
		return
	}
	fmt.Println(text)
}

func getPdfTextContent(inputPath string, pageNum int) (string, error) {
	pdfFile, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer pdfFile.Close()

	//获取某一页的信息
	pdfReader, err := pdf.NewPdfReader(pdfFile)
	if err != nil {
		return "", err
	}
	page, err := pdfReader.GetPage(pageNum)
	//导出文本
	extract, err := extractor.New(page)
	if err != nil {
		return "", err
	}
	text, err := extract.ExtractText()
	if err != nil {
		return "", err
	}

	return text, nil
}
