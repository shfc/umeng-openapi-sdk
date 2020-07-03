package sdk

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"

	sdkErr "github.com/shfc/umeng-openapi-sdk/sdk/errors"
	"github.com/shfc/umeng-openapi-sdk/sdk/utils"
)

func CreateUmengClient(apiKey, apiSecurity string) *UmengClient {
	return &UmengClient{
		defaultServer: "gateway.open.umeng.com",
		apiKey:        apiKey,
		apiSecurity:   apiSecurity,
	}
}

type UmengAPI interface {
	GetApiUri() string
	GetRequiredParams() []string
	NeedSign() bool
	IsInnerApi() bool
}

type UmengClient struct {
	defaultServer string
	apiKey        string
	apiSecurity   string
	method        UmengAPI
}

func (cli *UmengClient) SetDefaultServer(newDefaultServer string) {
	cli.defaultServer = newDefaultServer
}

func (cli *UmengClient) buildSignUrlPath() string {
	return fmt.Sprintf("%s/%s/%s", P_PARAM2, cli.method.GetApiUri(), cli.apiKey)
}

func (cli *UmengClient) buildUrl(signedUrlPath string) string {
	apiType := P_OPENAPI
	if cli.method.IsInnerApi() {
		apiType = P_API
	}
	return fmt.Sprintf("https://%s/%s/%s", cli.defaultServer, apiType, signedUrlPath)
}

func (cli *UmengClient) checkRequiredParams(params url.Values) error {
	paramKeys := []string{}
	for k, _ := range params {
		paramKeys = append(paramKeys, k)
	}
	apiRequiredParams := cli.method.GetRequiredParams()
	if !utils.StrArrContainsAll(apiRequiredParams, paramKeys) {
		return errors.New(fmt.Sprintf("Required params missing: %v", apiRequiredParams))
	}
	return nil
}

func (cli *UmengClient) sign(urlPath string, params url.Values, secret string) (signature string, err error) {
	var paramList = []string{}
	if urlPath == "" {
		return "", errors.New("sign error: urlPath missing")
	}
	if secret == "" {
		return "", errors.New("sign error: secret missing")
	}
	for k, vals := range params {
		for _, v := range vals {
			paramList = append(paramList, fmt.Sprintf("%v%v", k, v))
		}
	}
	sort.Strings(paramList)
	for _, i := range paramList {
		urlPath = urlPath + i
	}
	return strings.ToUpper(hex.EncodeToString(utils.ShaHmac1(urlPath, secret))), nil
}

func (cli *UmengClient) Do(method UmengAPI, params url.Values) (respDeco *json.Decoder, err error) {
	cli.method = method
	signedUrl := cli.buildSignUrlPath()
	apiUrl := cli.buildUrl(signedUrl)
	formData := url.Values{}
	if len(cli.method.GetRequiredParams()) > 0 {
		err := cli.checkRequiredParams(params)
		if err != nil {
			return nil, err
		}
	}
	for k, vals := range params {
		for _, v := range vals {
			formData.Add(k, v)
		}
	}
	if method.NeedSign() {
		sign, err := cli.sign(signedUrl, formData, cli.apiSecurity)
		if err != nil {
			return nil, err
		}
		formData.Set(P_SIGN, sign)
	}
	hc := http.Client{}
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Submit the request
	res, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	if !(res.StatusCode >= 200 && res.StatusCode < 400) {
		errMsg := sdkErr.Err{}
		if err:=json.NewDecoder(res.Body).Decode(&errMsg);err!=nil{
			return nil,err
		}
		return nil, errors.New(fmt.Sprintf("request failed, statusCode %v, %v", res.StatusCode,errMsg))
	}
	return json.NewDecoder(res.Body), nil
}
