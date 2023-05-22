package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/marceloaguero/go-odyssey-voyage-I/reviews/graph/model"
	"github.com/pkg/errors"
)

type Repository interface {
	Create(review *model.Review) (*model.Review, error)
	GetByID(id string) (*model.Review, error)
	GetAll() ([]*model.Review, error)
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

func (u *usecase) Create(review *model.Review) (*model.Review, error) {
	validate := validator.New()
	err := validate.Struct(review)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return nil, errors.Wrap(validationErrors, "UC - Create - Error during review data validation")
	}

	review.ID = uuid.NewString()
	review, err = u.repository.Create(review)
	if err != nil {
		return nil, errors.Wrap(err, "UC - Create - Error creating a new review")
	}

	return review, nil
}

func (u *usecase) GetByID(id string) (*model.Review, error) {
	review, err := u.repository.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "UC - GetByID - Error fetching a review")
	}

	return review, nil
}

func (u *usecase) GetAll() ([]*model.Review, error) {
	reviews, err := u.repository.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "UC - GetAll - Error fetching all reviews")
	}

	return reviews, nil
}
