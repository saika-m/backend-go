package database

import "github.com/Takina-Space/backend-go/app/models/entity"

func Migrate() {
	db := PostgreDB
	err := db.AutoMigrate(&entity.ExampleData{})
	if err != nil {
		return
	}
}
