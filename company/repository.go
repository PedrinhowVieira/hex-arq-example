package company

type CompanyRepository interface {
	Find(id int64) (*Company, error)
	Store(company *Company) (int64, error)
}
