package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"simp-service/pkg/model"
)

func ConnectDB(cfg Config) *gorm.DB {
	db, err := gorm.Open("postgres",
		"host="+cfg.Database.Host+" "+
			"port="+cfg.Database.Port+" "+
			"user="+cfg.Database.Username+" "+
			"password="+cfg.Database.Password+" "+
			"dbname="+cfg.Database.Db+" "+
			"sslmode=disable",
	)
	if err != nil {
		log.Print(err)
		panic("Failed to connect to ConnectDB")
	}

	db.AutoMigrate(&model.Comment{})

	return db
}
