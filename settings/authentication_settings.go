package settings

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
