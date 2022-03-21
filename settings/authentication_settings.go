package httpclient

type AuthenticationSettings struct {
	keyId     string
	keySecret string
}

func newAuthenticationSettings(keyId string, keySecret string) *AuthenticationSettings {
	settings := new(AuthenticationSettings)
	settings.keyId = keyId
	settings.keySecret = keySecret
	return settings
}
