package migrate

import (
	"github.com/grahamquan/go-crud/models"
	"github.com/grahamquan/go-crud/setup"
)

func MigrateDB() {
	setup.DB.AutoMigrate(models.Post{})
}
