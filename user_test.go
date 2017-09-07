package weixin

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestUserList(t *testing.T) {
	appId := os.Getenv("APP_ID")
	appSecret := os.Getenv("APP_SECRET")
	if appId == "" || appSecret == "" {
		t.Error("未找到环境变量APP_ID或APP_SECRET")
		return
	}

	wx := New(appId, appSecret)
	wx.LoadAccessTokenFileCache("weixin-access_token.cache")
	defer wx.SaveAccessTokenFileCache("weixin-access_token.cache")

	openIds, err := wx.UserList()
	if err != nil {
		t.Error(err)
	}

	t.Log(openIds)
}

func TestUserInfo(t *testing.T) {
	appId := os.Getenv("APP_ID")
	appSecret := os.Getenv("APP_SECRET")
	if appId == "" || appSecret == "" {
		t.Error("未找到环境变量APP_ID或APP_SECRET")
		return
	}

	wx := New(appId, appSecret)
	wx.LoadAccessTokenFileCache("weixin-access_token.cache")
	defer wx.SaveAccessTokenFileCache("weixin-access_token.cache")

	openIds, err := wx.UserList()
	if err != nil {
		t.Error(err)
	}

	for _, openId := range openIds {
		info, err := wx.UserInfo(openId)
		if err != nil {
			t.Error(err)
		}
		t.Log(info)
	}
}

func TestUserBatchInfo(t *testing.T) {
	appId := os.Getenv("APP_ID")
	appSecret := os.Getenv("APP_SECRET")
	if appId == "" || appSecret == "" {
		t.Error("未找到环境变量APP_ID或APP_SECRET")
		return
	}

	wx := New(appId, appSecret)
	wx.LoadAccessTokenFileCache("weixin-access_token.cache")
	defer wx.SaveAccessTokenFileCache("weixin-access_token.cache")

	openIds, err := wx.UserList()
	if err != nil {
		t.Error(err)
	}

	info, err := wx.UserBatchInfo(openIds)
	if err != nil {
		t.Error(err)
	}
	t.Log(info)
}

func TestUserUpdateRemark(t *testing.T) {
	appId := os.Getenv("APP_ID")
	appSecret := os.Getenv("APP_SECRET")
	if appId == "" || appSecret == "" {
		t.Error("未找到环境变量APP_ID或APP_SECRET")
		return
	}

	wx := New(appId, appSecret)
	wx.LoadAccessTokenFileCache("weixin-access_token.cache")
	defer wx.SaveAccessTokenFileCache("weixin-access_token.cache")

	openId := os.Getenv("REMARK_USER_OPENID")

	info, err := wx.UserInfo(openId)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Before: %s", info.Remark)

	err = wx.UserUpdateRemark(openId, fmt.Sprintf("remark_test_%d", time.Now().Unix()))
	if err != nil {
		t.Error(err)
	}

	info, err = wx.UserInfo(openId)
	if err != nil {
		t.Error(err)
	}

	t.Logf("After: %s", info.Remark)
}
