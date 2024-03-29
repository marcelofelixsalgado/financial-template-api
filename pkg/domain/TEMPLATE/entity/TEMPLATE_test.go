package entity_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/domain/TEMPLATE/entity"
)

type testCase struct {
	tenantId string
	name     string
	expected string
}

func TestNewTEMPLATESuccess(t *testing.T) {

	testCases := []testCase{
		{
			tenantId: "TEMPLATE",
			name:     "TEMPLATE",
		},
	}

	for _, testCase := range testCases {
		received, err := entity.Create(testCase.tenantId, testCase.name)
		if err != nil {
			t.Errorf("Should not return an error: %s", err)
		}
		if testCase.name != received.GetName() {
			t.Errorf("Name expected: %s - received: %s", testCase.name, received.GetName())
		}
		if received.GetCreatedAt().IsZero() {
			t.Errorf("CreatedAt must not be zero")
		}
	}
}

func TestNewTEMPLATETrimSpaces(t *testing.T) {
	testCase := testCase{
		tenantId: "1234",
		name:     "     	TEMPLATE      ",
	}
	expectedName := "TEMPLATE"

	received, err := entity.Create(testCase.tenantId, testCase.name)
	if err != nil {
		t.Errorf("Should not return an error: %s", err)
	}
	if strings.Compare(expectedName, received.GetName()) != 0 {
		t.Errorf("Name expected: [%s] - received: [%s]", expectedName, received.GetName())
	}
}

func TestNewTEMPLATEInvalidName(t *testing.T) {
	testCase := testCase{
		tenantId: "1234",
		name:     "",
		expected: "name is required",
	}
	_, err := entity.Create(testCase.tenantId, testCase.name)

	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf(formatErrorDiff(testCase.expected, err))
	}
}

func formatErrorDiff(expected string, received error) string {
	return fmt.Sprintf("Error expected: %s - Error received: %s", expected, received)
}
