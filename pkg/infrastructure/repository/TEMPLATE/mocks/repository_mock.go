package mocks

import (
	repositoryInternalStatus "github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"

	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/domain/TEMPLATE/entity"
	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/filter"
	"github.com/stretchr/testify/mock"
)

type TEMPLATERepositoryMock struct {
	mock.Mock
}

func (m *TEMPLATERepositoryMock) Create(TEMPLATE entity.ITEMPLATE) (repositoryInternalStatus.RepositoryInternalStatus, error) {
	args := m.Called(TEMPLATE)
	return repositoryInternalStatus.Success, args.Error(0)
}

func (m *TEMPLATERepositoryMock) Update(TEMPLATE entity.ITEMPLATE) (repositoryInternalStatus.RepositoryInternalStatus, error) {
	args := m.Called(TEMPLATE)
	return repositoryInternalStatus.Success, args.Error(0)
}

func (m *TEMPLATERepositoryMock) FindById(id string) (entity.ITEMPLATE, error) {
	args := m.Called(id)
	return args.Get(0).(entity.ITEMPLATE), args.Error(1)
}

func (m *TEMPLATERepositoryMock) List(filterParameters []filter.FilterParameter, tenantId string) ([]entity.ITEMPLATE, error) {
	args := m.Called(filterParameters, tenantId)
	return args.Get(0).([]entity.ITEMPLATE), args.Error(1)
}

func (m *TEMPLATERepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
