package app

import (
	"log"

	appCfg "../config/"
	ctrl "../gin/controllers/"
	"github.com/gin-gonic/gin"
)

//Run runs web app
func Run() {
	err := appCfg.ReadConfig()
	if err != nil {
		log.Fatalf("Can't start http server:%s.\n", err)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	v1 := router.Group("/v1")

	dicController := ctrl.DictionaryController{}
	dicController.Register(v1)

	err = router.Run(":8999")
	if err != nil {
		log.Fatalf("Can't start http server:%s.\n", err)
	}
}
