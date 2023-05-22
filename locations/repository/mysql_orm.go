package mysql_orm

import (
	"fmt"

	"github.com/marceloaguero/go-odyssey-voyage-I/locations/graph/model"
	"github.com/marceloaguero/go-odyssey-voyage-I/locations/usecase"
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

	db.AutoMigrate(&model.Location{})

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

func (r *ormRepo) Create(location *model.Location) (*model.Location, error) {
	result := r.db.Create(&location)
	return location, result.Error
}

func (r *ormRepo) GetByName(name string) (*model.Location, error) {
	var location model.Location
	result := r.db.Take(&location, "name = ?", name)
	return &location, result.Error
}
