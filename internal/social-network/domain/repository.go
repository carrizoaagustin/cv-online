package domain

import (
	"github.com/carrizoaagustin/cv-online/internal/social-network/domain/model"
)

type SocialNetworkRepository interface {
	Create(socialNetWork model.SocialNetwork) error
}
