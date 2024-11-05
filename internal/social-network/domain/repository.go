package domain

type SocialNetworkRepository interface {
	Create(socialNetWork SocialNetwork) error
}
