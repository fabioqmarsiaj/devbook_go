package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

// NewUserRepositories creates a repository of users
func NewUserRepositories(db *sql.DB) *users {
	return &users{db}
}

// insert an user into DB
func (u users) Create(user models.User) (uint64, error) {
	return 0, nil
}