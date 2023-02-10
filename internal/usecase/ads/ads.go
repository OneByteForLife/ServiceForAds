package ads

import (
	"ServiceForAds/internal/entity"
)

type (
	// Для service
	Service interface {
		GetOne(id string) (entity.Advertisements, error)
		GetAll(limit string, offset string, sortBy string, sortType string) ([]entity.Advertisements, error)
		Create(body []byte) error
	}

	// Для storage
	Storage interface {
		GetOne(id int) (entity.Advertisements, error)
		GetAll(limit int, offset int, sortBy string, sortType string) ([]entity.Advertisements, error)
		Create(ads entity.Advertisements) error
	}
)
