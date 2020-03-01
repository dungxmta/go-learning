package repository

import (
	models "testProject/learning/example/db_postgres_with_repository_pattern/model"
)

type UserRepo interface {
	// Select() ([]*models.User, error)
	Select() ([]models.User, error)
	Insert(u models.User) error
	Exists(key, value string) bool
}
