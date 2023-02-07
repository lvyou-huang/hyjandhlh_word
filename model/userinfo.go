package model

type Userinfo struct {
	Name         string `json:"name"`
	Position     string `json:"position"`
	Company      string `json:"company"`
	Web          string `json:"web"`
	Introduce    string `json:"introduce"`
	Phoneoremail string `json:"phoneoremail"`
	Cover        string `json:"cover"`
}
