package main

import (
	"fmt"

	"github.com/1Vewton/WikiSpider/utils/config"
	"github.com/1Vewton/WikiSpider/utils/logger"
)

var service_logger = logger.NewLogger("WikiSpider")

func main() {
	service_logger.Info("Program started")
	service_logger.Info(fmt.Sprintf("Target URL: %s", config.GetWikiUrl()))
}
