package app

import (
	"log"

	config "github.com/m1t0k/dictionaryService/api/infrastructure/v1/config"
	controllers "github.com/m1t0k/dictionaryService/api/infrastructure/v1/gin/controllers/"
	middleware "github.com/m1t0k/dictionaryService/api/infrastructure/v1/gin/middleware"

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
	//router.Use(gin.Logger)
	router.Use(middleware.GlobalTraceLogger)

	v1 := router.Group("/v1")
	dicController := controllers.CreateDictionaryController(settings)
	dicController.Register(v1)

	router.Use(middleware.ProcessResultsHandler)

	err = router.Run(":" + config.Port)
	if err != nil {
		log.Fatalf("Can't start http server:%s.\n", err)
	}
}
