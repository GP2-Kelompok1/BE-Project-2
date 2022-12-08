package mysql

import (
	"fmt"
	"immersive-dashboard/config"
	classRepo "immersive-dashboard/features/class/repository"
	feedbackRepo "immersive-dashboard/features/feedback/repository"
	menteeRepo "immersive-dashboard/features/mentee/repository"
	teamRepo "immersive-dashboard/features/team/repository"
	userRepo "immersive-dashboard/features/user/repository"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	return db
}

func DBMigration(db *gorm.DB) {
	db.AutoMigrate(&teamRepo.Team{})
	db.AutoMigrate(&userRepo.User{})
	db.AutoMigrate(&classRepo.Class{})
	db.AutoMigrate(&menteeRepo.Mentee{})
	db.AutoMigrate(&feedbackRepo.Feedback{})

}
