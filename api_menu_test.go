package weixin

import (
	"encoding/json"
	"os"
	"testing"
)

func TestMenuCreate(t *testing.T) {
	appId := os.Getenv("APP_ID")
	appSecret := os.Getenv("APP_SECRET")
	menuJson := os.Getenv("MENU_JSON")
	if appId == "" || appSecret == "" || menuJson == "" {
		t.Error("未找到环境变量APP_ID或APP_SECRET")
		return
	}

	wx := New(appId, appSecret)
	wx.LoadAccessTokenFileCache("weixin-access_token.cache")
	defer wx.SaveAccessTokenFileCache("weixin-access_token.cache")

	var buttons []MenuButton
	err := json.Unmarshal([]byte(menuJson), &buttons)
	if err != nil {
		t.Errorf("无法解析菜单配置:%s", err)
		return
	}

	return

	err = wx.MenuCreate(buttons)
	if err != nil {
		t.Error(err)
	}
}
