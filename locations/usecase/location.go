package usecase

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/marceloaguero/go-odyssey-voyage-I/locations/graph/model"
	"github.com/pkg/errors"
)

type Repository interface {
	Create(location *model.LocationInput) (*model.Location, error)
	GetByID(id string) (*model.Location, error)
	GetByName(name string) (*model.Location, error)
	GetAll() ([]*model.Location, error)
}

type Usecase interface {
	Repository
}

type usecase struct {
	repository Repository
}

func NewUsecase(repo Repository) Usecase {
	return &usecase{
		repository: repo,
	}
}

func (u *usecase) Create(input *model.LocationInput) (*model.Location, error) {

	// Verify name uniqueness
	input.Name = strings.TrimSpace(input.Name)
	_, err := u.GetByName(input.Name)
	if err == nil {
		return nil, errors.Errorf("UC - Create - Location with name %s already exists", input.Name)
	}

	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return nil, errors.Wrap(validationErrors, "UC - Create - Error during location data validation")
	}

	location, err := u.repository.Create(input)
	if err != nil {
		return nil, errors.Wrap(err, "UC - Create - Error creating a new location")
	}

	return location, nil
}

func (u *usecase) GetByID(id string) (*model.Location, error) {
	location, err := u.repository.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "UC - GetByID - Error fetching a location")
	}

	return location, nil
}

func (u *usecase) GetByName(name string) (*model.Location, error) {
	location, err := u.repository.GetByName(name)
	if err != nil {
		return nil, errors.Wrap(err, "UC - GetByName - Error fetching a location")
	}

	return location, nil
}

func (u *usecase) GetAll() ([]*model.Location, error) {
	locations, err := u.repository.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "UC - GetAll - Error fetching all locations")
	}

	return locations, nil
}
