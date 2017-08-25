package weixin

type Client struct {
	accessTokenInfo AccessTokenInfo

	AppId     string
	AppSecret string
}

func New(appId, appSecret string) *Client {
	return &Client{AppId: appId, AppSecret: appSecret}
}
