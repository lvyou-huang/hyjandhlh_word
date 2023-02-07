package model

type Column struct {
	Column_title        string `json:"column_title"`
	Column_intruduction string `json:"column_intruduction"`
	Phoneoremail        string `json:"phoneoremail"`
	Id                  int    `json:"id"`
	Cover               string `json:"cover"`
	Column_number       int    `json:"column_number"`
	Time                string `json:"time"`
}
