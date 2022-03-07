package app

import (
	"github.com/gin-gonic/gin"
	"github.com/laithrafid/bookstore_utils-go/config_utils"
	"github.com/laithrafid/bookstore_utils-go/logger_utils"
)

var (
	router = gin.Default()
)

func StartApplication() {
	config, err := config_utils.LoadConfig(".")
	if err != nil {
		logger_utils.Error("cannot load config of application:", err)
	}
	mapUrls()

	logger_utils.Info("starting the application ....")
	router.Run(config.GnewsApiAddress)
}
