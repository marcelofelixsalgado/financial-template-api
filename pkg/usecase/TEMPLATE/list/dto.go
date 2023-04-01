package list

type InputListTEMPLATEDto struct {
	TenantId string
}

type template struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type OutputListTEMPLATEDto struct {
	TEMPLATES []template `json:"-"`
}
