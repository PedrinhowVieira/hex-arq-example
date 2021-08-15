package psql

import (
	"fmt"

	"gitlab.com/bavatech/development/backend/retail/company-service/company"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type psqlRepository struct {
	client *gorm.DB
}

type psqlClientConfigs struct {
	host     string
	user     string
	dbname   string
	password string
	port     string
}

func (pg psqlClientConfigs) String() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		pg.host,
		pg.user,
		pg.password,
		pg.dbname,
		pg.port,
	)
}

func newPsqlClient() (*gorm.DB, error) {
	psqlConfigs := psqlClientConfigs{
		host:     "localhost",
		user:     "tupi",
		password: "ancora",
		dbname:   "companydb",
		port:     "5432",
	}
	gormConfig := gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
	psqlClient, err := gorm.Open(postgres.Open(psqlConfigs.String()), &gormConfig)
	if err != nil {
		return nil, err
	}
	return psqlClient, nil
}

func NewPsqlRepository() (company.CompanyRepository, error) {
	repo := &psqlRepository{}
	client, err := newPsqlClient()
	if err != nil {
		return nil, fmt.Errorf("%v: repository error", err)
	}
	repo.client = client
	return repo, nil
}

func (psql *psqlRepository) Find(id int64) (*company.Company, error) {
	var c company.Company
	result := psql.client.First(&c, id)
	if result.Error != nil {
		return nil, fmt.Errorf("%v: repository error", result.Error)
	}
	return &c, nil
}

func (psql *psqlRepository) Store(c *company.Company) (int64, error) {
	result := psql.client.Create(&c)
	if result.Error != nil {
		return 0, fmt.Errorf("%v: repository error", result.Error)
	}
	return *c.Id, nil
}
