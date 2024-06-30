package model

type Todo struct {
	Base
	Title string `json:"title" gorm:"type:varchar(64) not null"`
}
