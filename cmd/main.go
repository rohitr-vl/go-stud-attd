package main

import (
	route "clean-architecture/api/route"
	"clean-architecture/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.Psql

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, *db, gin)

	gin.Run(env.ServerAddress)
}
