package app

import (
	"log"

	config "../config"

	controllers "../gin/controllers"
	middleware "../gin/middleware"
	"github.com/gin-gonic/gin"
)

//Run runs web app
func Run() {
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("Can't start http server:%s.\n", err)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.GlobalTraceLogger)

	v1Public := router.Group("/api/v1")
	controllers.RegisterDictionaryController(v1Public, config)

	v1Admin := router.Group("/api/admin/v1")
	controllers.RegisterAdminDictionaryController(v1Admin, config)

	err = router.Run(":" + config.Port)
	if err != nil {
		log.Fatalf("Can't start http server:%s.\n", err)
	}
}
