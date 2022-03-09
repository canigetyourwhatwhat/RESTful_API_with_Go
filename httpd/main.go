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

	//r.GET("/", handlers.Get_one_todo(db))
	r.GET("/", handlers.Get_all_todo(db))
	r.POST("/post", handlers.Post_todo(db))
	//r.PUT("/", handlers.Put_todo())	
	//r.DELETE("/", handlers.Delete_todo())
	

	r.Run()

}