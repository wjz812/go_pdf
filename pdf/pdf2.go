package pdf

import (
	"fmt"

	"rsc.io/pdf"
)

func ReadPdf5() {
	file, err := pdf.Open("111.pdf")
	if err != nil {
		panic(err)
	}
	fmt.Println(file.NumPage())
	fmt.Println(file.Page(2).Content())
}
