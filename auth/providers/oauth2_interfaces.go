package providers

type Endpoints struct {
	AuthorizationURL string
	TokenURL         string
	UserInfoURL      string
}

type oauth2Config struct {
	ClientID  string
	SecretKey string
	Name      string
	Endpoints Endpoints
	Scopes    []string
}
