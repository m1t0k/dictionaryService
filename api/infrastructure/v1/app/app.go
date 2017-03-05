package app

import (
	"log"
	"time"

	config "../config"
	controllers "../gin/controllers"
	middleware "../gin/middleware"
	"github.com/m1t0k/gin-cors"
	"gopkg.in/gin-gonic/gin.v1"
)

//Run runs web app
func Run() {
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("Can't start http server:%s.\n", err)
	}

	router := gin.New()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
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
