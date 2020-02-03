package matrixbot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type AuthResponse struct {
	Access_token string `json:"access_token"`
	Device_id    string `json:"device_id"`
	Home_server  string `json:"home_server"`
	User_id      string `json:"user_id"`
}

// GetTokenJSON gets a session json object
func GetTokenJSON(server string, username string, password string) AuthResponse {

	authcontext := "/_matrix/client/r0/login"

	requestBody, err := json.Marshal(map[string]string{
		"type":     "m.login.password",
		"user":     username,
		"password": password,
	})

	e(err)
	// send post request to server
	resp, err := http.Post(server+authcontext, "application/json", bytes.NewBuffer(requestBody))
	e(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	e(err)

	var authResp AuthResponse
	err = json.Unmarshal(body, &authResp)
	e(err)

	return authResp
}
