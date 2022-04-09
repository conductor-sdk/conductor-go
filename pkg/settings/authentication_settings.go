package settings

import "encoding/json"

type AuthenticationSettings struct {
	KeyId     string
	KeySecret string
}

func NewAuthenticationSettings(keyId string, keySecret string) *AuthenticationSettings {
	settings := new(AuthenticationSettings)
	settings.KeyId = keyId
	settings.KeySecret = keySecret
	return settings
}

func (s *AuthenticationSettings) GetFormattedSettings() string {
	keys := map[string]string{
		"keyId":     s.KeyId,
		"keySecret": s.KeySecret,
	}
	result, _ := json.Marshal(keys)
	return string(result)
}
