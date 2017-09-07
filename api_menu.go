package weixin

import (
	"bytes"
	"encoding/json"
)

const (
	MenuButtonTypeView        = "view"
	MenuButtonTypeClick       = "click"
	MenuButtonTypeMiniProgram = "miniprogram"
)

type MenuButton struct {
	Type string `json:"type"`
	Name string `json:"name"`

	Key      string `json:"key"`
	Url      string `json:"url"`
	AppId    string `json:"appid"`
	MediaId  string `json:"media_id"`
	PagePath string `json:"pagepath"`

	SubButton []MenuButton `json:"sub_button,omitempty"`
}

func (client *Client) MenuCreate(buttons []MenuButton) error {
	var reqData struct {
		Button []MenuButton `json:"button"`
	}
	reqData.Button = buttons
	reqBytes, _ := json.Marshal(reqData)

	var resp BaseResponse
	return client.request("POST", "/menu/create", nil, bytes.NewBuffer(reqBytes), &resp)
}
