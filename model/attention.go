package model

type Attention struct {
	Followed int    `json:"followed"`
	Follow   int    `json:"follow"`
	Date     string `json:"date"`
	Status   string `json:"status"`
}
