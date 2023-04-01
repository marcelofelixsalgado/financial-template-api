package mocks

import (
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/usecase/TEMPLATE/find"
	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
	"github.com/stretchr/testify/mock"
)

type FindUseCaseMock struct {
	mock.Mock
}

func (m *FindUseCaseMock) Execute(input find.InputFindTEMPLATEDto) (find.OutputFindTEMPLATEDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(find.OutputFindTEMPLATEDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
