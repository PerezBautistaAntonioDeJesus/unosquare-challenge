package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"

	"unosquare-challenge/api"
	"unosquare-challenge/api/configs"
)

func main() {
	engine := gin.New()
	engine.Use(cors.New(configs.Cors))
	engine.SetTrustedProxies(nil)

	engine.Use(static.Serve("/swagger", static.LocalFile("api/swagger", true)))
	ops := middleware.RedocOpts{SpecURL: "swagger/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)
	engine.GET("docs", gin.WrapH(sh))

	api.AddRoutes(engine)
	log.Println("listening on: http://localhost:8080")
	if err := engine.Run(":8080"); err != nil {
		log.Fatalln("couldn't start service: ", err)
	}

}
