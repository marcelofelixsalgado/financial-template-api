package template

import "time"

type TemplateCreated struct {
	Name    string
	Payload interface{}
}

func NewTemplateCreated() *TemplateCreated {
	return &TemplateCreated{
		Name: "TemplateCreated",
	}
}

func (event *TemplateCreated) GetName() string {
	return event.Name
}

func (event *TemplateCreated) GetDateTime() time.Time {
	return time.Now()
}

func (event *TemplateCreated) GetPayload() interface{} {
	return event.Payload
}

func (event *TemplateCreated) SetPayload(payload interface{}) {
	event.Payload = payload
}
