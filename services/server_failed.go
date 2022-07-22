package services

import (
	"fmt"
	"go-reporting-server/app"
	"go-reporting-server/entity"
	"strings"

	log "github.com/sirupsen/logrus"
)

func IfServerShutdown(start int, end int) {
	var servers []entity.ReportingServer
	var serverDead []string

	log.Infof(fmt.Sprintf("server jobs is running, Start: %d End: %d", start, end))

	dataServer := []string{"SERVER-10", "SERVER-13"}

	app.Instance.Raw("SELECT *  FROM `reporting_servers` WHERE `created_at` >= ? AND `created_at` <= ?", start, end).Scan(&servers)

	for _, serverName := range dataServer {
		count := 0
		for _, serverLog := range servers {
			if serverLog.ServerName == serverName {
				count++
			}
		}

		if count == 0 {
			serverDead = append(serverDead, serverName)
		}
	}

	if serverDead != nil {
		message := strings.Join(serverDead, ", ")
		SendMessageTelegram(message)
	}

	serverDead = nil
}
