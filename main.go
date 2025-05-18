package main

import (
	"example.com/mygolangproj/db"
	"example.com/mygolangproj/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	db.InitDB()
	server.Run(":8080")
}
