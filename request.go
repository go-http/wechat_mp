package weixin

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

//各种API请求的基础URL
const BaseURI = "https://api.weixin.qq.com/cgi-bin"

type BaseResponser interface {
	FetchError() error
}

//响应的基础结构，用于判断全局性的错误。所有的API响应都会包含该结构
type BaseResponse struct {
	ErrCode int
	ErrMsg  string
}

func (resp *BaseResponse) FetchError() error {
	if resp.ErrCode == 0 {
		return nil
	}

	return fmt.Errorf("[%d]%s", resp.ErrCode, resp.ErrMsg)
}

//执行API请求，会自动添加AccessToken
func (client *Client) request(method, path string, getParams url.Values, body io.Reader, respInfo BaseResponser) error {
	if getParams == nil {
		getParams = url.Values{}
	}

	accessToken, err := client.getAccessToken()
	if err != nil {
		return fmt.Errorf("获取AccessToken错误: %s", err)
	}
	getParams.Set("access_token", accessToken)

	uri := BaseURI + path + "?" + getParams.Encode()
	req, err := http.NewRequest(method, uri, body)
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
