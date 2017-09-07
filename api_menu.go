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
	Type string `json:"type,omitempty"`
	Name string `json:"name"`

	Key      string `json:"key,omitempty"`
	Url      string `json:"url,omitempty"`
	AppId    string `json:"appid,omitempty"`
	MediaId  string `json:"media_id,omitempty"`
	PagePath string `json:"pagepath,omitempty"`

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
