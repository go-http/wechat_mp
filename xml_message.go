package weixin

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

//用<![CDATA[和]]包裹的文本
type CDATA struct {
	String string `xml:",cdata"`
}

type XMLMessage struct {
	XMLName xml.Name `xml:"xml"`

	CreateTime int
	MsgId      int

	ToUserName   CDATA
	FromUserName CDATA
	MsgType      CDATA
	Content      CDATA

	PicUrl  CDATA
	MediaId CDATA

	Event    CDATA
	EventKey CDATA

	SendPicsInfo struct {
		Count   int
		PicList []struct {
			PicMd5Sum CDATA
		} `xml:"PicList>item"`
	}
}

func (msg XMLMessage) SaveImageTo(path string) error {
	if msg.MsgType.String != "image" {
		return fmt.Errorf("非图片消息")
	}

	if msg.PicUrl.String == "" {
		return fmt.Errorf("图片地址为空")
	}

	resp, err := http.Get(msg.PicUrl.String)
	if err != nil {
		return fmt.Errorf("请求错误: %s", err)
	}
	defer resp.Body.Close()

	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取错误: %s", err)
	}

	err = ioutil.WriteFile(path, buffer, 0755)
	if err != nil {
		return fmt.Errorf("保存错误: %s", err)
	}

	return nil
}
