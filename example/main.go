package main

import (
	"fmt"
	"github.com/steffenfritz/matrixbot"
)

func main() {

	server := "https://example.com"
	user := ""
	password := ""

	resp := matrixbot.GetTokenJSON(server, user, password)

	fmt.Println(resp.Access_token)

}
