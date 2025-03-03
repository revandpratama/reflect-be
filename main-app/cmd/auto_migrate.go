package cmd

import (
	"fmt"

	"github.com/revandpratama/reflect/internal/entities"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {

	createSchemaIfNotExists(db, "public")

	db.AutoMigrate(&entities.Post{}, &entities.Comment{})
}

func createSchemaIfNotExists(db *gorm.DB, schemaName string) {
	db.Exec(fmt.Sprintf(`CREATE SCHEMA IF NOT EXISTS "%s"`, schemaName))
}
