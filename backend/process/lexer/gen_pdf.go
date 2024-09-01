package main

// func GeneratePDF(tokens []Token) ([]byte, error) {
// 	pdf := gofpdf.New("P", "mm", "A4", "")
// 	pdf.AddPage()

// 	defaultTextSize := 11.0
// 	pdf.SetFont("Arial", "", defaultTextSize)

// 	var currentWidth float64
// 	var currentHeight float64 = 10.0
// 	var setStyle string = ""

// 	for _, token := range tokens {
// 		switch token.Type {
// 		case TEXT:
// 			pdf.Cell(currentWidth, currentHeight, token.Literal)
// 		case TEXTMOD:
// 			if token.Attributes.Bold {
// 				setStyle += "B"
// 			}
// 			if token.Attributes.Italic {
// 				setStyle += "I"
// 			}
// 			if token.Attributes.Underline {
// 				setStyle += "U"
// 			}
// 			if token.Attributes.Size > 0 {
// 				pdf.SetFontSize(float64(token.Attributes.Size))
// 			}
// 		default:
// 			fmt.Println("token not recognized")
// 			pdf.SetFontStyle(setStyle)
// 			currentWidth = pdf.GetStringWidth(token.Literal)
// 			pdf.Cell(currentWidth, currentHeight, token.Literal)
// 			setStyle = ""
// 		}
// 		pdf.SetFontSize(defaultTextSize)
// 		pdf.SetFontStyle("")

// 	}

// 	var buf bytes.Buffer
// 	err := pdf.Output(&buf)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return buf.Bytes(), nil
// }
