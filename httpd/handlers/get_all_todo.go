package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get_all_todo() gin.HandlerFunc {
	return func (c *gin.Context){		
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	}
}