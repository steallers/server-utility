package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseService struct {
	Connectors *gorm.DB
}
type DatabaseServiceInterface interface {
	GetConnectors() *gorm.DB
}

func StartDatabaseService() DatabaseServiceInterface {
	dsn := "host=localhost user=postgres password=123123 dbname=steallers port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(errors.New(fmt.Sprintf("failed to start database with error: -> %e", err)))
	}
	var NewService DatabaseService
	NewService.Connectors = database
	return NewService
}

func (db DatabaseService) GetConnectors() *gorm.DB {
	return db.Connectors
}
