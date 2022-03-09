package handlers

import (
	"todoList/platform"

	"github.com/gin-gonic/gin"
)

type itemRequest struct {
	Title string `json:"title"`
	Detail string `json:"detail"`
	Done bool `json:"done"`
}

func Post_todo (itemList *platform.ItemList) gin.HandlerFunc{
	return func (c *gin.Context) {
		requestBody := itemRequest{}
		c.Bind(&requestBody)

		item := platform.Item{
			Title: requestBody.Title,
			Detail: requestBody.Detail,
			Done: requestBody.Done,
		}



		itemList.Add(item)
		//itemList.Items = append(itemList.Items, requestBody)

		c.JSON(200, gin.H{
			"Title": requestBody.Title,
			"Detail": requestBody.Detail,
			"Done": requestBody.Done,
		})
	}
}