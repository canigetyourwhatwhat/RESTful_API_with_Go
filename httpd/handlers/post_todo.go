package handlers

import (
	"todoList/platform"

	"github.com/gin-gonic/gin"
)

type itemRequest struct {
	ID string `json:"id"`
	Title string `json:"title"`	
	Done bool `json:"done"`
}

func Post_todo (itemList *platform.ItemList) gin.HandlerFunc{
	return func (c *gin.Context) {
		requestBody := itemRequest{}
		c.Bind(&requestBody)

		item := platform.Item{
			ID: requestBody.ID,
			Title: requestBody.Title,			
			Done: requestBody.Done,
		}

		itemList.Add(item)

		c.JSON(200, item)
	}
}