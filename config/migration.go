package config

import (
	"fmt"
	"gorbac/app/entity"
	"log"
)

func Migration(db Database) {
	fmt.Println("Process Migrating...")

	err := db.DB.AutoMigrate(
		&entity.User{},
		&entity.Role{},
		&entity.Menu{},
		&entity.UserRole{},
		&entity.RoleMenu{},
		&entity.Permission{},
		&entity.RolePermission{},
		&entity.Status{},
		&entity.Products{},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database migrated successfully...")
}
