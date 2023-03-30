package TEMPLATE

import (
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/domain/TEMPLATE/entity"
	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/filter"
	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"
)

type ITEMPLATERepository interface {
	Create(entity.ITEMPLATE) (status.RepositoryInternalStatus, error)
	Update(entity.ITEMPLATE) (status.RepositoryInternalStatus, error)
	FindById(id string) (entity.ITEMPLATE, error)
	List(filterParameters []filter.FilterParameter, tenantId string) ([]entity.ITEMPLATE, error)
	Delete(id string) error
}
