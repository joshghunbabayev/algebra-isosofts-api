package modules

import (
	"algebra-isosofts-api/middlewares"
	"encoding/json"
	"net/http"
	"os"
)

func GetAccountById(Id string) middlewares.RemoteAccount {
	isosoftsUrl := os.Getenv("ISOSOFTS_API_URL") + "/api/algebra/account/" + Id
	resp, err := http.Get(isosoftsUrl)

	var account middlewares.RemoteAccount

	if err != nil {
		return account
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return account
	}

	if err := json.NewDecoder(resp.Body).Decode(&account); err != nil {
		return account
	}
	return account
}
