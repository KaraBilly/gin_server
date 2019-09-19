package wechat

type GetTokenResponse struct {
	Token string `json:"token"`
}

type SendMessageRequest struct {
	OpenId          string      `json:"open_id"`
	TemplateId      string      `json:"template_id"`
	Page            string      `json："page"`
	FormId          string      `json:"form_id"`
	Data            interface{} `json:"data"`
	EmphasisKeyword string      `json:"emphasis_keyword"`
}
