package weixin

import (
	"encoding/json"
)

type TemplateMessageData struct {
	Value string `json:"value"`
	Color string `json:"color"` //#rrggbb
}

type MiniProgram struct {
	Appid    string `json:"appid,omitempty"`
	PagePath string `json:"pagepath,omitempty"`
}

type TemplateMessage struct {
	ToUser      string                         `json:"touser"` //收件人的OPENID
	TemplateId  string                         `json:"template_id"`
	Url         string                         `json:"url,omitempty"`
	MiniProgram *MiniProgram                   `json:"miniprogram,omitempty"`
	Data        map[string]TemplateMessageData `json:"data,omitempty"`
}

func (client *Client) SendTemplateMessage(msg *TemplateMessage) (int, error) {
	reqBytes, _ := json.Marshal(msg)

	var resp struct {
		BaseResponse
		MsgId int
	}

	err := client.request("POST", "/message/template/send", nil, reqBytes, &resp)

	return resp.MsgId, err
}

func NewTemplateMessage(templateId, toOpenId string) *TemplateMessage {
	return &TemplateMessage{
		ToUser:     toOpenId,
		TemplateId: templateId,
		Data:       make(map[string]TemplateMessageData),
	}
}

func (msg *TemplateMessage) SetData(key, value, color string) {
	data := TemplateMessageData{Value: value, Color: color}
	msg.Data[key] = data
}

func (msg *TemplateMessage) SetMiniProgram(appId, pagePath string) {
	msg.MiniProgram = &MiniProgram{Appid: appId, PagePath: pagePath}
}
