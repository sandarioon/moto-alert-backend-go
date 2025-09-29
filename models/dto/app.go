package dto

type GetAboutResponse struct {
	Status  int       `json:"status" example:"200"`
	Data    AboutData `json:"data"`
	Message string    `json:"message" example:"OK"`
}

type AboutData struct {
	Text string `json:"text"`
}

type SocialLinkResponse struct {
	Status  int          `json:"status" example:"200"`
	Data    []SocialLink `json:"data"`
	Message string       `json:"message" example:"OK"`
}

type SocialLink struct {
	Type string `json:"type" example:"Telegram"`
	Name string `json:"name" example:"Канал в Telegram"`
	Link string `json:"link" example:"https://t.me/moto_alert"`
}

type GetSettingsResponse struct {
	Status  int         `json:"status" example:"200"`
	Data    AppSettings `json:"data"`
	Message string      `json:"message" example:"OK"`
}

type AppSettings struct {
	Env     string `json:"env" example:"development"`
	Version string `json:"version" example:"1.0.0"`
}
