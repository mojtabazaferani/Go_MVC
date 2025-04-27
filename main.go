package main

import (
	"mvc/routes"
	"github.com/gin-gonic/gin"
) 

func main() {

	r := gin.Default()

	routes.Routes(r)

	r.Run(":8080")

}