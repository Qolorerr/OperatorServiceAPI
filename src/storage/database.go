package storage

import (
	"fmt"
	"operator_text_channel/src/storage/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(cfg PGConfig) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Moscow",
		cfg.Addr, cfg.Port, cfg.User, cfg.Password, cfg.DbName)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, err
	}
	err = autoMigrateModels(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func autoMigrateModels(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Appeal{},
		&models.Tag{},
		&models.OperatorGroup{},
		&models.Operator{},
		&models.AppealTagLink{},
		&models.GroupTagLink{},
		&models.OperatorGroupLink{},
	)
	return err
}
