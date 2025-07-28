package main

import (
	"task_management/Delivery/config"
	route "task_management/Delivery/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.Config{}
	config.LoadEnv()
	database := config.Connect_db()
	router := gin.Default()
	route.Setup(*database, router)
}
