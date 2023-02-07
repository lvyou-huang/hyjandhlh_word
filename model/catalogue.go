package model

type Catalogue struct {
	Catalogue_id    int    `json:"catalogue_id"`
	Course_id       string `json:"course_id"`
	Catalogue_title string `json:"catalogue_title"`
	Trialornot      string `json:"trialornot"`
	Order           int    `json:"order"`
	Content         string `json:"content"`
}
