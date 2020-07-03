package uapp

/*
U-App新建数据源

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.createApp/1
*/
type CreateApp struct{}

func (CreateApp) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.createApp"
}

func (CreateApp) GetRequiredParams() []string {
	return []string{"name", "type", "platform", "language", "firstLevel", "secondLevel"}
}

func (CreateApp) NeedSign() bool {
	return true
}

func (CreateApp) IsInnerApi() bool {
	return false
}

/*
获取新增账号（仅游戏类型app）

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getNewAccounts/1
*/
type GetNewAccounts struct{}

func (GetNewAccounts) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getNewAccounts"
}

func (GetNewAccounts) GetRequiredParams() []string {
	return []string{"appkey", "startDate", "endDate"}
}

func (GetNewAccounts) NeedSign() bool {
	return true
}

func (GetNewAccounts) IsInnerApi() bool {
	return false
}

/*
获取活跃账号（仅游戏类型app）

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getActiveAccounts/1
*/
type GetActiveAccounts struct{}

func (GetActiveAccounts) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getActiveAccounts"
}

func (GetActiveAccounts) GetRequiredParams() []string {
	return []string{"appkey", "startDate", "endDate"}
}

func (GetActiveAccounts) NeedSign() bool {
	return true
}

func (GetActiveAccounts) IsInnerApi() bool {
	return false
}

/*
创建自定义事件

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.event.create/1
*/
type EventCreate struct{}

func (EventCreate) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.event.create"
}

func (EventCreate) GetRequiredParams() []string {
	return []string{"appkey", "eventName", "eventDisplayName"}
}

func (EventCreate) NeedSign() bool {
	return true
}

func (EventCreate) IsInnerApi() bool {
	return false
}

/*
根据渠道或版本条件，获取指定App某个时间范围内的启动次数

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getLaunchesByChannelOrVersion/1
*/
type GetLaunchesByChannelOrVersion struct{}

func (GetLaunchesByChannelOrVersion) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getLaunchesByChannelOrVersion"
}

func (GetLaunchesByChannelOrVersion) GetRequiredParams() []string {
	return []string{"appkey", "startDate", "endDate", "periodType"}
}

func (GetLaunchesByChannelOrVersion) NeedSign() bool {
	return true
}

func (GetLaunchesByChannelOrVersion) IsInnerApi() bool {
	return false
}

/*
根据渠道或版本条件，获取指定App某个时间范围内的活跃用户数

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getActiveUsersByChannelOrVersion/1
*/
type GetActiveUsersByChannelOrVersion struct{}

func (GetActiveUsersByChannelOrVersion) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getActiveUsersByChannelOrVersion"
}

func (GetActiveUsersByChannelOrVersion) GetRequiredParams() []string {
	return []string{"appkey", "startDate", "endDate", "periodType"}
}

func (GetActiveUsersByChannelOrVersion) NeedSign() bool {
	return true
}

func (GetActiveUsersByChannelOrVersion) IsInnerApi() bool {
	return false
}

/*
根据渠道或版本条件，获取指定App某个时间范围内的新增用户数

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getNewUsersByChannelOrVersion/1
*/
type GetNewUsersByChannelOrVersion struct{}

func (GetNewUsersByChannelOrVersion) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getNewUsersByChannelOrVersion"
}

func (GetNewUsersByChannelOrVersion) GetRequiredParams() []string {
	return []string{"appkey", "startDate", "endDate", "periodType"}
}

func (GetNewUsersByChannelOrVersion) NeedSign() bool {
	return true
}

func (GetNewUsersByChannelOrVersion) IsInnerApi() bool {
	return false
}

/*
根据自定义事件参数值，获取使用时长

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.event.param.getValueDurationList/1
*/
type EventParamGetValueDurationList struct{}

func (EventParamGetValueDurationList) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.event.param.getValueDurationList"
}

func (EventParamGetValueDurationList) GetRequiredParams() []string {
	return []string{"appkey", "startDate", "endDate", "eventName", "eventParamName"}
}

func (EventParamGetValueDurationList) NeedSign() bool {
	return true
}

func (EventParamGetValueDurationList) IsInnerApi() bool {
	return false
}

/*
获取指定App今天与昨天的统计数据

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getTodayYesterdayData/1
*/
type GetTodayYesterdayData struct{}

