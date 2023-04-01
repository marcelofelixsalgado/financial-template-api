package mocks

import (
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/usecase/TEMPLATE/update"
	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
	"github.com/stretchr/testify/mock"
)

type UpdateUseCaseMock struct {
	mock.Mock
}

func (m *UpdateUseCaseMock) Execute(input update.InputUpdateTEMPLATEDto) (update.OutputUpdateTEMPLATEDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(update.OutputUpdateTEMPLATEDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
