package campaign

import (
	di "emailn/cmd/container/injector"
	campaignApplication "emailn/internal/campaign/application"
	campaignRepository "emailn/internal/campaign/domain/repository"
	campaignRepositoryImpl "emailn/internal/campaign/infra/repository"
	campaignPresentation "emailn/internal/campaign/presentation"
	"emailn/internal/core/infra/database"
)

func RegisterModuleBindings() {
	/// Repositories
	di.Register[campaignRepository.Repository](&campaignRepositoryImpl.PostgresRepository{
		DbInterface: di.Resolve[database.DatabaseInterface](),
	})

	/// Services
	di.Register(campaignApplication.Service{
		Repository: di.Resolve[campaignRepository.Repository](),
	})

	/// Handler
	di.Register(campaignPresentation.Handler{
		Service: di.Resolve[campaignApplication.Service](),
	})
}
