package update

type InputUpdateTEMPLATEDto struct {
	Id   string `json:"-"`
	Name string `json:"name"`
}

type OutputUpdateTEMPLATEDto struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
