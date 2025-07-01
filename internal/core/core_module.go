package core

import (
	di "emailn/cmd/container/injector"
	"emailn/internal/core/infra/database"
)

func RegisterModuleBindings() {
	di.Register[database.DatabaseInterface](&database.PostgresInterfaceImpl{})
}
