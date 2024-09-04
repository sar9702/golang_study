package main

import (
	"github.com/jung-kurt/gofpdf"
)

func main() {
	// gofpdf.New(orientationStr string, unitStr string, sizeStr string, fontDirStr string)
	// orientationStr : P(세로), L(가로)
	// unitStr : pt(point), mm, cm, in(inch)
	pdf := gofpdf.New("L", "mm", "A4", "")

	// 페이지 추가
	pdf.AddPage()

	// 페이지 크기를 가져온다.
	pageWidth, _ := pdf.GetPageSize()

	// 1. 제목
	fontSize := 20.0
	pdf.SetFont("Arial", "B", fontSize)

	text := "Hi, I'm Aerim"
	// width := pdf.GetStringWidth(text)

	// 텍스트를 페이지 정중앙에 배치하기 위해서 페이지의 정중앙을 계산한다.
	textWidth := pdf.GetStringWidth(text)
	x := (pageWidth - textWidth) / 2
	pdf.SetX(x)

	lineHeight := fontSize / 2.0 // 기본적으로 폰트 크기의 절반 정도를 줄 높이로 설정한다.
	pdf.Cell(0, lineHeight, text)

	// 줄 바꿈
	pdf.Ln(20)

	// 폰트 변경
	fontSize = 12.0
	pdf.SetFont("Arial", "", fontSize)

	// 두 번째 줄에 텍스트 출력
	boxWidth := (pageWidth - 40) / 2
	text = "Who am i\n" +
		"- Developer who works hard\n" +
		"- I worked as a 2D pipeline TD in the VFX industry for 4 years and 9 months.\n" +
		"- I like to heal while eating delicious food."
	pdf.MultiCell(boxWidth, fontSize/2.0, text, "1", "L", false)

	// PDF 파일을 저장하고 닫는다.
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		panic(err)
	}
}
