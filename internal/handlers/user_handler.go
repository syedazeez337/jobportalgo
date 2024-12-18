package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/syedazeez337/jobportalgo/internal/services"
)

func GetUserByIDHandler(db *sql.Db) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"error": "Invalid user ID",
			})
		}

		user, err := services.GetUserByID(db, id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, user)
	}
}