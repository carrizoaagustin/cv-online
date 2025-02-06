package auth

type UserInformation struct {
	ID    string
	Email string
	Name  string
}

type OAuthConfigParams struct {
	ClientID  string
	SecretKey string
	Scopes    []string
}

type OAuthProvider interface {
	GetAuthorizationURL() string
	GetName() string
	GetUserInformation(accessToken string) (*UserInformation, error)
	RedeemToken(authorizationCode string) (string, error)
}

type LocalProvider interface {
	CheckUserCredentials(username string, password string) bool
	GetUserInformation() UserInformation
}
