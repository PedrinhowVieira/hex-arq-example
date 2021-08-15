package company

import (
	"errors"
	"testing"
)

func createMockCompany(id *int64) *Company {
	return &Company{
		Id:          id,
		Cnpj:        "38099278000130",
		DisplayName: "displayname",
		LegalName:   "legalname",
	}
}

type companyRepositoryMock struct{}

func (c *companyRepositoryMock) Store(company *Company) (int64, error) {
	id := int64(123)
	company.Id = &id
	return *company.Id, nil
}

func (p *companyRepositoryMock) Find(id int64) (*Company, error) {
	return createMockCompany(&id), nil
}

func TestPersonService(t *testing.T) {
	t.Run("Test store company", func(t *testing.T) {
		var repo CompanyRepository = &companyRepositoryMock{}
		companyService := NewCompanyService(repo)

		companyMock := createMockCompany(nil)

		got, _ := companyService.Store(companyMock)
		want := int64(123)

		if got != want {
			t.Errorf("\ngot: %v \nwnt: %v", got, want)
		}
	})

	t.Run("Test store company with invalid CPF", func(t *testing.T) {
		var repo CompanyRepository = &companyRepositoryMock{}
		companyService := NewCompanyService(repo)

		companyMock := createMockCompany(nil)
		companyMock.Cnpj = ""

		_, got := companyService.Store(companyMock)
		want := ErrInvalidFormat

		if !errors.As(got, &want) {
			t.Errorf("\ngot: %v \nwnt: %v", got, want)
		}
	})
}
