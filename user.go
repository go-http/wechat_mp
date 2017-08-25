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

type UserInfo struct {
	OpenId     string
	Nickname   string //昵称
	Sex        int    //性别：0未知、1男性、2女性
	Language   string
	City       string
	Province   string
	Country    string
	HeadImgUrl string //用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	Unionid    string
	Remark     string //公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	GroupId    int    //用户所在的分组ID（兼容旧的用户分组接口）
	TagidList  []int  `json:"tagid_list"` //用户被打上的标签ID列表

	Subscribe     int //是否已关注，为0时表示未关注，其他信息无法获取
	SubscribeTime int `json:"subscribe_time"` //用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
}

//获取指定用户的信息
func (client *Client) UserInfo(openId string) (*UserInfo, error) {
	getParams := url.Values{}
	getParams.Set("openid", openId)

	var resp struct {
		BaseResponse
		*UserInfo
	}
	err := client.request("GET", "/user/info", getParams, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("获取用户OpenID列表出错: %s", err)
	}

	return resp.UserInfo, nil
}
