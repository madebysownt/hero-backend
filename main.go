package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	SetupRoutes(r)

	if err := r.Run(); err != nil {
		log.Println(err)
	}

}
