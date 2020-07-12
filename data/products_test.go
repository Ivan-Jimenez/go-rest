package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "fuckme",
		Price: 1.00,
		SKU:   "abs-asas-asas",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
