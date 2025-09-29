package app

import "github.com/spf13/viper"

type service struct {
}

type Service interface {
	GetAbout() (string, error)
	GetSocialLinks() ([]SocialLink, error)
	GetSettings() (AppSettings, error)
}

func NewService() Service {
	return service{}
}

func (s service) GetAbout() (string, error) {
	text := "Наша цель: помогать в спасении людей.\n\nПриложение создано и поддерживается энтузиастами за свой счёт — некоммерческое, без рекламы, всегда бесплатно. Каждая функция — для быстрой помощи при аварии.\n\nНапишите нам о ваших идеях или проблемах с приложением."

	return text, nil
}

type SocialLink struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Link string `json:"link"`
}

func (s service) GetSocialLinks() ([]SocialLink, error) {
	socialLinks := []SocialLink{
		{
			Type: "Telegram",
			Name: "Канал в Telegram",
			Link: "https://t.me/moto_alert",
		},
		{
			Type: "Telegram",
			Name: "Тех. поддержка",
			Link: "https://t.me/moto_alert_support",
		},
		{
			Type: "Telegram",
			Name: "Чат",
			Link: "https://t.me/moto_alert_chat",
		},
		{
			Type: "VK",
			Name: "ВКонтакте",
			Link: "https://vk.com/moto_alert",
		},
	}

	return socialLinks, nil
}

type AppSettings struct {
	Env     string `json:"env"`
	Version string `json:"version"`
}

func (s service) GetSettings() (AppSettings, error) {
	settings := AppSettings{
		Env:     viper.GetString("general.env"),
		Version: "1.0.0",
	}

	return settings, nil
}
