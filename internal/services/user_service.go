package services

import (
	"database/sql"

	"github.com/syedazeez337/jobportalgo/internal/models"
	"github.com/syedazeez337/jobportalgo/internal/repository"
)

func GetUserByID(db *sql.DB, id int) (*models.User, error) {
	return repository.GetUserByID(db, id)
}