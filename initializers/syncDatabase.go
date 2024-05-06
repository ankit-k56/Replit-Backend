package initializers

import "github.com/ankit-k56/Repelit/models"

func SyncDatabase() {
	Db.AutoMigrate(&models.User{}, &models.Project{})
}