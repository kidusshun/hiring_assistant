package resumes

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/fumiama/go-docx"
	"github.com/ledongthuc/pdf"
)

func extractPDFText(fileBytes []byte) (string, error) {
	reader := bytes.NewReader(fileBytes)
	pdfReader, err := pdf.NewReader(reader, reader.Size())
	if err != nil {
		return "", fmt.Errorf("failed to create PDF reader: %v", err)
	}

	var text strings.Builder
	for i := 1; i <= pdfReader.NumPage(); i++ {
		page := pdfReader.Page(i)
		if page.V.IsNull() {
			continue
		}
		
		pageText, err := page.GetPlainText(nil)
		if err != nil {
			return "", fmt.Errorf("failed to extract text from page %d: %v", i, err)
		}
		text.WriteString(pageText)
	}

	return text.String(), nil
}

func extractDocxText(fileBytes []byte) (string, error) {
	reader := bytes.NewReader(fileBytes)
	size := int64(len(fileBytes))
	doc, err := docx.Parse(reader, size)
	
    if err != nil {
        return "", err
    }


	var text strings.Builder
	for _, item := range doc.Document.Body.Items {
		switch v := item.(type) {
		case *docx.Paragraph:
			text.WriteString(v.String())
			text.WriteString("\n")
		case *docx.Table:
			for _, row := range v.TableRows {
				for _, cell := range row.TableCells {
					for _, para := range cell.Paragraphs {
						text.WriteString(para.String())
						text.WriteString("\n")
					}
				}
			}
		}
	}
	result := text.String()
	
    
    // Return the extracted text
    return result, nil
}