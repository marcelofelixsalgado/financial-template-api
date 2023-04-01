package mocks

import (
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/usecase/TEMPLATE/list"
	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/filter"
	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
	"github.com/stretchr/testify/mock"
)

type ListUseCaseMock struct {
	mock.Mock
}

func (m *ListUseCaseMock) Execute(input list.InputListTEMPLATEDto, filterParameters []filter.FilterParameter) (list.OutputListTEMPLATEDto, status.InternalStatus, error) {
	args := m.Called(input, filterParameters)
	return args.Get(0).(list.OutputListTEMPLATEDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
