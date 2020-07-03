package uapp

import (
	"github.com/shfc/umeng-openapi-sdk/sdk/uapp"
	"net/url"
	"strconv"

	"github.com/shfc/umeng-openapi-sdk/sdk"
)

func NewUAppClient(apiKey, apiSecurity string) *UAppClient {
	uCli := sdk.CreateUmengClient(apiKey, apiSecurity)
	return &UAppClient{uCli}
}

type UAppClient struct {
	*sdk.UmengClient
}

func (cli *UAppClient) GetNewAccounts(appKey, startDate, endDate, periodType, channel string) (*uapp.GetNewAccountsResp, error) {
	resp := &uapp.GetNewAccountsResp{}
	formVals := url.Values{
		"appkey":     {appKey},
		"startDate":  {startDate},
		"endDate":    {endDate},
		"periodType": {periodType},
		"channel":    {channel},
	}
	respDeco, err := cli.Do(uapp.GetNewAccounts{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetActiveAccounts(appKey, startDate, endDate, periodType, channel string) (*uapp.GetActiveAccountsResp, error) {
	resp := &uapp.GetActiveAccountsResp{}
	formVals := url.Values{
		"appkey":     {appKey},
		"startDate":  {startDate},
		"endDate":    {endDate},
		"periodType": {periodType},
		"channel":    {channel},
	}
	respDeco, err := cli.Do(uapp.GetActiveAccounts{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) EventCreate(appKey, eventName, eventDisplayName string, eventType bool) (*uapp.EventCreateResp, error) {
	resp := &uapp.EventCreateResp{}
	formVals := url.Values{
		"appkey":           {appKey},
		"eventName":        {eventName},
		"eventDisplayName": {eventDisplayName},
		"eventType":        {strconv.FormatBool(eventType)},
	}
	respDeco, err := cli.Do(uapp.EventCreate{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetLaunchesByChannelOrVersion(appKey, startDate, endDate, periodType, channels, versions string) (*uapp.GetLaunchesByChannelOrVersionResp, error) {
	resp := &uapp.GetLaunchesByChannelOrVersionResp{}
	formVals := url.Values{
		"appkey":     {appKey},
		"startDate":  {startDate},
		"endDate":    {endDate},
		"periodType": {periodType},
		"channels":   {channels},
		"versions":   {versions},
	}
	respDeco, err := cli.Do(uapp.GetLaunchesByChannelOrVersion{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetActiveUsersByChannelOrVersion(appKey, startDate, endDate, periodType, channels, versions string) (*uapp.GetActiveUsersByChannelOrVersionResp, error) {
	resp := &uapp.GetActiveUsersByChannelOrVersionResp{}
	formVals := url.Values{
		"appkey":     {appKey},
		"startDate":  {startDate},
		"endDate":    {endDate},
		"periodType": {periodType},
		"channels":   {channels},
		"versions":   {versions},
	}
	respDeco, err := cli.Do(uapp.GetActiveUsersByChannelOrVersion{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetNewUsersByChannelOrVersion(appKey, startDate, endDate, periodType, channels, versions string) (*uapp.GetNewUsersByChannelOrVersionResp, error) {
	resp := &uapp.GetNewUsersByChannelOrVersionResp{}
	formVals := url.Values{
		"appkey":     {appKey},
		"startDate":  {startDate},
		"endDate":    {endDate},
		"periodType": {periodType},
		"channels":   {channels},
		"versions":   {versions},
	}
	respDeco, err := cli.Do(uapp.GetNewUsersByChannelOrVersion{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) EventParamGetValueDurationList(appKey, startDate, endDate, eventName, eventParamName string) (*uapp.EventParamGetValueDurationListResp, error) {
	resp := &uapp.EventParamGetValueDurationListResp{}
	formVals := url.Values{
		"appkey":         {appKey},
		"startDate":      {startDate},
		"endDate":        {endDate},
		"eventName":      {eventName},
		"eventParamName": {eventParamName},
	}
	respDeco, err := cli.Do(uapp.EventParamGetValueDurationList{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetTodayYesterdayData(appKey string) (*uapp.GetTodayYesterdayDataResp, error) {
	resp := &uapp.GetTodayYesterdayDataResp{}
	formVals := url.Values{
		"appkey": {appKey},
	}
	respDeco, err := cli.Do(uapp.GetTodayYesterdayData{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetYesterdayData(appKey string) (*uapp.GetYesterdayDataResp, error) {
	resp := &uapp.GetYesterdayDataResp{}
	formVals := url.Values{
		"appkey": {appKey},
	}
	respDeco, err := cli.Do(uapp.GetYesterdayData{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetTodayData(appKey string) (*uapp.GetTodayDataResp, error) {
	resp := &uapp.GetTodayDataResp{}
	formVals := url.Values{
		"appkey": {appKey},
	}
	respDeco, err := cli.Do(uapp.GetTodayData{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) EventGetUniqueUsers(appKey, startDate, endDate, eventName string) (*uapp.EventGetUniqueUsersResp, error) {
	resp := &uapp.EventGetUniqueUsersResp{}
	formVals := url.Values{
		"appkey":    {appKey},
		"startDate": {startDate},
		"endDate":   {endDate},
		"eventName": {eventName},
	}
	respDeco, err := cli.Do(uapp.EventGetUniqueUsers{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetAllAppData() (*uapp.GetAllAppDataResp, error) {
	resp := &uapp.GetAllAppDataResp{}
	respDeco, err := cli.Do(uapp.GetAllAppData{}, nil)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetAppCount() (*uapp.GetAppCountResp, error) {
	resp := &uapp.GetAppCountResp{}
	formVals := url.Values{
	}
	respDeco, err := cli.Do(uapp.GetAppCount{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetChannelData(appKey, date string, perPage, page int64) (*uapp.GetChannelDataResp, error) {
	resp := &uapp.GetChannelDataResp{}
	formVals := url.Values{
		"appkey":  {appKey},
		"date":    {date},
		"perPage": {strconv.FormatInt(perPage, 10)},
		"page":    {strconv.FormatInt(page, 10)},
	}
	respDeco, err := cli.Do(uapp.GetChannelData{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetVersionData(appKey, date string) (*uapp.GetVersionDataResp, error) {
	resp := &uapp.GetVersionDataResp{}
	formVals := url.Values{
		"appkey": {appKey},
		"date":   {date},
	}
	respDeco, err := cli.Do(uapp.GetVersionData{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) EventParamGetData(appKey, startDate, endDate, eventName, eventParamName, paramValueName string) (*uapp.EventParamGetDataResp, error) {
	resp := &uapp.EventParamGetDataResp{}
	formVals := url.Values{
		"appkey":         {appKey},
		"startDate":      {startDate},
		"endDate":        {endDate},
		"eventName":      {eventName},
		"eventParamName": {eventParamName},
		"paramValueName": {paramValueName},
	}
	respDeco, err := cli.Do(uapp.EventParamGetData{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) EventParamGetValueList(appKey, startDate, endDate, eventName, eventParamName string) (*uapp.EventParamGetValueListResp, error) {
	resp := &uapp.EventParamGetValueListResp{}
	formVals := url.Values{
		"appkey":         {appKey},
		"startDate":      {startDate},
		"endDate":        {endDate},
		"eventName":      {eventName},
		"eventParamName": {eventParamName},
	}
	respDeco, err := cli.Do(uapp.EventParamGetValueList{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) EventGetData(appKey, startDate, endDate, eventName string) (*uapp.EventGetDataResp, error) {
	resp := &uapp.EventGetDataResp{}
	formVals := url.Values{
		"appkey":    {appKey},
		"startDate": {startDate},
		"endDate":   {endDate},
		"eventName": {eventName},
	}
	respDeco, err := cli.Do(uapp.EventGetData{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) EventParamList(startDate, endDate, eventId, appKey string) (*uapp.EventParamListResp, error) {
	resp := &uapp.EventParamListResp{}
	formVals := url.Values{
		"startDate": {startDate},
		"endDate":   {endDate},
		"eventId":   {eventId},
		"appkey":    {appKey},
	}
	respDeco, err := cli.Do(uapp.EventParamList{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) EventList(appKey, startDate, endDate, version string, perPage, page int64) (*uapp.EventListResp, error) {
	resp := &uapp.EventListResp{}
	formVals := url.Values{
		"appkey":    {appKey},
		"startDate": {startDate},
		"endDate":   {endDate},
		"perPage":   {strconv.FormatInt(perPage, 10)},
		"page":      {strconv.FormatInt(page, 10)},
		"version":   {version},
	}
	respDeco, err := cli.Do(uapp.EventList{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetRetentions(appKey, startDate, endDate, periodType, channel, version, typ string) (*uapp.GetRetentionsResp, error) {
	resp := &uapp.GetRetentionsResp{}
	formVals := url.Values{
		"appkey":     {appKey},
		"startDate":  {startDate},
		"endDate":    {endDate},
		"periodType": {periodType},
		"channel":    {channel},
		"version":    {version},
		"type":       {typ},
	}
	respDeco, err := cli.Do(uapp.GetRetentions{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetDurations(appKey, date, statType, channel, version string) (*uapp.GetDurationsResp, error) {
	resp := &uapp.GetDurationsResp{}
	formVals := url.Values{
		"appkey":   {appKey},
		"date":     {date},
		"statType": {statType},
		"channel":  {channel},
		"version":  {version},
	}
	respDeco, err := cli.Do(uapp.GetDurations{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetLaunches(appKey, startDate, endDate, periodType string) (*uapp.GetLaunchesResp, error) {
	resp := &uapp.GetLaunchesResp{}
	formVals := url.Values{
		"appkey":     {appKey},
		"startDate":  {startDate},
		"endDate":    {endDate},
		"periodType": {periodType},
	}
	respDeco, err := cli.Do(uapp.GetLaunches{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetActiveUsers(appKey, startDate, endDate, periodType string) (*uapp.GetActiveUsersResp, error) {
	resp := &uapp.GetActiveUsersResp{}
	formVals := url.Values{
		"appkey":     {appKey},
		"startDate":  {startDate},
		"endDate":    {endDate},
		"periodType": {periodType},
	}
	respDeco, err := cli.Do(uapp.GetActiveUsers{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetNewUsers(appKey, startDate, endDate, periodType string) (*uapp.GetNewUsersResp, error) {
	resp := &uapp.GetNewUsersResp{}
	formVals := url.Values{
		"appkey":     {appKey},
		"startDate":  {startDate},
		"endDate":    {endDate},
		"periodType": {periodType},
	}
	respDeco, err := cli.Do(uapp.GetNewUsers{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetDailyData(appKey, date, version, channel string) (*uapp.GetDailyDataResp, error) {
	resp := &uapp.GetDailyDataResp{}
	formVals := url.Values{
		"appkey":  {appKey},
		"date":    {date},
		"version": {version},
		"channel": {channel},
	}
	respDeco, err := cli.Do(uapp.GetDailyData{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *UAppClient) GetAppList(page, perPage, accessToken string) (*uapp.GetAppListResp, error) {
	resp := &uapp.GetAppListResp{}
	formVals := url.Values{
		"page":        {page},
		"perPage":     {perPage},
		"accessToken": {accessToken},
	}
	respDeco, err := cli.Do(uapp.GetAppList{}, formVals)
	if err != nil {
		return nil, err
	}
	if err := respDeco.Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}
