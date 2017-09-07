package weixin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//各种API请求的基础URL
const BaseURI = "https://api.weixin.qq.com/cgi-bin"

type BaseResponser interface {
	CleanError()
	FetchError() error
	FetchErrorCode() int
}

//响应的基础结构，用于判断全局性的错误。所有的API响应都会包含该结构
type BaseResponse struct {
	ErrCode int
	ErrMsg  string
}

func (resp *BaseResponse) FetchErrorCode() int {
	return resp.ErrCode
}

func (resp *BaseResponse) CleanError() {
	resp.ErrCode = 0
	resp.ErrMsg = ""
}

func (resp *BaseResponse) FetchError() error {
	if resp.ErrCode == 0 {
		return nil
	}

	return fmt.Errorf("[%d]%s", resp.ErrCode, resp.ErrMsg)
}

//执行API请求，会自动添加AccessToken
func (client *Client) request(method, path string, getParams url.Values, bodyBytes []byte, respInfo BaseResponser) error {
	//先带缓存请求
	err := client.requestWithTokenCache(method, path, getParams, bodyBytes, respInfo)

	if respInfo.FetchErrorCode() != 40001 {
		return err
	}

	fmt.Println(respInfo.FetchError().Error(), "Retry")

	//如果返回AccessToken相关的错误，则清缓存后重试一次
	client.cleanAccessTokenCache()
	respInfo.CleanError()
	return client.requestWithTokenCache(method, path, getParams, bodyBytes, respInfo)
}

func (client *Client) requestWithTokenCache(method, path string, getParams url.Values, bodyBytes []byte, respInfo BaseResponser) error {
	if getParams == nil {
		getParams = url.Values{}
	}

	accessToken, err := client.getAccessToken()
	if err != nil {
		return fmt.Errorf("获取AccessToken错误: %s", err)
	}
	getParams.Set("access_token", accessToken)

	uri := BaseURI + path + "?" + getParams.Encode()
	req, err := http.NewRequest(method, uri, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return fmt.Errorf("构造请求错误: %s", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("执行请求错误: %s", err)
	}

	defer resp.Body.Close()

	if respInfo == nil {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("读取响应错误: %s", err)
		}

		fmt.Println(string(b))
		return nil
	}

	err = json.NewDecoder(resp.Body).Decode(respInfo)
	if err != nil {
		return fmt.Errorf("读取响应错误: %s", err)
	}

	return respInfo.FetchError()
}
