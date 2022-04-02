package interfaces

import "github.com/fxamauri/api-go-gorm/src/structs"

type InstallmentService interface {
	Create(installment structs.Installment) error
}
type InstallmentRepository interface {
	Create(installment structs.Installment) error
}
