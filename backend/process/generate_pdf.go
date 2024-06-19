package process

import (
	"bytes"

	"github.com/carsonkiibi/pdfapp/backend/process/commands"
	"github.com/jung-kurt/gofpdf"
)

// actual spaghetti rn for demo
func GeneratePDF(tokens []commands.Token) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	defaultSize := 11.0
	pdf.SetFont("Arial", "", defaultSize)

	var currentWidth float64
	var currentHeight float64 = 10.0
	var setStyle string = ""

	for _, token := range tokens {
		if token.Type == "TEXT" {
			currentWidth = pdf.GetStringWidth(token.Literal)
			pdf.Cell(currentWidth, currentHeight, token.Literal)

		} else if token.Type == "TEXT_MOD" {
			if token.Attributes.Bold {
				setStyle += "B"
			}
			if token.Attributes.Underline {
				setStyle += "U"
			}
			if token.Attributes.Italic {
				setStyle += "I"
			}
			if token.Attributes.Size > 0 {
				pdf.SetFontSize(float64(token.Attributes.Size))
			}
			pdf.SetFontStyle(setStyle)
			currentWidth = pdf.GetStringWidth(token.Literal)
			pdf.Cell(currentWidth, currentHeight, token.Literal)
			setStyle = ""
		}
		pdf.SetFontSize(defaultSize)
		pdf.SetFontStyle("")
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
