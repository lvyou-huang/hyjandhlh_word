package model

type Article struct {
	Article_title   string `json:"article_title"`
	Article_content string `json:"article_content"`
	Date            string `json:"date"`
	Category        string `json:"category"`
	Label           string `json:"label"`
	Column          string `json:"column"`
	Id              int    `json:"id"`
	Author          string `json:"author"`
	Author_id       int    `json:"author_id"`
	Like            int64  `json:"like"`
	View            int    `json:"view"`
	Comment         int    `json:"comment"`
	Postorboil      int    `json:"postorboil"`
	Cover           string `json:"cover"`
	Collection      int    `json:"collection"`
}
