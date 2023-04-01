package update

import (
	"time"

	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/domain/TEMPLATE/entity"
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/infrastructure/repository/TEMPLATE"
	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	repositoryStatus "github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"
)

type IUpdateUseCase interface {
	Execute(InputUpdateTEMPLATEDto) (OutputUpdateTEMPLATEDto, status.InternalStatus, error)
}

type UpdateUseCase struct {
	repository TEMPLATE.ITEMPLATERepository
}

func NewUpdateUseCase(repository TEMPLATE.ITEMPLATERepository) IUpdateUseCase {
	return &UpdateUseCase{
		repository: repository,
	}
}

func (updateUseCase *UpdateUseCase) Execute(input InputUpdateTEMPLATEDto) (OutputUpdateTEMPLATEDto, status.InternalStatus, error) {

	// Find the entity before update
	currentEntity, err := updateUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputUpdateTEMPLATEDto{}, status.InternalServerError, err
	}
	if currentEntity == nil {
		return OutputUpdateTEMPLATEDto{}, status.InvalidResourceId, nil
	}

	entity, err := entity.NewTEMPLATE(input.Id, currentEntity.GetTenantId(), input.Name, currentEntity.GetCreatedAt(), time.Now())
	if err != nil {
		return OutputUpdateTEMPLATEDto{}, status.InternalServerError, err
	}

	// Persists in dabatase
	repositoryInternalStatus, err := updateUseCase.repository.Update(entity)
	if repositoryInternalStatus == repositoryStatus.EntityWithSameKeyAlreadyExists {
		return OutputUpdateTEMPLATEDto{}, status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return OutputUpdateTEMPLATEDto{}, status.InternalServerError, err
	}

	outputUpdateTEMPLATEDto := OutputUpdateTEMPLATEDto{
		Id:        entity.GetId(),
		Name:      entity.GetName(),
		CreatedAt: entity.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: entity.GetUpdatedAt().Format(time.RFC3339),
	}

	return outputUpdateTEMPLATEDto, status.Success, nil
}
