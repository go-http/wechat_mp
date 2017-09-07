package weixin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type AccessTokenInfo struct {
	AccessToken string
	ExpireAt    time.Time
}

//从文件中加载AccessToken的缓存
func (client *Client) LoadAccessTokenFileCache(file string) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &client.accessTokenInfo)
}

//保存AccessToken的缓存到文件
func (client *Client) SaveAccessTokenFileCache(file string) error {
	b, err := json.Marshal(client.accessTokenInfo)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, b, 0700)
}

//获取AccessToken，如果缓存未过期则优先从缓存获取
func (client *Client) cleanAccessTokenCache() {
	client.accessTokenInfo.AccessToken = ""
}

//获取AccessToken，如果缓存未过期则优先从缓存获取
func (client *Client) getAccessToken() (string, error) {
	//先检查缓存是否有效
	if client.accessTokenInfo.AccessToken != "" && time.Now().Before(client.accessTokenInfo.ExpireAt) {
		return client.accessTokenInfo.AccessToken, nil
	}

	getParams := url.Values{}
	getParams.Set("appid", client.AppId)
	getParams.Set("secret", client.AppSecret)
	getParams.Set("grant_type", "client_credential")

	uri := BaseURI + "/token" + "?" + getParams.Encode()
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", fmt.Errorf("构造请求错误: %s", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("执行请求错误: %s", err)
	}

	defer resp.Body.Close()

	var respInfo struct {
		BaseResponse

		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}

	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	if err != nil {
		return "", fmt.Errorf("读取响应错误: %s", err)
	}

	client.accessTokenInfo.AccessToken = respInfo.AccessToken
	client.accessTokenInfo.ExpireAt = time.Now().Add(time.Second * time.Duration(respInfo.ExpiresIn))

	return respInfo.AccessToken, nil
}
