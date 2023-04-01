package create

type InputCreateTEMPLATEDto struct {
	TenantId string
	Name     string `json:"name"`
}

type OutputCreateTEMPLATEDto struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}
