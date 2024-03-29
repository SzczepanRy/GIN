package initializers

import "jwt.com/api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
