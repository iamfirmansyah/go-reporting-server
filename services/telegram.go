package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-reporting-server/config"
	"go-reporting-server/helper"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	Token  string
	ChatId string
)

func getUrl() string {
	Token = config.AppConfig.TELEGRAM_BOT_ID
	return fmt.Sprintf("https://api.telegram.org/bot%s", Token)
}

func SendMessageTelegram(text string) (bool, error) {
	ChatId = config.AppConfig.TELEGRAM_CHAT_ID
	text = fmt.Sprintf("YesDok %s %s", text, "telah mati")

	var err error
	var response *http.Response

	// Send the message
	url := fmt.Sprintf("%s/sendMessage", getUrl())
	body, _ := json.Marshal(map[string]string{
		"chat_id": ChatId,
		"text":    text,
	})

	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)

	helper.PanicIfError(err)

	// Close the request at the end
	defer response.Body.Close()

	// Body
	body, err = ioutil.ReadAll(response.Body)

	helper.PanicIfError(err)

	// Log
	log.Infof("Message '%s' was sent", text)
	log.Infof("Response JSON: %s", string(body))

	// Return
	return true, nil
}
