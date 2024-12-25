package postgres

import (
	"database/sql"
	"fmt"
	"github.com/qahvazor/qahvazor/internal/model"
	"log"
	"time"

	"github.com/qahvazor/qahvazor/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DB,
	)

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:                 sqlDB,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}

	log.Println("Postgres connection successfully done!")

	err = db.AutoMigrate(&model.Subscription{}, &model.UserSubscription{})
	if err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}
	log.Println("Migration completed successfully!")

	return db
}
