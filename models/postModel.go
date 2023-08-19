package models

import (

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string `json:"title" gorm:"not null" validate:"required"`
	Body  string `json:"body" gorm:"not null" validate:"required"`
}

func (p *Post) Validate() error {
	return validator.New().Struct(p)
}
