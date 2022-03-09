package handlers

import (
	"todoList/platform"

	"github.com/gin-gonic/gin"
)

func Get_one_todo (itemList *platform.ItemList) gin.HandlerFunc{
	return func (c *gin.Context) {
		//picked_item := itemList.GetOneItem()
	}
}