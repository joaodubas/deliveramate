package storage

import (
	"fmt"
	"regexp"
)

func DocumentFormatter(doc string) (string, error) {
	rg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return doc, fmt.Errorf("DocumentFormatter: regex compilation error (%w)", ErrorDocumentRegex)
	}
	docFormatted := rg.ReplaceAllString(doc, "")
	if len(docFormatted) == 0 {
		return doc, fmt.Errorf("DocumentFormatter: got document %s | formatted document %s (%w)", doc, docFormatted, ErrorDocumentMalformed)
	} else if len(docFormatted) <= 11 {
		docFormatted = cpfFormatter(docFormatted)
	} else {
		docFormatted = cnpjFormatter(docFormatted)
	}
	return docFormatted, nil
}

func cpfFormatter(doc string) string {
	rg, _ := regexp.Compile(`(\d{3})(\d{3})(\d{3})(\d{2})`)
	return rg.ReplaceAllString(fmt.Sprintf("%011s", doc), `$1.$2.$3-$4`)
}

func cnpjFormatter(doc string) string {
	rg, _ := regexp.Compile(`(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})`)
	return rg.ReplaceAllString(fmt.Sprintf("%14s", doc), `$1.$2.$3/$4-$5`)
}
