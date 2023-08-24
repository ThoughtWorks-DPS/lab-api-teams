package datastore

import (
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
	DB gorm.DB
}

func NewGormDatastore() GormDatastore {
	host := os.Getenv("DATABASE_URL")
	fmt.Println(host)
	dsn := fmt.Sprintf("host=%s user=postgres dbname=postgres port=5433 sslmode=disable TimeZone=Asia/Shanghai", host)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &gormDatastoreImpl{
		DB: *db,
	}
}

func (g *gormDatastoreImpl) Migrate(models ...interface{}) error {
	g.DB.AutoMigrate(models...)
	return g.DB.Error
}

func (g *gormDatastoreImpl) Create(data interface{}) error {
	return g.DB.Create(data).Error
}

func (g *gormDatastoreImpl) ReadByID(id string, out interface{}) error {
	return g.DB.First(out, id).Error
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
