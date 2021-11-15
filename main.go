package main

import (
	"github.com/Presbyter/ddd/gateway"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.Use(gin.Recovery())
	app = gateway.RouterBind(app)

	app.Run(":3000")
}
