package company

import "fmt"

type Company struct {
	Id          *int64 `json:"id"`
	Cnpj        string `json:"cnpj"`
	LegalName   string `json:"legalName"`
	DisplayName string `json:"displayName"`
}

func (c Company) String() string {
	return fmt.Sprintf(
		"\nId: %v\nCnpj: %v\nLegalName: %v\nDisplayName: %v",
		*c.Id,
		c.Cnpj,
		c.LegalName,
		c.DisplayName,
	)
}
