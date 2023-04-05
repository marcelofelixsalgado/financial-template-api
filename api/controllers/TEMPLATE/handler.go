package TEMPLATE

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/marcelofelixsalgado/financial-commons/api/requests"
	"github.com/marcelofelixsalgado/financial-commons/api/responses"
	"github.com/marcelofelixsalgado/financial-commons/api/responses/faults"

	"github.com/labstack/echo/v4"
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/usecase/TEMPLATE/create"
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/usecase/TEMPLATE/delete"
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/usecase/TEMPLATE/find"
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/usecase/TEMPLATE/list"
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/usecase/TEMPLATE/update"
	"github.com/marcelofelixsalgado/financial-commons/pkg/auth"
	"github.com/marcelofelixsalgado/financial-commons/pkg/commons/logger"
	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type ITEMPLATEHandler interface {
	CreateTEMPLATE(ctx echo.Context) error
	ListTEMPLATEs(ctx echo.Context) error
	GetTEMPLATEById(ctx echo.Context) error
	UpdateTEMPLATE(ctx echo.Context) error
	DeleteTEMPLATE(ctx echo.Context) error
}

type TEMPLATEHandler struct {
	createUseCase create.ICreateUseCase
	deleteUseCase delete.IDeleteUseCase
	findUseCase   find.IFindUseCase
	listUseCase   list.IListUseCase
	updateUseCase update.IUpdateUseCase
}

const requestBodyErrorMessage = "Error trying to read the request body: "
const inputConversionErrorMessage = "Error trying to convert the input data: "
const outputConversionErrorMessage = "Error trying to convert the output to response body: "
const unableFindEntityErrorMessage = "Unable to find the entity"

func NewTEMPLATEHandler(createUseCase create.ICreateUseCase, deleteUseCase delete.IDeleteUseCase, findUseCase find.IFindUseCase, listUseCase list.IListUseCase, updateUseCase update.IUpdateUseCase) ITEMPLATEHandler {
	return &TEMPLATEHandler{
		createUseCase: createUseCase,
		deleteUseCase: deleteUseCase,
		findUseCase:   findUseCase,
		listUseCase:   listUseCase,
		updateUseCase: updateUseCase,
	}
}

func (TEMPLATEHandler *TEMPLATEHandler) CreateTEMPLATE(ctx echo.Context) error {
	log := logger.GetLogger()

	tenantId, err := auth.ExtractUserId(ctx.Request())
	if err != nil {
		log.Errorf("Error extracting tenantId from access token: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	requestBody, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Warnf("%s%v", requestBodyErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	var input create.InputCreateTEMPLATEDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Warnf("%s%v", inputConversionErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	// Validating input parameters
	if responseMessage := ValidateCreateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	input.TenantId = tenantId

	output, internalStatus, err := TEMPLATEHandler.createUseCase.Execute(input)
	if internalStatus == status.EntityWithSameKeyAlreadyExists {
		log.Warnf("Error trying to create the entity - duplicate key: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.EntityWithSameKeyAlreadyExists, "body", "name", input.Name)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus != status.Success {
		log.Errorf("Error trying to create the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusCreated, output)
}

func (TEMPLATEHandler *TEMPLATEHandler) ListTEMPLATEs(ctx echo.Context) error {
	log := logger.GetLogger()

	tenantId, err := auth.ExtractUserId(ctx.Request())
	if err != nil {
		log.Errorf("Error extracting tenantId from access token: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	input := list.InputListTEMPLATEDto{
		TenantId: tenantId,
	}

	filterParameters, err := requests.SetupFilters(ctx.Request())
	if err != nil {
		log.Warnf("Error parsing the querystring parameters: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "query_parameter", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	output, internalStatus, err := TEMPLATEHandler.listUseCase.Execute(input, filterParameters)
	if internalStatus == status.InternalServerError {
		log.Errorf("Error listing the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.NoRecordsFound {
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(internalStatus, "", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusOK, output.TEMPLATES)
}

func (TEMPLATEHandler *TEMPLATEHandler) GetTEMPLATEById(ctx echo.Context) error {
	log := logger.GetLogger()

	id := ctx.Param("id")

	input := find.InputFindTEMPLATEDto{
		Id: id,
	}

	output, internalStatus, err := TEMPLATEHandler.findUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Errorf("Error finding the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.InvalidResourceId {
		log.Infof("%s", unableFindEntityErrorMessage)
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", id)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusOK, output)
}

func (TEMPLATEHandler *TEMPLATEHandler) UpdateTEMPLATE(ctx echo.Context) error {
	log := logger.GetLogger()

	id := ctx.Param("id")

	requestBody, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Warnf("%s%v", requestBodyErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	var input update.InputUpdateTEMPLATEDto

	if erro := json.Unmarshal([]byte(requestBody), &input); erro != nil {
		log.Warnf("%s%v", inputConversionErrorMessage, err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.MalformedRequest, "body", "", "")
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	input.Id = id

	// Validating input parameters
	if responseMessage := ValidateUpdateRequestBody(input).GetMessage(); responseMessage.ErrorCode != "" {
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	output, internalStatus, err := TEMPLATEHandler.updateUseCase.Execute(input)
	if internalStatus == status.EntityWithSameKeyAlreadyExists {
		log.Warnf("Error trying to create the entity - duplicate key: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByIssue(faults.EntityWithSameKeyAlreadyExists, "body", "name", input.Name)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.InternalServerError {
		log.Errorf("Error updating the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.InvalidResourceId {
		log.Infof("%s", unableFindEntityErrorMessage)
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", id)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.JSON(http.StatusOK, output)
}

func (TEMPLATEHandler *TEMPLATEHandler) DeleteTEMPLATE(ctx echo.Context) error {
	log := logger.GetLogger()

	id := ctx.Param("id")

	var input = delete.InputDeleteTEMPLATEDto{
		Id: id,
	}

	_, internalStatus, err := TEMPLATEHandler.deleteUseCase.Execute(input)
	if internalStatus == status.InternalServerError {
		log.Errorf("Error removing the entity: %v", err)
		responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.InternalServerError)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}
	if internalStatus == status.InvalidResourceId {
		log.Infof("%s", unableFindEntityErrorMessage)
		responseMessage := responses.NewResponseMessage().AddMessageByInternalStatus(status.InvalidResourceId, responses.PathParameter, "id", id)
		return ctx.JSON(responseMessage.HttpStatusCode, responseMessage)
	}

	return ctx.NoContent(http.StatusNoContent)
}
