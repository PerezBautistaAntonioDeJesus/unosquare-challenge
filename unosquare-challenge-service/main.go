package main

import (
	"log"
	"unosquare-challenge/api/configs"

	/*
		"context"
		"flag"
		"k8s.io/apimachinery/pkg/api/errors"
		metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
		"k8s.io/client-go/kubernetes"
		"k8s.io/client-go/tools/clientcmd"
		"k8s.io/client-go/util/homedir"
	*/

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"

	"unosquare-challenge/api"
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
