package initializers

import (
	"github.com/lordofthemind/ginJwtAuthPostgres/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