func (GetTodayYesterdayData) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getTodayYesterdayData"
}

func (GetTodayYesterdayData) GetRequiredParams() []string {
	return []string{"appkey"}
}

func (GetTodayYesterdayData) NeedSign() bool {
	return true
}

func (GetTodayYesterdayData) IsInnerApi() bool {
	return false
}

/*
获取指定App昨日的统计数据

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getYesterdayData/1
*/
type GetYesterdayData struct{}

func (GetYesterdayData) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getYesterdayData"
}

func (GetYesterdayData) GetRequiredParams() []string {
	return []string{"appkey"}
}

func (GetYesterdayData) NeedSign() bool {
	return true
}

func (GetYesterdayData) IsInnerApi() bool {
	return false
}

/*
获取指定App今日的统计数据

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getTodayData/1
*/
type GetTodayData struct{}

func (GetTodayData) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getTodayData"
}

func (GetTodayData) GetRequiredParams() []string {
	return []string{"appkey"}
}

func (GetTodayData) NeedSign() bool {
	return true
}

func (GetTodayData) IsInnerApi() bool {
	return false
}

/*
获取自定义事件的独立用户数统计数据（按设备device统计 "data":[{"data":[123],"dates":["2018-01-01"]}）

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.event.getUniqueUsers/1
*/
type EventGetUniqueUsers struct{}

func (EventGetUniqueUsers) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.event.getUniqueUsers"
}

func (EventGetUniqueUsers) GetRequiredParams() []string {
	return []string{"appkey", "startDate", "endDate", "eventName"}
}

func (EventGetUniqueUsers) NeedSign() bool {
	return true
}

func (EventGetUniqueUsers) IsInnerApi() bool {
	return false
}

/*
获取当前用户所有App昨日和今日的基础统计数据（活跃用户数，新增用户数，启动次数，总用户数）

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getAllAppData/1
*/
type GetAllAppData struct{}

func (GetAllAppData) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getAllAppData"
}

func (GetAllAppData) GetRequiredParams() []string {
	return []string{}
}

func (GetAllAppData) NeedSign() bool {
	return true
}

func (GetAllAppData) IsInnerApi() bool {
	return false
}

/*
获取当前用户所有App的数量

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getAppCount/1
*/
type GetAppCount struct{}

func (GetAppCount) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getAppCount"
}

func (GetAppCount) GetRequiredParams() []string {
	return []string{}
}

func (GetAppCount) NeedSign() bool {
	return true
}

func (GetAppCount) IsInnerApi() bool {
	return false
}

/*
获取指定App按照分发渠道维度的统计数据

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getChannelData/1
*/
type GetChannelData struct{}

func (GetChannelData) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getChannelData"
}

func (GetChannelData) GetRequiredParams() []string {
	return []string{"appkey", "date"}
}

func (GetChannelData) NeedSign() bool {
	return true
}

func (GetChannelData) IsInnerApi() bool {
	return false
}

/*
获取指定App按照版本维度的统计数据

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getVersionData/1
*/
type GetVersionData struct{}

func (GetVersionData) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getVersionData"
}

func (GetVersionData) GetRequiredParams() []string {
	return []string{"appkey", "date"}
}

func (GetVersionData) NeedSign() bool {
	return true
}

func (GetVersionData) IsInnerApi() bool {
	return false
}

/*
获取自定义事件某个参数按照值维度的统计数据

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.event.param.getData/1
*/
type EventParamGetData struct{}

func (EventParamGetData) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.event.param.getData"
}

func (EventParamGetData) GetRequiredParams() []string {
	return []string{"appkey", "startDate", "endDate", "eventName", "eventParamName", "paramValueName"}
}

func (EventParamGetData) NeedSign() bool {
	return true
}

func (EventParamGetData) IsInnerApi() bool {
	return false
}

/*
获取自定义事件某个参数的取值范围列表

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.event.param.getValueList/1
*/
type EventParamGetValueList struct{}

func (EventParamGetValueList) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.event.param.getValueList"
}

func (EventParamGetValueList) GetRequiredParams() []string {
	return []string{"appkey", "startDate", "endDate", "eventName", "eventParamName"}
}

func (EventParamGetValueList) NeedSign() bool {
	return true
}

