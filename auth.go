package matrixbot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// AuthResponse is a struct that holds authentication information
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	DeviceID    string `json:"device_id"`
	HomeServer  string `json:"home_server"`
	UserID      string `json:"user_id"`
}

// GetTokenJSON gets a session json object
func GetTokenJSON(server string, username string, passwd string) AuthResponse {

	authcontext := "/_matrix/client/r0/login"

	requestBody, err := json.Marshal(map[string]string{
		"type":     "m.login.password",
		"user":     username,
		"password": passwd,
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
