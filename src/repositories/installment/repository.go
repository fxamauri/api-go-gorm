package installment

import (
	"github.com/fxamauri/api-go-gorm/src/structs"
	"gorm.io/gorm"
)

type InstallmentRepository struct {
	db *gorm.DB
}

func NewInstallmentRepository(db *gorm.DB) InstallmentRepository {
	return InstallmentRepository{
		db: db,
	}
}

func (ir InstallmentRepository) Create(installment structs.Installment) error {
	return ir.db.Create(installment).Error
}
