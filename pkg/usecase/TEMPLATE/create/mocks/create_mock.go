package mocks

import (
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/usecase/TEMPLATE/create"
	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
	"github.com/stretchr/testify/mock"
)

type CreateUseCaseMock struct {
	mock.Mock
}

func (m *CreateUseCaseMock) Execute(input create.InputCreateTEMPLATEDto) (create.OutputCreateTEMPLATEDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(create.OutputCreateTEMPLATEDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
