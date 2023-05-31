package mysql_orm

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/marceloaguero/go-odyssey-voyage-I/reviews/graph/model"
	"github.com/marceloaguero/go-odyssey-voyage-I/reviews/usecase"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ormRepo struct {
	db *gorm.DB
}

// NewRepo crea un repositorio implementado en ORM (MySQL)
func NewRepo(dsName, dbName string) (usecase.Repository, error) {
	db, err := dbConnect(dsName, dbName)
	if err != nil {
		return nil, errors.Wrap(err, "MySQL ORM - Can't connect to DB")
	}

	db.AutoMigrate(&model.Review{})

	return &ormRepo{
		db: db,
	}, nil
}

func dbConnect(dsName, dbName string) (*gorm.DB, error) {
	conn := fmt.Sprintf("%s/%s?charset=utf8&parseTime=True&loc=Local", dsName, dbName)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (r *ormRepo) Create(input *model.LocationReviewInput) (*model.Review, error) {
	var review *model.Review

	review.ID = uuid.NewString()
	review.Comment = input.Comment
	review.Rating = input.Rating
	review.LocationID = input.LocationID

	result := r.db.Create(&review)
	return review, result.Error
}

func (r *ormRepo) GetByID(id string) (*model.Review, error) {
	var review model.Review
	result := r.db.Take(&review, "id = ?", id)
	return &review, result.Error
}

func (r *ormRepo) GetByLocationID(locationID string) ([]*model.Review, error) {
	reviews := []*model.Review{}
	result := r.db.Find(&reviews, "location_id = ?", locationID)
	return reviews, result.Error
}

func (r *ormRepo) GetAll() ([]*model.Review, error) {
	reviews := []*model.Review{}
	result := r.db.Find(&reviews)
	return reviews, result.Error
}
