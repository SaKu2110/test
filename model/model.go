package model

type LoginRequest struct {
	ID			string
	PASSWORD	string
}

type User struct {
	ID			string
	PASSWORD	string
	ADMIN		string
}