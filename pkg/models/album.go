package models

type Album struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Singer string `json:"singer"`
	Year   int    `json:"year"`
}
