package model

type Course struct {
	Introduction string `json:"introduction"`
	Course_title string `json:"course_title"`
	Abstract     string `json:"abstract"`
	Course_id    int    `json:"course_id"`
	Cover        string `json:"cover"`
	Catalogue    string `json:"catalogue"`
	Hot          int    `json:"hot"`
}
