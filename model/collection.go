package model

type Collection struct {
	Id           int    `json:"id"`
	Article_id   int    `json:"article_id"`
	Collector_id int    `json:"collector_id"`
	Date         string `json:"date"`
}
