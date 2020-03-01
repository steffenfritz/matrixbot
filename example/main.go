package main

import (
	"github.com/steffenfritz/matrixbot"
	"log"
	"net/http"
)

func main() {

	c := matrixbot.ReadConfig("config.toml")
	resp := matrixbot.GetTokenJSON(c.Server, c.User, c.Password)

	retCode := matrixbot.SendTXTToRoom(c, resp.Access_token, "hello world")

	log.Println(http.StatusText(retCode))

}
