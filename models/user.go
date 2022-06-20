package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// User structure of a user in lenslocked
type User struct {
	ID       int
	Email    string
	Password string
}

// UserService contains all methods and fields for interacting with the `users` table in the database.
type UserService struct {
	DB *sql.DB
}

// NewUser represents a new user with a plain text password.
type NewUser struct {
	Email string

	// plain text password
	Password string
}

// Create creates a new user and returns a pointer to the user.
func (us *UserService) Create(nu NewUser) (*User, error) {
	email := strings.ToLower(nu.Email)

	hashedPassword, err := hashPassword(nu.Password)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	user := User{
		Email:    email,
		Password: hashedPassword,
	}

	row := us.DB.QueryRow(createUserQuery, email, hashedPassword)
	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &user, nil
}

const createUserQuery = `
-- models/user.go:Create
INSERT INTO users (email, password)
VALUES ($1, $2)
RETURNING id;
`

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing: %w", err)
	}
	return string(hash), nil
}
