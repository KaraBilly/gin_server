package service

type weChatResponse struct {
	Token   string `json:"access_token"`
	Expire  int    `json:"expires_in"`
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
