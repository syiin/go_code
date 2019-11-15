package models

import (
// "github.com/jinzhu/gorm"
)

//User struct declaration
type Transactions struct {
	// gorm.Model        //this is a struct embedding - kind of like inheritance
	ID       string `json:ID`
	Address  string `json:"Address"`
	County   string `json:"County"`
	District string `json:"District"`
}
