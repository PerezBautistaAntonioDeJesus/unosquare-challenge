package configs

import (
	"time"

	"github.com/gin-contrib/cors"
)

var Cors = cors.Config{
	AllowAllOrigins:  true,
	AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
	AllowHeaders:     []string{"Authorization", "Accept", "Origin", "Content-Length", "Content-Type"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
	MaxAge:           12 * time.Hour,
}
