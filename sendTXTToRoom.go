package matrixbot

import (
	"bytes"
	"log"
	"net/http"
)

// SendTXTToRoom sends a text message
func SendTXTToRoom(c Config, token string, txtmsg string) int {

	url := c.Server + "/_matrix/client/r0/rooms/" + c.RoomID + "/send/m.room.message?access_token=" + token

	msg := "{\"msgtype\": \"m.text\", \"body\": \"" + txtmsg + "\"}"
	body := []byte(msg)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	return resp.StatusCode
}
