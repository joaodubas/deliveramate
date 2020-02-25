package storage

import "testing"

func TestDocumentFormatter(t *testing.T) {
	testCases := []struct {
		name     string
		document string
		expected string
		err      error
	}{
		{
			"cnpj: do nothing if properly formatted",
			"26.810.612/0001-33",
			"26.810.612/0001-33",
			nil,
		},
		{
			"cnpj: add proper point separator",
			"24302190/0001-60",
			"24.302.190/0001-60",
			nil,
		},
		{
			"cnpj: add proper format",
			"23840372000121",
			"23.840.372/0001-21",
			nil,
		},
		{
			"cpf: add proper format",
			"36211693850",
			"362.116.938-50",
			nil,
		},
		{
			"cpf: add missing point sep",
			"960361.506-44",
			"960.361.506-44",
			nil,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			d, err := DocumentFormatter(c.document)
			if err != nil && c.err == nil {
				t.Errorf("Format document %s should not fail (%v)", c.document, err)
			} else if err == nil && c.err != nil {
				t.Errorf("Format document %s should fail (%v)", c.document, c.err)
			} else if c.err == nil && c.expected != d {
				t.Errorf("Format document %s | expected %s | got %s", c.document, c.expected, d)
			}
		})
	}
}
