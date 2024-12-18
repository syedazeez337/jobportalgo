package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syedazeez337/jobportalgo/internal/models"
)

func LoginHandler(db *sql.DB) gin.HandlerFunc {
	return func (c *gin.Context) {
		// var user models.User
	}
}

func RegisterHandler(bd *sql.DB) gin.HandlerFunc {
	return func (c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		err := service
	}
}
