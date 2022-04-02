package services

import (
	"github.com/fxamauri/api-go-gorm/src/interfaces"
	"github.com/fxamauri/api-go-gorm/src/repositories"
	"github.com/fxamauri/api-go-gorm/src/services/installment"
)

type ServiceContainer struct {
	InstallmentService interfaces.InstallmentService
}

func GetServices(repositoryContainer repositories.RepositoryContainer) ServiceContainer {
	return ServiceContainer{
		InstallmentService: installment.NewInstallmentService(repositoryContainer),
	}
}
