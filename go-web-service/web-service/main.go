package main

import (
    "task-management/web-service/router"
)

func main() {
	engine := router.RouterSetup() 

	engine.Run("localhost:8080")
}
