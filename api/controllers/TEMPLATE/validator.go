package TEMPLATE

import (
	"github.com/marcelofelixsalgado/financial-commons/api/responses"
	"github.com/marcelofelixsalgado/financial-commons/api/responses/faults"

	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/usecase/TEMPLATE/create"
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/usecase/TEMPLATE/update"
)

type InputTEMPLATEDto struct {
	name string
}

func ValidateCreateRequestBody(inputCreateTEMPLATEDto create.InputCreateTEMPLATEDto) *responses.ResponseMessage {
	inputTEMPLATEDto := InputTEMPLATEDto{
		name: inputCreateTEMPLATEDto.Name,
	}
	return validateRequestBody(inputTEMPLATEDto)
}
func ValidateUpdateRequestBody(inputUpdateTEMPLATEDto update.InputUpdateTEMPLATEDto) *responses.ResponseMessage {
	if inputUpdateTEMPLATEDto.Id == "" {
		return responses.NewResponseMessage().AddMessageByIssue(faults.MissingRequiredField, responses.PathParameter, "id", "")
	}
	inputTEMPLATEDto := InputTEMPLATEDto{
		name: inputUpdateTEMPLATEDto.Name,
	}
	return validateRequestBody(inputTEMPLATEDto)
}

func validateRequestBody(inputTEMPLATEDto InputTEMPLATEDto) *responses.ResponseMessage {

	responseMessage := responses.NewResponseMessage()

	if inputTEMPLATEDto.name == "" {
		responseMessage.AddMessageByIssue(faults.MissingRequiredField, responses.Body, "name", "")
	}

	return responseMessage
}
