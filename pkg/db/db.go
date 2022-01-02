package db

import (
	"github.com/Yefhem/basic-api-with-go/pkg/helpers"
	"github.com/Yefhem/basic-api-with-go/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	dbURI := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	helpers.CheckError(err)

	db.AutoMigrate(&models.Album{})

	return db
}
