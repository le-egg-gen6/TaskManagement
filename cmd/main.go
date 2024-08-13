package main

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce/bootstrap"
	"time"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)

	defer app.CloseDBConnection()

	timout := time.Duration(env.ContextTimeout) * time.Second

	server := gin.Default()

	server.Run(env.ServerAddress)
}
