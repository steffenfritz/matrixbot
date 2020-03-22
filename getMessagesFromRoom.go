package matrixbot

import (
	"net/http"
)

// GetInitialEvent gets the first event id
func GetInitialEvent(c Config, token string, txtmsg string) int {

	url := c.Server + "/_matrix/client/r0/rooms/" + c.RoomID + "/send/m.room.message?access_token=" + token

}
