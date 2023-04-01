package delete

import (
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/infrastructure/repository/TEMPLATE"
	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type IDeleteUseCase interface {
	Execute(InputDeleteTEMPLATEDto) (OutputDeleteTEMPLATEDto, status.InternalStatus, error)
}

type DeleteUseCase struct {
	repository TEMPLATE.ITEMPLATERepository
}

func NewDeleteUseCase(repository TEMPLATE.ITEMPLATERepository) IDeleteUseCase {
	return &DeleteUseCase{
		repository: repository,
	}
}

func (deleteUseCase *DeleteUseCase) Execute(input InputDeleteTEMPLATEDto) (OutputDeleteTEMPLATEDto, status.InternalStatus, error) {

	// Find the entity before update
	entity, err := deleteUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputDeleteTEMPLATEDto{}, status.InternalServerError, err
	}
	if entity == nil {
		return OutputDeleteTEMPLATEDto{}, status.InvalidResourceId, nil
	}

	// Apply in dabatase
	err = deleteUseCase.repository.Delete(input.Id)
	if err != nil {
		return OutputDeleteTEMPLATEDto{}, status.InternalServerError, err
	}

	var outputDeleteTEMPLATEDto OutputDeleteTEMPLATEDto

	return outputDeleteTEMPLATEDto, status.Success, nil
}
