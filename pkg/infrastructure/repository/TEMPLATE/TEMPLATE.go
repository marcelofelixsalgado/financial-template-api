package TEMPLATE

import (
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/domain/TEMPLATE/entity"
	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/filter"
	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"
)

type TEMPLATERepository struct {
	client *sql.DB
}

type TEMPLATEModel struct {
	id        string
	tenantId  string
	name      string
	createdAt time.Time
	updatedAt time.Time
}

func NewTEMPLATERepository(client *sql.DB) ITEMPLATERepository {
	return &TEMPLATERepository{
		client: client,
	}
}

func (repository *TEMPLATERepository) Create(entity entity.ITEMPLATE) (status.RepositoryInternalStatus, error) {
	var mysqlErr *mysql.MySQLError

	model := TEMPLATEModel{
		id:        entity.GetId(),
		tenantId:  entity.GetTenantId(),
		name:      entity.GetName(),
		createdAt: entity.GetCreatedAt(),
		updatedAt: entity.GetUpdatedAt(),
	}

	statement, err := repository.client.Prepare("insert into TEMPLATES (id, tenant_id, name, created_at) values (?, ?, ?, ?)")
	if err != nil {
		return status.InternalServerError, err
	}
	defer statement.Close()

	_, err = statement.Exec(model.id, model.tenantId, model.name, model.createdAt)
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		// Unique key violated
		return status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return status.InternalServerError, err
	}

	return status.Success, nil
}

func (repository *TEMPLATERepository) Update(entity entity.ITEMPLATE) (status.RepositoryInternalStatus, error) {
	var mysqlErr *mysql.MySQLError

	model := TEMPLATEModel{
		id:        entity.GetId(),
		name:      entity.GetName(),
		updatedAt: entity.GetUpdatedAt(),
	}

	statement, err := repository.client.Prepare("update TEMPLATES set name = ?, updated_at = ? where id = ?")
	if err != nil {
		return status.InternalServerError, err
	}
	defer statement.Close()

	_, err = statement.Exec(model.name, model.updatedAt, model.id)
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		// Unique key violated
		return status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return status.InternalServerError, err
	}

	return status.Success, nil
}

func (repository *TEMPLATERepository) FindById(id string) (entity.ITEMPLATE, error) {

	row, err := repository.client.Query("select id, tenant_id, name, created_at, updated_at from TEMPLATES where id = ?", id)
	if err != nil {
		return entity.TEMPLATE{}, err
	}
	defer row.Close()

	var TEMPLATEModel TEMPLATEModel
	if row.Next() {
		if err := row.Scan(&TEMPLATEModel.id, &TEMPLATEModel.tenantId, &TEMPLATEModel.name, &TEMPLATEModel.createdAt, &TEMPLATEModel.updatedAt); err != nil {
			return entity.TEMPLATE{}, err
		}

		TEMPLATE, err := entity.NewTEMPLATE(TEMPLATEModel.id, TEMPLATEModel.tenantId, TEMPLATEModel.name, TEMPLATEModel.createdAt, TEMPLATEModel.updatedAt)
		if err != nil {
			return entity.TEMPLATE{}, err
		}
		return TEMPLATE, nil
	}
	return nil, nil
}

func (repository *TEMPLATERepository) List(filterParameters []filter.FilterParameter, tenantId string) ([]entity.ITEMPLATE, error) {
	nameFilter := ""
	for _, filterParameter := range filterParameters {
		switch filterParameter.Name {
		case "name":
			nameFilter = filterParameter.Value
		}
	}

	var rows *sql.Rows
	var err error
	if len(filterParameters) == 0 {
		rows, err = repository.client.Query("select id, tenant_id, name, created_at, updated_at from TEMPLATES where tenant_id = ?", tenantId)
	} else {
		if len(nameFilter) > 0 {
			rows, err = repository.client.Query("select id, tenant_id, name, created_at, updated_at from TEMPLATES where tenant_id = ? and name = ?", tenantId, nameFilter)
		}
	}

	if err != nil {
		return []entity.ITEMPLATE{}, err
	}
	defer rows.Close()

	TEMPLATES := []entity.ITEMPLATE{}
	for rows.Next() {
		var TEMPLATEModel TEMPLATEModel

		if err := rows.Scan(&TEMPLATEModel.id, &TEMPLATEModel.tenantId, &TEMPLATEModel.name, &TEMPLATEModel.createdAt, &TEMPLATEModel.updatedAt); err != nil {
			return []entity.ITEMPLATE{}, err
		}

		TEMPLATE, err := entity.NewTEMPLATE(TEMPLATEModel.id, TEMPLATEModel.tenantId, TEMPLATEModel.name, TEMPLATEModel.createdAt, TEMPLATEModel.updatedAt)
		if err != nil {
			return []entity.ITEMPLATE{}, err
		}

		TEMPLATES = append(TEMPLATES, TEMPLATE)
	}

	return TEMPLATES, nil
}

func (repository *TEMPLATERepository) Delete(id string) error {

	statement, err := repository.client.Prepare("delete from TEMPLATES where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
