package main

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce/bootstrap"
)

func main() {
	app := bootstrap.App()

	env := app.Env
	//
	//db := app.Mongo.Database(env.DBName)
	//
	//defer app.CloseDBConnection()
	//
	//timeout := time.Duration(env.ContextTimeout) * time.Second

	server := gin.Default()

	server.Run(env.ServerAddress)
}
