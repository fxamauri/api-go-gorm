package repositories

import (
	"github.com/fxamauri/api-go-gorm/src/interfaces"
	"github.com/fxamauri/api-go-gorm/src/repositories/installment"
	"gorm.io/gorm"
)

type RepositoryContainer struct {
	InstallmentRepository interfaces.InstallmentRepository
}

func GetRepositories(db *gorm.DB) RepositoryContainer {
	return RepositoryContainer{
		InstallmentRepository: installment.NewInstallmentRepository(db),
	}
}
