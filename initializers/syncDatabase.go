package initializers

import (
	"github.com/crud-go/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Post{})
}
