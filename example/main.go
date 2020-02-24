package main

import (
	"fmt"
	"github.com/steffenfritz/matrixbot"
)

func main() {

	c := matrixbot.ReadConfig("config.toml")
	resp := matrixbot.GetTokenJSON(c.Server, c.User, c.Password)

	fmt.Println(resp.Access_token)

}
