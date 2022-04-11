package settings

import "encoding/json"

type AuthenticationSettings struct {
	keyId     string
	keySecret string
}

func NewAuthenticationSettings(keyId string, keySecret string) *AuthenticationSettings {
	return &AuthenticationSettings{
		keyId:     keyId,
		keySecret: keySecret,
	}
}

func (s *AuthenticationSettings) GetFormattedSettings() string {
	keys := map[string]string{
		"keyId":     s.keyId,
		"keySecret": s.keySecret,
	}
	result, _ := json.Marshal(keys)
	return string(result)
}
