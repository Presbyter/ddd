package main

import (
	"github.com/Presbyter/ddd/router"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.Use(gin.Recovery())
	app = router.RouterBind(app)

	app.Run(":3000")
}
