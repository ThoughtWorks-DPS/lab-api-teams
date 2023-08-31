package datastore

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDatastore interface {
	Datastore
	Migrate(models ...interface{}) error
}

type gormDatastoreImpl struct {
	DB     gorm.DB
	domain string
}

func NewGormDatastore(domain string) GormDatastore {
	host := os.Getenv("DATABASE_URL")
	dsn := fmt.Sprintf("host=%s user=gorm dbname=gorm password=gorm port=5432 sslmode=disable TimeZone=UTC", host)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &gormDatastoreImpl{
		DB:     *db,
		domain: domain,
	}
}

func (g *gormDatastoreImpl) ReadByAttributes(filter Filter, out interface{}) error {
	return g.DB.Where(filter).Find(out).Error
}

func (g *gormDatastoreImpl) Migrate(models ...interface{}) error {
	err := g.DB.AutoMigrate(models...)
	if err != nil {
		return err
	}
	return nil
}

func (g *gormDatastoreImpl) Create(data interface{}) error {
	return g.DB.Create(data).Error
}

// ReadByID retrieves a record by its ID
func (g *gormDatastoreImpl) ReadByID(id string, out interface{}) error {
	if err := g.DB.First(out, fmt.Sprintf("%s_id = ?", g.domain), id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}
		return err
	}
	return nil
}

func (g *gormDatastoreImpl) ReadAll(out interface{}) error {
	return g.DB.Find(out).Error
}

func (g *gormDatastoreImpl) Update(data interface{}) error {
	return g.DB.Save(data).Error
}

func (g *gormDatastoreImpl) Delete(data interface{}) error {
	return g.DB.Delete(data).Error
}

func (g *gormDatastoreImpl) IsDatabaseAvailable() (bool, error) {
	sqlDB, err := g.DB.DB()
	if err != nil {
		return false, fmt.Errorf("DB was not connectable: %v", err)
	}
	err = sqlDB.Ping()
	if err != nil {
		return false, fmt.Errorf("DB was not reachable: %v", err)
	}
	return true, nil
}

var ErrNotFound = errors.New("record not found")
