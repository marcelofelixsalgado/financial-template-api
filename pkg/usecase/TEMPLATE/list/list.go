package list

import (
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/infrastructure/repository/TEMPLATE"
	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/filter"
	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type IListUseCase interface {
	Execute(InputListTEMPLATEDto, []filter.FilterParameter) (OutputListTEMPLATEDto, status.InternalStatus, error)
}

type ListUseCase struct {
	repository TEMPLATE.ITEMPLATERepository
}

func NewListUseCase(repository TEMPLATE.ITEMPLATERepository) IListUseCase {
	return &ListUseCase{
		repository: repository,
	}
}

func (listUseCase *ListUseCase) Execute(input InputListTEMPLATEDto, filterParameters []filter.FilterParameter) (OutputListTEMPLATEDto, status.InternalStatus, error) {

	TEMPLATEs, err := listUseCase.repository.List(filterParameters, input.TenantId)
	if err != nil {
		return OutputListTEMPLATEDto{}, status.InternalServerError, err
	}

	outputListTEMPLATEDto := OutputListTEMPLATEDto{}

	if len(TEMPLATEs) == 0 {
		return outputListTEMPLATEDto, status.NoRecordsFound, nil
	}

	for _, item := range TEMPLATEs {
		TEMPLATE := template{
			Id:   item.GetId(),
			Name: item.GetName(),
		}
		outputListTEMPLATEDto.TEMPLATES = append(outputListTEMPLATEDto.TEMPLATES, TEMPLATE)
	}
	return outputListTEMPLATEDto, status.Success, nil
}
