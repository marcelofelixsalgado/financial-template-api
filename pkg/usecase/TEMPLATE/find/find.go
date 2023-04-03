package find

import (
	"time"

	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/infrastructure/repository/TEMPLATE"
	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type IFindUseCase interface {
	Execute(InputFindTEMPLATEDto) (OutputFindTEMPLATEDto, status.InternalStatus, error)
}

type FindUseCase struct {
	repository TEMPLATE.ITEMPLATERepository
}

func NewFindUseCase(repository TEMPLATE.ITEMPLATERepository) IFindUseCase {
	return &FindUseCase{
		repository: repository,
	}
}

func (findUseCase *FindUseCase) Execute(input InputFindTEMPLATEDto) (OutputFindTEMPLATEDto, status.InternalStatus, error) {

	entity, err := findUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputFindTEMPLATEDto{}, status.InternalServerError, err
	}
	if entity == nil {
		return OutputFindTEMPLATEDto{}, status.InvalidResourceId, err
	}

	if entity.GetId() == "" {
		return OutputFindTEMPLATEDto{}, status.NoRecordsFound, err
	}

	outputFindTEMPLATEDto := OutputFindTEMPLATEDto{
		Id:        entity.GetId(),
		Name:      entity.GetName(),
		CreatedAt: entity.GetCreatedAt().Format(time.RFC3339),
	}

	if !entity.GetUpdatedAt().IsZero() {
		outputFindTEMPLATEDto.UpdatedAt = entity.GetUpdatedAt().Format(time.RFC3339)
	}

	return outputFindTEMPLATEDto, status.Success, nil
}
