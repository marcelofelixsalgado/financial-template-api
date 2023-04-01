package create

import (
	"time"

	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/domain/TEMPLATE/entity"
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/infrastructure/repository/TEMPLATE"
	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	repositoryStatus "github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"
)

type ICreateUseCase interface {
	Execute(InputCreateTEMPLATEDto) (OutputCreateTEMPLATEDto, status.InternalStatus, error)
}

type CreateUseCase struct {
	repository TEMPLATE.ITEMPLATERepository
}

func NewCreateUseCase(repository TEMPLATE.ITEMPLATERepository) ICreateUseCase {
	return &CreateUseCase{
		repository: repository,
	}
}

func (createUseCase *CreateUseCase) Execute(input InputCreateTEMPLATEDto) (OutputCreateTEMPLATEDto, status.InternalStatus, error) {

	// Creates an entity
	entity, err := entity.Create(input.TenantId, input.Name)
	if err != nil {
		return OutputCreateTEMPLATEDto{}, status.InternalServerError, err
	}

	// Persists in dabatase
	repositoryInternalStatus, err := createUseCase.repository.Create(entity)
	if repositoryInternalStatus == repositoryStatus.EntityWithSameKeyAlreadyExists {
		return OutputCreateTEMPLATEDto{}, status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return OutputCreateTEMPLATEDto{}, status.InternalServerError, err
	}

	outputCreateTEMPLATEDto := OutputCreateTEMPLATEDto{
		Id:        entity.GetId(),
		Name:      entity.GetName(),
		CreatedAt: entity.GetCreatedAt().Format(time.RFC3339),
	}

	return outputCreateTEMPLATEDto, status.Success, nil
}
