package company

import (
	"errors"
	"fmt"

	doc "github.com/paemuri/brdoc"
)

var (
	ErrInvalidFormat = errors.New("format invalid")
)

type companyService struct {
	companyRepo CompanyRepository
}

func NewCompanyService(companyRepo CompanyRepository) CompanyService {
	return &companyService{
		companyRepo,
	}
}

func (c *companyService) Find(id int64) (*Company, error) {
	return c.companyRepo.Find(id)
}

func (c *companyService) Store(company *Company) (int64, error) {
	if !doc.IsCNPJ(company.Cnpj) {
		return 0, fmt.Errorf("cnpj invalid: %v", ErrInvalidFormat)
	}
	return c.companyRepo.Store(company)
}
