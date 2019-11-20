package models

//User struct declaration
type Transaction struct {
	ID       string `json:ID`
	Address  string `json:"Address" validate:"required,address"`
	County   string `json:"County" validate:"required,min=1,max=100"`
	District string `json:"District" validate:"required,min=1,max=100"`
}
