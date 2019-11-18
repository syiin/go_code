package models

//User struct declaration
type Transaction struct {
	ID       string `json:ID`
	Address  string `json:"Address"`
	County   string `json:"County"`
	District string `json:"District"`
}
