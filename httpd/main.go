package main

import (
	"todoList/httpd/handlers"
	"todoList/platform"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// I will modify here later on to add the real DB!!
	db := platform.New()

	r.GET("/:id", handlers.Get_one_todo(db))            // get one specific item
	r.GET("/", handlers.Get_all_todo(db))               // get all todo items
	r.POST("/post/:id/:title", handlers.Post_todo(db))  // add one todo item
	r.PUT("/put/:id", handlers.Put_todo(db))            // modify one specific todo item
	r.DELETE("/delete/:id", handlers.Delete_todo(db))  // delete the specific item

	r.Run()

}