package main

import (
	"emailn/cmd/api"
	"emailn/cmd/container"
	"emailn/db/migration"
	"emailn/internal/shared/validator"
	"net/http"
)

func main() {
	migration.RunMigrations()
	validator.ConfigValidator()
	container.InitInjector()
	http.ListenAndServe(":3000", api.NewRouter())
}
