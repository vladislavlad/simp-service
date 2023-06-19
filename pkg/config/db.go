package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"simp-service/pkg/model"
)

func ConnectDB(cfg Config) *gorm.DB {
	db, err := gorm.Open(
		postgres.Open(
			"host="+cfg.Database.Host+" "+
				"port="+cfg.Database.Port+" "+
				"user="+cfg.Database.Username+" "+
				"password="+cfg.Database.Password+" "+
				"dbname="+cfg.Database.Db+" "+
				"sslmode=disable",
		),
		&gorm.Config{},
	)
	if err != nil {
		log.Print(err)
		panic("Failed to connect to DB")
	}

	err = db.AutoMigrate(&model.Comment{})
	if err != nil {
		log.Print(err)
		panic("Failed to migrate DB")
	}

	return db
}
