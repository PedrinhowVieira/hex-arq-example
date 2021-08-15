package main

import (
	"fmt"
	"log"

	"gitlab.com/bavatech/development/backend/retail/company-service/company"
	psql "gitlab.com/bavatech/development/backend/retail/company-service/repository"
)

func main() {
	companyToStore := &company.Company{
		Cnpj:        "2323232",
		LegalName:   "legalname",
		DisplayName: "displayname",
	}

	repo, err := psql.NewPsqlRepository()
	if err != nil {
		log.Fatal(err)
	}
	service := company.NewCompanyService(repo)

	result, err := service.Store(companyToStore)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID from company Saved %v", result)
}
