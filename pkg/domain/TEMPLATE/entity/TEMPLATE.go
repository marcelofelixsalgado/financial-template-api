package entity

import (
	"errors"
	"strings"
	"time"
)

type ITEMPLATE interface {
	GetId() string
	GetTenantId() string
	GetName() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type TEMPLATE struct {
	id        string
	tenantId  string
	name      string
	createdAt time.Time
	updatedAt time.Time
}

func NewTEMPLATE(id string, tenantId string, name string, createdAt time.Time, updatedAt time.Time) (ITEMPLATE, error) {
	TEMPLATE := &TEMPLATE{
		id:        id,
		tenantId:  tenantId,
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}

	TEMPLATE.format()
	if err := TEMPLATE.validate(); err != nil {
		return nil, err
	}

	return TEMPLATE, nil
}

func (TEMPLATE TEMPLATE) GetId() string {
	return TEMPLATE.id
}

func (TEMPLATE TEMPLATE) GetTenantId() string {
	return TEMPLATE.tenantId
}

func (TEMPLATE TEMPLATE) GetName() string {
	return TEMPLATE.name
}

func (TEMPLATE TEMPLATE) GetCreatedAt() time.Time {
	return TEMPLATE.createdAt
}

func (TEMPLATE TEMPLATE) GetUpdatedAt() time.Time {
	return TEMPLATE.updatedAt
}

func (TEMPLATE *TEMPLATE) format() {
	TEMPLATE.name = strings.TrimSpace(TEMPLATE.name)
}

func (TEMPLATE *TEMPLATE) validate() error {
	var idRequired error
	var tenantIdRequired error
	var nameRequired error
	var createdAtRequired error
	var upatedAtGreaterThanCreatedAt error

	if TEMPLATE.id == "" {
		idRequired = errors.New("id is required")
	}

	if TEMPLATE.tenantId == "" {
		tenantIdRequired = errors.New("tenant id is required")
	}

	if TEMPLATE.name == "" {
		nameRequired = errors.New("name is required")
	}

	if TEMPLATE.createdAt.IsZero() {
		createdAtRequired = errors.New("created at is required")
	}

	if !TEMPLATE.updatedAt.IsZero() && !TEMPLATE.createdAt.IsZero() && TEMPLATE.createdAt.After(TEMPLATE.updatedAt) {
		upatedAtGreaterThanCreatedAt = errors.New("upated at must be greater than the created at date")
	}

	return errors.Join(idRequired, tenantIdRequired, nameRequired, createdAtRequired, upatedAtGreaterThanCreatedAt)
}
