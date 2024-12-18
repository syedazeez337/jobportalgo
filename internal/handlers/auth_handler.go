package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syedazeez337/jobportalgo/internal/models"
	"github.com/syedazeez337/jobportalgo/internal/services"
)

func LoginHandler(db *sql.DB) gin.HandlerFunc {
	return func (c *gin.Context) {
		// var user models.User
	}
}

func RegisterHandler(db *sql.DB) gin.HandlerFunc {
	return func (c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		err := services.RegisterUser(db, &user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"error" : "Error creating user",
			})
		}

		c.JSON(http.StatusCreated, gin.H {
			"message": "User created successfully",
		})
	}
}
