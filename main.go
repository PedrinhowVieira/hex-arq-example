package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gitlab.com/bavatech/development/backend/retail/company-service/company"
	psql "gitlab.com/bavatech/development/backend/retail/company-service/repository"
)

func main() {
	loadEnv()

	companyToStore := &company.Company{
		Cnpj:        "38099278000130",
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

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
