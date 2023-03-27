package TEMPLATE

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

func Create(tenantId string, name string) (ITEMPLATE, error) {
	return NewTEMPLATE(uuid.NewV4().String(), tenantId, name, time.Now(), time.Time{})
}
