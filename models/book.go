package models

type Book struct {
	ID     int    `gorm:"primary_key" json:"id" from:"id"`
	Title  string `gorm:"not null" json:"title" form:"title" validate:"required~Title is required"`
	Author string `gorm:"not null" json:"author" form:"author" validate:"required~Author is required"`
}
