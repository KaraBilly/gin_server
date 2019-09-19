package service

import (
	"encoding/json"
	"gin_server/pkg/cache"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func GetAccessToken() (string, error) {

	token, exist := cache.GetTokenCache(cache.AccessToken)
	if exist {
		return token.(string), nil
	}

	wechatClient, err := url.Parse(viper.GetString("wechat.base_url") + viper.GetString("wechat.get_token"))
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("grant_type", viper.GetString("wechat.grant_type"))
	params.Set("appid", viper.GetString("wechat.appid"))
	params.Set("secret", viper.GetString("wechat.secret"))
	wechatClient.RawQuery = params.Encode()
	urlPath := wechatClient.String()

	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result weChatResponse
	err = json.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}

	cache.SetTokenCache(cache.AccessToken, result.Token, time.Duration(result.Expire)*time.Second)
	return result.Token, nil
}

//
//func SendMessage(r *wechat.SendMessageRequest) error {
//	//var request = &r
//	//wechatClient,err:= url.Parse(viper.GetString("wechat.base_url") + viper.GetString("wechat.send_message"))
//	//if err == nil{
//	//	return err
//	//}
//	return nil
//}
