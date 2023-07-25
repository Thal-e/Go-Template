package migrations

import (
	"github.com/Thal-e/go-template/api/db"
	"github.com/Thal-e/go-template/pkg/models"
)

func Migrate() {
	database := db.GetDB()
	database.Raw("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	database.AutoMigrate(&models.Users{}, &models.Items{})
}
