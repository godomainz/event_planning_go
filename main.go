package main

import (
	"event_planning_go/db"
	"event_planning_go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8081")
}


