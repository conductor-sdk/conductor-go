package settings

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

func (s *AuthenticationSettings) GetBody() map[string]string {
	body := map[string]string{
		"keyId":     s.keyId,
		"keySecret": s.keySecret,
	}
	return body
}
