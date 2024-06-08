package main

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"hero-backend/legacy"
	"time"
)

func SetupRoutes(engine *gin.Engine) {
	store := persistence.NewInMemoryStore(time.Second)

	// Legacy APIs
	engine.GET("/id", cache.CachePage(store, time.Hour*24, legacy.GetId))
	engine.GET("/awards", cache.CachePage(store, time.Minute, legacy.GetAwards))

	// New APIs
	v1 := engine.Group("/v1")
	{
		v1.GET("/csb_profile", cache.CachePage(store, time.Minute, GetCsbProfile))
	}
}
