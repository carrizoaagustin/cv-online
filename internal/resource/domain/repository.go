package domain

import "github.com/carrizoaagustin/cv-online/internal/resource/domain/model"

type ResourceRepository interface {
	Create(resource model.Resource) error
}
