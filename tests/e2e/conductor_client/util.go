package conductor_client

import "github.com/conductor-sdk/conductor-go/pkg/settings"

func GetAuthenticationSettings() *settings.AuthenticationSettings {
	return settings.NewAuthenticationSettings(
		"",
		"",
	)
}

func GetHttpSettingsWithAuth() *settings.HttpSettings {
	return settings.NewHttpSettings(
		"https://play.orkes.io",
		nil,
	)
}
