package main

import (
	"log"
	"net/http"

	"github.com/steffenfritz/matrixbot"
)

func main() {

	c := matrixbot.ReadConfig("config.toml")
	resp := matrixbot.GetTokenJSON(c.Server, c.User, c.Password)

	retCode := matrixbot.SendTXTToRoom(c, resp.AccessToken, "hello world")

	log.Println(http.StatusText(retCode))

}
