package container

import (
	campaignModule "emailn/internal/campaign"
	"emailn/internal/core"
)

func InitInjector() {
	core.RegisterModuleBindings()
	campaignModule.RegisterModuleBindings()
}
