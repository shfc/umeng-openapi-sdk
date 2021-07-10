package uapp

type GetAllAppDataResp struct {
	AllAppData []struct {
		YesterdayNewUsers        int64 `json:"yesterdayNewUsers"`
		YesterdayUniqNewUsers    int64 `json:"yesterdayUniqNewUsers"`
		TodayLaunches            int64 `json:"todayLaunches"`
		TotalUsers               int64 `json:"totalUsers"`
		TodayNewUsers            int64 `json:"todayNewUsers"`
		YesterdayUniqActiveUsers int64 `json:"yesterdayUniqActiveUsers"`
		TodayActivityUsers       int64 `json:"todayActivityUsers"`
		YesterdayLaunches        int64 `json:"yesterdayLaunches"`
		YesterdayActivityUsers   int64 `json:"yesterdayActivityUsers"`
	} `json:"allAppData"`
}

type GetNewAccountsResp struct {
	NewAccountInfo []struct {
		Date           string `json:"date"`
		HourNewUser    string `json:"hourNewUser"`
		NewUser        int64  `json:"newUser"`
		NewAccount     int64  `json:"newAccount"`
		HourNewAccount string `json:"hourNewAccount"`
	} `json:"newAccountInfo"`
}

type GetActiveAccountsResp []struct {
	ActiveAccountInfo []struct {
		Date          string `json:"date"`
		ActiveAccount int64  `json:"activeAccount"`
		ActiveUser    int64  `json:"activeUser"`
	} `json:"activeAccountInfo"`
}

type EventCreateResp struct {
	Msg    string `json:"msg"`
	Status int64  `json:"status"`
}

type GetLaunchesByChannelOrVersionResp struct {
	LaunchInfo []struct {
		Date       string `json:"date"`
		DailyValue []struct {
			Name  string `json:"name"`
			Value int64  `json:"value"`
		} `json:"dailyValue"`
		HourValue string `json:"hourValue"`
		Value     int64  `json:"value"`
	} `json:"launchInfo"`
}

type GetActiveUsersByChannelOrVersionResp struct {
	ActiveUserInfo []struct {
		Date       string `json:"date"`
		DailyValue []struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"dailyValue"`
		HourValue string `json:"hourValue"`
		Value     int    `json:"value"`
	} `json:"activeUserInfo"`
}

type GetNewUsersByChannelOrVersionResp struct {
	NewUserInfo []struct {
		Date       string `json:"date"`
		DailyValue []struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"dailyValue"`
		HourValue string `json:"hourValue"`
		Value     int    `json:"value"`
	} `json:"newUserInfo"`
}

type EventParamGetValueDurationListResp struct {
	ParamInfos []struct {
		Name    string `json:"name"`
		Count   int    `json:"count"`
		Percent string `json:"percent"`
	} `json:"paramInfos"`
}

type GetTodayYesterdayDataResp struct {
	YesterdayData struct {
		Date          string `json:"date"`
		NewUsers      int    `json:"newUsers"`
		TotalUsers    int    `json:"totalUsers"`
		ActivityUsers int    `json:"activityUsers"`
		Launches      int    `json:"launches"`
		PayUsers      int    `json:"payUsers"`
	} `json:"yesterdayData"`
	TodayData struct {
		Date          string `json:"date"`
		NewUsers      int    `json:"newUsers"`
		TotalUsers    int    `json:"totalUsers"`
		ActivityUsers int    `json:"activityUsers"`
		Launches      int    `json:"launches"`
		PayUsers      int    `json:"payUsers"`
	} `json:"todayData"`
}

type GetYesterdayDataResp struct {
	YesterdayData struct {
		Date          string `json:"date"`
		NewUsers      int    `json:"newUsers"`
		TotalUsers    int    `json:"totalUsers"`
		ActivityUsers int    `json:"activityUsers"`
		Launches      int    `json:"launches"`
		PayUsers      int    `json:"payUsers"`
	} `json:"yesterdayData"`
}

type GetTodayDataResp struct {
	TodayData struct {
		Date          string `json:"date"`
		NewUsers      int    `json:"newUsers"`
		TotalUsers    int    `json:"totalUsers"`
		ActivityUsers int    `json:"activityUsers"`
		Launches      int    `json:"launches"`
		PayUsers      int    `json:"payUsers"`
	} `json:"todayData"`
}

