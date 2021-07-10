# umeng-openapi-sdk
`umeng-openapi-sdk` 是友盟+ Open API的golang版本

## 功能实现
- [x] <a href="https://developer.umeng.com/open-api/ns/com.umeng.uapp/apply">U-App-移动统计</a>
- [ ] <a href="https://developer.umeng.com/open-api/ns/com.umeng.apptrack/apply">AppTrack-移动广告监测</a>
- [ ] <a href="https://developer.umeng.com/open-api/ns/com.umeng.umini/apply">U-MiniProgram-小程序统计</a>
- [ ] <a href="https://developer.umeng.com/open-api/ns/com.umeng.uapm/apply">U-APM-应用性能监控</a>


## 如何使用
```go
var (
	apiKey      = "your api key"
	apiSecurity = "your api security"
	appKey      = "your app key"
)

func main() {
	uAppCli := uapp.NewUAppClient(apiKey, apiSecurity)
	
	// 获取当前用户所有App昨日和今日的基础统计数据
	allAppDataResp, err := uAppCli.GetAllAppData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n",allAppDataResp)
	
	// 获取新增账号
	newAccountsResp, err := uAppCli.GetNewAccounts(appKey, "2020-05-01", "2020-05-11", "daily", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n",newAccountsResp)
}
```