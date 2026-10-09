package main

import (
	"fmt"

	"github.com/1Vewton/WikiSpider/datasource"
	"github.com/1Vewton/WikiSpider/target"
	"github.com/1Vewton/WikiSpider/ui/cmd/operations"
	"github.com/1Vewton/WikiSpider/utils/config"
	"github.com/1Vewton/WikiSpider/utils/logger"
)

var service_logger = logger.NewLogger("WikiSpider")

func main() {
	var dataSourceUsed datasource.DataSourceType
	service_logger.Info("Program started")
	service_logger.Info(
		fmt.Sprintf(
			"Target URL: %s",
			config.GetWikiURL(),
		),
	)
	operations.Welcome()
	resDataSource := datasource.GetDataSource(
		config.GetWikiURL(),
	)
	// Decide data source
	useAutomDecidedDataSource := operations.YesOrNo(
		true,
		fmt.Sprintf(
			"The program decide to use data source type %s, is that correct?",
			resDataSource.ToString(),
		),
	)
	if useAutomDecidedDataSource {
		dataSourceUsed = resDataSource
	} else {
		fmt.Println("Here are the available data sources supported, decide it by your own")
		supportedDataSources := datasource.GetAllSuportedTypes()
		for _, dataSource := range supportedDataSources {
			fmt.Printf(
				"- %s\n",
				dataSource.ToString(),
			)
		}
		var selectedDataSource string
		_, err := fmt.Scan(&selectedDataSource)
		if err != nil {
			service_logger.Error(err.Error())
			panic(err)
		}
		dataSourceUsed = datasource.ToDataSourceType(
			selectedDataSource,
		)
	}
	if dataSourceUsed == datasource.NotSupported {
		service_logger.Error("data source not supported!")
		panic("data source not supported!")
	}
	fmt.Printf(
		"The program will use data source %s\n",
		dataSourceUsed.ToString(),
	)
	// Decide target
	fmt.Println(
		"Here are the available targets to store data: ",
	)
	for _, target := range target.GetAllTypes() {
		fmt.Println(target.ToString())
	}
	var target string
	fmt.Println(
		"Input the name of the datasource you want to use: ",
	)
	_, err := fmt.Scan(&target)
	if err != nil {
		service_logger.Error(err.Error())
		panic(err.Error())
	}
}
