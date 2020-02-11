package storage

import (
	"errors"
	"testing"
)

func TestNewPartnerDocumentFormat(t *testing.T) {
	testCases := []struct {
		name        string
		doc         string
		expectedDoc string
		err         error
	}{
		{
			"success cnpj",
			"00.000.000/0000-00",
			"00.000.000/0000-00",
			nil,
		},
		{
			"success cpf",
			"1111111111",
			"011.111.111-11",
			nil,
		},
		{
			"failure",
			"./-",
			"",
			ErrorDocumentMalformed,
		},
	}
	for _, d := range testCases {
		t.Run(d.name, func(t *testing.T) {
			p, err := NewPartner(
				1,
				"Sample 1",
				"Sample 1",
				d.doc,
				[]byte(`{
					"type": "Polygon",
					"coordinates": [[
						[-46.720004081726074, -23.533084428991376],
						[-46.71858787536621, -23.533084428991376],
						[-46.71858787536621, -23.532041754244908],
						[-46.720004081726074, -23.532041754244908],
						[-46.720004081726074, -23.533084428991376]
					]]
				}`),
				[]byte(`{
					"type": "Point",
					"coordinates": [-46.71938180923462, -23.53242538082001]
				}`),
			)
			if d.err != nil && err == nil {
				t.Errorf("NewPartner: should've failed.")
			} else if d.err != nil && !errors.Is(err, d.err) {
				t.Errorf("NewPartner: expected error %v | got error %v", d.err, err)
			} else if d.err == nil && p.Document != d.expectedDoc {
				t.Errorf("NewPartner: expected document %s | got document %s", d.expectedDoc, p.Document)
			}
		})
	}
}

func TestNewPartnerCoverageArea(t *testing.T) {
	testCases := []struct {
		name         string
		coverageArea []byte
		err          error
	}{
		{
			"success",
			[]byte(`{
				"type": "Polygon",
				"coordinates": [[
					[-46.720004081726074, -23.533084428991376],
					[-46.71858787536621, -23.533084428991376],
					[-46.71858787536621, -23.532041754244908],
					[-46.720004081726074, -23.532041754244908],
					[-46.720004081726074, -23.533084428991376]
				]]
			}`),
			nil,
		},
		{
			"failuer",
			[]byte(`{
				"type": "Point",
				"coordinates": [-46.71938180923462, -23.53242538082001]
			}`),
			ErrorWrongCoverageArea,
		},
	}
	for _, d := range testCases {
		t.Run(d.name, func(t *testing.T) {
			_, err := NewPartner(
				1,
				"Sample 1",
				"Sample 1",
				"00.000.000/0000-00",
				d.coverageArea,
				[]byte(`{
					"type": "Point",
					"coordinates": [-46.71938180923462, -23.53242538082001]
				}`),
			)
			if d.err != nil && err == nil {
				t.Errorf("NewPartner: should've failed.")
			} else if d.err != nil && !errors.Is(err, d.err) {
				t.Errorf("NewPartner: expected error %v | got error %v", d.err, err)
			} else if d.err == nil && err != nil {
				t.Errorf("NewPartner: should've passed (got error %v)", err)
			}
		})
	}
}

func TestNewPartnerAddress(t *testing.T) {
	testCases := []struct {
		name    string
		address []byte
		err     error
	}{
		{
			"success",
			[]byte(`{
				"type": "Point",
				"coordinates": [-46.71938180923462, -23.53242538082001]
			}`),
			nil,
		},
		{
			"failure",
			[]byte(`{
				"type": "LineString",
				"coordinates": [
					[-46.71878635883331, -23.533384442363957],
					[-46.718958020210266, -23.53286802547826]
				]
			}`),
			ErrorWrongAddress,
		},
	}
	for _, d := range testCases {
		t.Run(d.name, func(t *testing.T) {
			_, err := NewPartner(
				1,
				"Sample 1",
				"Sample 1",
				"00.000.000/0000-00",
				[]byte(`{
					"type": "Polygon",
					"coordinates": [[
						[-46.720004081726074, -23.533084428991376],
						[-46.71858787536621, -23.533084428991376],
						[-46.71858787536621, -23.532041754244908],
						[-46.720004081726074, -23.532041754244908],
						[-46.720004081726074, -23.533084428991376]
					]]
				}`),
				d.address,
			)
			if d.err != nil && err == nil {
				t.Errorf("NewPartner: should've failed.")
			} else if d.err != nil && !errors.Is(err, d.err) {
				t.Errorf("NewPartner: expected error %v | got error %v", d.err, err)
			} else if d.err == nil && err != nil {
				t.Errorf("NewPartner: should've passed (got error %v)", err)
			}
		})
	}
}
