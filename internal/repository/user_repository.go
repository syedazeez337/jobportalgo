package repository

import (
	"database/sql"

	"github.com/syedazeez337/jobportalgo/internal/models"
)

func CreateUser(db *sql.DB, user *models.User) error {
	_, err := db.Exec(
		`INSERT INTO users (username, password, email) VALUES (?, ?, ?)`,
		user.Username, user.Password, user.Email,
	)

	return err	
}
