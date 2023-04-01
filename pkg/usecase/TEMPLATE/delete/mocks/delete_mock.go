package mocks

import (
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/usecase/TEMPLATE/delete"
	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
	"github.com/stretchr/testify/mock"
)

type DeleteUseCaseMock struct {
	mock.Mock
}

func (m *DeleteUseCaseMock) Execute(input delete.InputDeleteTEMPLATEDto) (delete.OutputDeleteTEMPLATEDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(delete.OutputDeleteTEMPLATEDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
