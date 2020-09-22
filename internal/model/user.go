package model

type User struct {
	BasicModel
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
