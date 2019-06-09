package main

import (
	"fmt"

	"github.com/eventlist/config"
	"github.com/eventlist/dao"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Setup Database
	postgresConfig := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.DBHost, config.Port,
		config.DBName, config.Username,
		config.Password, "disable")
	db, err := gorm.Open("postgres", postgresConfig)
	if err != nil {
		panic(fmt.Sprintf("Can not create db client: %v", err))
	}
	defer db.Close()

	// Proccess Migrations
	db.AutoMigrate(&dao.Category{})
	db.AutoMigrate(&dao.Event{}).AddForeignKey("category_id", "category(id)", "RESTRICT", "RESTRICT")
}
