package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/syedazeez337/jobportalgo/internal/handlers"
)

func InitRoutes(r *gin.Engine, db *sql.DB)  {
	// Auth Routes
	r.POST("/login", handlers.LoginHandler(db))
	r.POST("/register", handlers.RegisterHandler(db))
}