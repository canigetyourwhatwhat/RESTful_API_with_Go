package main

import (
	"fmt"
	"todoList/platform"
)

func main() {
	// r := gin.Default()

	// r.GET("/", handlers.Get_all_todo())

	// r.Run()

	todos := platform.New()
	fmt.Println(todos)
}