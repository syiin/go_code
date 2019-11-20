package models

import (
	"github.com/jinzhu/gorm"
)

//User struct declaration
type Admin struct {
	gorm.Model        //this is a struct embedding - kind of like inheritance
	Name       string `json:"Name"`
	Email      string `json:"Email"`
	Password   string `json:"Password" validate:"required"`
}
