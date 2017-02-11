package app

import (
	"log"

	appCfg "../config/"
	controllers "../gin/controllers/"
	middleware "../gin/middleware/"

	"github.com/gin-gonic/gin"
)

//Run runs web app
func Run() {
	config, err := appCfg.ReadConfig()
	if err != nil {
		log.Fatalf("Can't start http server:%s.\n", err)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	v1 := router.Group("/v1")

	dicController := controllers.CreateDictionaryController(config)
	dicController.Register(v1)
	router.Use(middleware.ProcessResultsHandler)

	err = router.Run(":" + config.Port)
	if err != nil {
		log.Fatalf("Can't start http server:%s.\n", err)
	}
}
