package initializers

import "github.com/abhinavthapa1998/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
