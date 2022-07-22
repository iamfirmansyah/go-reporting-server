package main

import (
	"go-reporting-server/app"
	"go-reporting-server/config"
	"go-reporting-server/services"
)

func main() {
	// * Load Configuration
	config.LoadAppConfig()

	// * Initialize Database
	app.ConnectMysql(config.AppConfig.MYSQL_CONNECTION)

	// * runing cronjobs
	services.RunCronJobs()
}
