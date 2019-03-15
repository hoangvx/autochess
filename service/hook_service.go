package service

import (
	"log"
	"net/http"
	"fmt"
	"bytes"
)

// Const
const PAGE_ACCESS_TOKEN = "EAAf2quZAUyVsBANVaKkI7gZAHxiLTOZCCKHSJ8OLhorkVWUUpB5VFx6BuFJ2Rm1tgx4qyKjiLOY75TIIPC9pYAcZBErozvb8xJcjKBanyYcMZCGbqERvd0bpaBJOdnXGCDIMYZCHZBtGSk9VtB78GBWOqwGdYG97RDY7oMLsrjY1AZDZD"

type HookService struct{}

func (h HookService) CallBack(sender_psid string, response string) {
	log.Println("callback", response)
	url := "https://graph.facebook.com/v2.6/me/messages"
	post_url := fmt.Sprintf("%s?access_token=%s", url, PAGE_ACCESS_TOKEN)

	post_data := fmt.Sprintf(`{"recipient":{"id": "%s"}, "message": { "text": "User said: %s" }}`, sender_psid, response)
	var jsonStr = []byte(post_data)
	req, err := http.NewRequest("POST", post_url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("access_token", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			panic(err)
	}
	defer resp.Body.Close()
}