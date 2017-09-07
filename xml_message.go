package weixin

import (
	"encoding/xml"
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
