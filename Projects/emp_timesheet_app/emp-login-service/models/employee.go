package models

type Employee struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
}
