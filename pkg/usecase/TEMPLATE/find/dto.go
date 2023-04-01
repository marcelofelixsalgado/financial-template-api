package find

type InputFindTEMPLATEDto struct {
	Id string
}

type OutputFindTEMPLATEDto struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
