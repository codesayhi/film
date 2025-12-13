package bootstrap

import (
	domaincountry "github.com/codesayhi/golang-clean/internal/domain/country"
	httpcountry "github.com/codesayhi/golang-clean/internal/http/country"
	gormcountry "github.com/codesayhi/golang-clean/internal/infrastructure/db/gormrepo/country"
	gosimpleslug "github.com/codesayhi/golang-clean/internal/infrastructure/slug/gosimple"
	usecasecountry "github.com/codesayhi/golang-clean/internal/service/country"
	"gorm.io/gorm"
)

type Application struct {
	CountryHandler *httpcountry.Handler
}

func NewApplication(db *gorm.DB) *Application {
	//======================================================================================================================
	//Interface Other
	slugGen := gosimpleslug.New()

	//======================================================================================================================
	// Repository (GORM)
	var countryRepo domaincountry.Repository = gormcountry.NewCountryRepository(db)

	//======================================================================================================================
	// Usecase
	countrySvc := usecasecountry.NewService(countryRepo, slugGen)

	//======================================================================================================================
	// HTTP handler
	countryHandler := httpcountry.NewHandler(countrySvc)

	//======================================================================================================================
	return &Application{
		CountryHandler: countryHandler,
	}
}
