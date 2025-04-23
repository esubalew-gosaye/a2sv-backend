package main

import (
	"task-management/web-service/db"
	"task-management/web-service/router"
)

func main() {
	db.ConnectDB()

	engine := router.RouterSetup()
	engine.Run("localhost:8080")
}