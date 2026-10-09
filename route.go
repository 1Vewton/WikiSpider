package main

import (
	"fmt"

	"github.com/1Vewton/WikiSpider/datasource"
	"github.com/1Vewton/WikiSpider/datasource/fandom"
	"github.com/1Vewton/WikiSpider/datasource/wikipedia"
	"github.com/1Vewton/WikiSpider/target"
)

// NewTarget creates new target
func NewTarget(
	targetType target.TargetType,
) (target.Target, error) {
	switch targetType {
	default:
		return nil, fmt.Errorf(
			"%s does not supported",
			targetType.ToString(),
		)
	}
}

// NewDataSourceTask creates new data source task
func NewDataSourceTask(
	dataSourceType datasource.DataSourceType,
	title string,
	retryCount int,
	linkLimit int,
	userAgent string,
	targetURL string,
) (datasource.DataSource, error) {
	switch dataSourceType {
	case datasource.Fandom:
		return fandom.New(
			title,
			retryCount,
			linkLimit,
			userAgent,
			targetURL,
		), nil
	case datasource.Wikipedia:
		return wikipedia.New(
			title,
			retryCount,
			linkLimit,
			userAgent,
			targetURL,
		), nil
	default:
		return nil, fmt.Errorf(
			"data source type %s not supported",
			dataSourceType,
		)
	}
}
