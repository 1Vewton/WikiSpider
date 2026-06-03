package main

import (
	"fmt"

	"github.com/1Vewton/WikiSpider/utils/config"
	"github.com/1Vewton/WikiSpider/utils/logger"
)

var service_name string = "WikiSpider"

func main() {
	logger.Info("Program started", service_name)
	logger.Info(fmt.Sprintf("Target URL: %s", config.GetWikiUrl()), service_name)
}
