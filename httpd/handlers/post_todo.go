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
		// requestBody := itemRequest{}
		// c.Bind(&requestBody)

		// item := platform.Item{
		// 	ID: requestBody.ID,
		// 	Title: requestBody.Title,			
		// 	Done: requestBody.Done,
		// }

		id := c.Param("id")
		title := c.Param("title")

		item := platform.Item{
			ID: id,
			Title: title,
			Done: false,
		}

		itemList.Add(item)

		c.JSON(200, item)
	}
}