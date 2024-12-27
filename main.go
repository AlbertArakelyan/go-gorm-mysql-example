package main

import (
	"go-mysql-phpmyadmin/db"
	"go-mysql-phpmyadmin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
