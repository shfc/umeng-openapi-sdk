package main

import (
	"encoding/json"
	"fmt"

	"github.com/shfc/umeng-openapi-sdk/services/uapp"
)

var (
	apiKey      = ""
	apiSecurity = ""
	appKey      = ""
)

func main() {
	uAppCli := uapp.NewUAppClient(apiKey, apiSecurity)
	allAppDataResp, err := uAppCli.GetAllAppData()
	if err != nil {
		fmt.Println(err)
		return
	}
	opt, _ := json.MarshalIndent(allAppDataResp, "", "    ")
	fmt.Println(string(opt))

	newAccountsResp, err := uAppCli.GetNewAccounts(appKey, "2020-05-01", "2020-05-11", "daily", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	opt, _ = json.MarshalIndent(newAccountsResp, "", "    ")
	fmt.Println(string(opt))

}
