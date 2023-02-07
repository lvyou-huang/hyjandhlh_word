package model

type Comment struct {
	ID        uint   `json:"id"`
	ProjectId uint   `json:"projectid"`
	ParentId  uint   `json:"parentid"`
	UserId    string `json:"userid"`
	Content   string `json:"content"`
	Time      string `json:"time"`
}
