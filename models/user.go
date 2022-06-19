package models

import "database/sql"

type User struct {
	ID       int
	Email    string
	Password string
}

type UserService struct {
	DB *sql.DB
}
