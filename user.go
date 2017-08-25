package weixin

import (
	"fmt"
	"net/url"
)

type UserGetResponse struct {
	BaseResponse

	Total int //总用户数
	Count int //当前返回用户数，最大10000
	Data  struct {
		OpenId []string
	}
	NextOpenId string `json:"next_openid"` //为空时表示已经获取完成
}

//获取关注的用户openid列表
func (client *Client) UserList() ([]string, error) {
	var openIds []string
	var nextOpenId string

	for {
		getParams := url.Values{}
		if nextOpenId != "" {
			getParams.Set("next_openid", nextOpenId)
		}

		var resp UserGetResponse
		err := client.request("GET", "/user/get", getParams, nil, &resp)
		if err != nil {
			return nil, fmt.Errorf("获取用户OpenID列表出错: %s", err)
		}

		openIds = append(openIds, resp.Data.OpenId...)

		//如果不存在next_openid说明已获取完
		if resp.NextOpenId == "" {
			break
		}

		nextOpenId = resp.NextOpenId
	}

	return openIds, nil
}
