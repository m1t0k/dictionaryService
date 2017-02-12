package app

import (
	"log"

	config "../config"

	controllers "../gin/controllers/"
	middleware "../gin/middleware/"
	"github.com/gin-gonic/gin"
)

//Run runs web app
func Run() {

	settings, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("Can't start http server:%s.\n", err)
	}
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middleware.GlobalTraceLogger)
	//v1 := router.Group("/v1")
	controllers.RegisterDictionaryController(router)

	err = router.Run(":" + settings.Port)
	if err != nil {
		log.Fatalf("Can't start http server:%s.\n", err)
	}
}
