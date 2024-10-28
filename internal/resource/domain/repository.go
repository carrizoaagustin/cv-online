package domain

type ResourceRepository interface {
	Create(resource Resource) error
}
