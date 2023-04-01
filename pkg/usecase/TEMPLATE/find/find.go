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

	TEMPLATE, err := findUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputFindTEMPLATEDto{}, status.InternalServerError, err
	}
	if TEMPLATE == nil {
		return OutputFindTEMPLATEDto{}, status.InvalidResourceId, err
	}

	if TEMPLATE.GetId() == "" {
		return OutputFindTEMPLATEDto{}, status.NoRecordsFound, err
	}

	outputFindTEMPLATEDto := OutputFindTEMPLATEDto{
		Id:        TEMPLATE.GetId(),
		Name:      TEMPLATE.GetName(),
		CreatedAt: TEMPLATE.GetCreatedAt().Format(time.RFC3339),
	}

	if !TEMPLATE.GetUpdatedAt().IsZero() {
		outputFindTEMPLATEDto.UpdatedAt = TEMPLATE.GetUpdatedAt().Format(time.RFC3339)
	}

	return outputFindTEMPLATEDto, status.Success, nil
}