type EventGetUniqueUsersResp struct {
	UniqueUsers []struct {
		Data  string `json:"data"`
		Dates string `json:"dates"`
	} `json:"uniqueUsers"`
}

type GetAppCountResp struct {
	Count int64 `json:"count"`
}

type GetChannelDataResp struct {
	ChannelInfos []struct {
		Duration      string  `json:"duration"`
		Date          string  `json:"date"`
		ActiveUser    int     `json:"activeUser"`
		NewUser       int     `json:"newUser"`
		TotalUser     int     `json:"totalUser"`
		Channel       string  `json:"channel"`
		Launch        int     `json:"launch"`
		ID            string  `json:"id"`
		TotalUserRate float64 `json:"totalUserRate"`
	} `json:"channelInfos"`
}

type GetVersionDataResp struct {
	VersionInfos []struct {
		Date          string  `json:"date"`
		ActiveUser    int     `json:"activeUser"`
		NewUser       int     `json:"newUser"`
		TotalUser     int     `json:"totalUser"`
		Version       string  `json:"version"`
		TotalUserRate float64 `json:"totalUserRate"`
	} `json:"versionInfos"`
}

type EventParamGetDataResp struct {
	ParamValueData []struct {
		Data  string `json:"data"`
		Dates string `json:"dates"`
	} `json:"paramValueData"`
}

type EventParamGetValueListResp struct {
	ParamInfos []struct {
		Name    string `json:"name"`
		Count   int    `json:"count"`
		Percent string `json:"percent"`
	} `json:"paramInfos"`
}

type EventGetDataResp struct {
	EventData []struct {
		Data  string `json:"data"`
		Dates string `json:"dates"`
	} `json:"eventData"`
}

type EventParamListResp struct {
	ParamInfos []struct {
		DisplayName string `json:"displayName"`
		Name        string `json:"name"`
		ParamID     string `json:"paramId"`
	} `json:"paramInfos"`
}

type EventListResp struct {
	EventInfo []struct {
		DisplayName string `json:"displayName"`
		Name        string `json:"name"`
		Count       int    `json:"count"`
		ID          string `json:"id"`
	} `json:"eventInfo"`
	TotalPage int `json:"totalPage"`
	Page      int `json:"page"`
}

type GetRetentionsResp struct {
	RetentionInfo []struct {
		Date             string    `json:"date"`
		TotalInstallUser int       `json:"totalInstallUser"`
		RetentionRate    []float64 `json:"retentionRate"`
	} `json:"retentionInfo"`
}

type GetDurationsResp struct {
	DurationInfos []struct {
		Name    string `json:"name"`
		Value   int    `json:"value"`
		Percent string `json:"percent"`
	} `json:"durationInfos"`
	Average float64 `json:"average"`
}

type GetLaunchesResp struct {
	LaunchInfo []struct {
		Date       string `json:"date"`
		DailyValue []struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"dailyValue"`
		HourValue string `json:"hourValue"`
		Value     int    `json:"value"`
	} `json:"launchInfo"`
}

type GetActiveUsersResp struct {
	ActiveUserInfo []struct {
		Date       string `json:"date"`
		DailyValue []struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"dailyValue"`
		HourValue string `json:"hourValue"`
		Value     int    `json:"value"`
	} `json:"activeUserInfo"`
}

type GetNewUsersResp struct {
	NewUserInfo []struct {
		Date       string `json:"date"`
		DailyValue []struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"dailyValue"`
		HourValue string `json:"hourValue"`
		Value     int    `json:"value"`
	} `json:"newUserInfo"`
}

type GetDailyDataResp struct {
	DailyData struct {
		Date          string `json:"date"`
		NewUsers      int    `json:"newUsers"`
		TotalUsers    int    `json:"totalUsers"`
		ActivityUsers int    `json:"activityUsers"`
		Launches      int    `json:"launches"`
		PayUsers      int    `json:"payUsers"`
	} `json:"dailyData"`
}

type GetAppListResp struct {
	AppInfos []struct {
		CreatedAt  string `json:"createdAt"`
		UseGameSdk string `json:"useGameSdk"`
		Name       string `json:"name"`
		Appkey     string `json:"appkey"`
		Category   string `json:"category"`
		Popular    int    `json:"popular"`
		Platform   string `json:"platform"`
		UpdatedAt  string `json:"updatedAt"`
	} `json:"appInfos"`
	TotalPage int `json:"totalPage"`
	Page      int `json:"page"`
}