func (EventParamGetValueList) IsInnerApi() bool {
	return false
}

/*
获取自定义事件事件的统计数据

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.event.getData/1
*/
type EventGetData struct{}

func (EventGetData) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.event.getData"
}

func (EventGetData) GetRequiredParams() []string {
	return []string{"appkey", "startDate", "endDate", "eventName"}
}

func (EventGetData) NeedSign() bool {
	return true
}

func (EventGetData) IsInnerApi() bool {
	return false
}

/*
获取自定义事件的参数列表

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.event.param.list/1
*/
type EventParamList struct{}

func (EventParamList) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.event.param.list"
}

func (EventParamList) GetRequiredParams() []string {
	return []string{"startDate", "endDate", "eventId", "appkey"}
}

func (EventParamList) NeedSign() bool {
	return true
}

func (EventParamList) IsInnerApi() bool {
	return false
}

/*
获取自定义事件列表

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.event.list/1
*/
type EventList struct{}

func (EventList) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.event.list"
}

func (EventList) GetRequiredParams() []string {
	return []string{"appkey", "startDate", "endDate"}
}

func (EventList) NeedSign() bool {
	return true
}

func (EventList) IsInnerApi() bool {
	return false
}

/*
获取指定App某个时间范围内的新增用户留存率

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getRetentions/1
*/
type GetRetentions struct{}

func (GetRetentions) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getRetentions"
}

func (GetRetentions) GetRequiredParams() []string {
	return []string{"appkey", "startDate", "endDate"}
}

func (GetRetentions) NeedSign() bool {
	return true
}

func (GetRetentions) IsInnerApi() bool {
	return false
}

/*
获取指定App某个时间范围内的使用时长统计数据

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getDurations/1
*/
type GetDurations struct{}

func (GetDurations) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getDurations"
}

func (GetDurations) GetRequiredParams() []string {
	return []string{"appkey", "date"}
}

func (GetDurations) NeedSign() bool {
	return true
}

func (GetDurations) IsInnerApi() bool {
	return false
}

/*
获取指定App某个时间范围内的启动次数

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getLaunches/1
*/
type GetLaunches struct{}

func (GetLaunches) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getLaunches"
}

func (GetLaunches) GetRequiredParams() []string {
	return []string{"appkey", "startDate", "endDate"}
}

func (GetLaunches) NeedSign() bool {
	return true
}

func (GetLaunches) IsInnerApi() bool {
	return false
}

/*
获取指定App某个时间范围内的活跃用户数

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getActiveUsers/1
*/
type GetActiveUsers struct{}

func (GetActiveUsers) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getActiveUsers"
}

func (GetActiveUsers) GetRequiredParams() []string {
	return []string{"appkey", "startDate", "endDate"}
}

func (GetActiveUsers) NeedSign() bool {
	return true
}

func (GetActiveUsers) IsInnerApi() bool {
	return false
}

/*
获取指定App某个时间范围内的新增用户数

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getNewUsers/1
*/
type GetNewUsers struct{}

func (GetNewUsers) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getNewUsers"
}

func (GetNewUsers) GetRequiredParams() []string {
	return []string{"appkey", "startDate", "endDate"}
}

func (GetNewUsers) NeedSign() bool {
	return true
}

func (GetNewUsers) IsInnerApi() bool {
	return false
}

/*
获取指定App特定日期的统计数据

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getDailyData/1
*/
type GetDailyData struct{}

func (GetDailyData) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getDailyData"
}

func (GetDailyData) GetRequiredParams() []string {
	return []string{"appkey", "date"}
}

func (GetDailyData) NeedSign() bool {
	return true
}

func (GetDailyData) IsInnerApi() bool {
	return false
}

/*
获取当前用户的所有App列表

    References
    ----------
    https://developer.umeng.com/open-api/docs/com.umeng.uapp/umeng.uapp.getAppList/1
*/
type GetAppList struct{}

func (GetAppList) GetApiUri() string {
	return "1/com.umeng.uapp/umeng.uapp.getAppList"
}

func (GetAppList) GetRequiredParams() []string {
	return []string{}
}

func (GetAppList) NeedSign() bool {
	return true
}

func (GetAppList) IsInnerApi() bool {
	return false
}
