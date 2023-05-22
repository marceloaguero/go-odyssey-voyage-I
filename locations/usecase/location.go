package usecase

import (
	"github.com/marceloaguero/go-odyssey-voyage-I/locations/graph/model"
)

type Repository interface {
	Create(location *model.LocationInput) (*model.Location, error)
}
