package config

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/its-lana/coffee-shop/logger"
	"github.com/its-lana/coffee-shop/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	databaseLogger "gorm.io/gorm/logger"
)

type GormDatabase struct {
	DB *gorm.DB
}

func NewPG(ctx context.Context, file *os.File) (*GormDatabase, error) {
	connString := LoadConnString()
	dbLogger := NewDatabaseLogger(file)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		PrepareStmt: true,
		Logger:      dbLogger,
	})
	if err != nil {
		logger.Log.Errorf("unable to create connection pool: %w", err.Error())
		return nil, err
	}
	return &GormDatabase{
		DB: db,
	}, nil
}

func (gd *GormDatabase) MigratingDatabase() {
	gd.DB.AutoMigrate(&model.Customer{}, &model.Merchant{}, &model.Category{}, &model.Menu{}, &model.OrderItem{}, &model.Cart{})
}

func LoadConnString() string {
	requiredEnvVars := []string{"DB_USERNAME", "DB_PASSWORD", "DB_PORT", "DB_HOST", "DB_NAME"}
	config := make(map[string]string)

	// Load environment variables and check if they are set
	for _, env := range requiredEnvVars {
		value := os.Getenv(env)
		if value == "" {
			log.Fatalf("Environment variable %s not set", env)
		}
		config[env] = value
	}

	// Load the time zone location
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal(err)
	}

	// Construct the connection string
	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=%s",
		config["DB_USERNAME"],
		config["DB_PASSWORD"],
		config["DB_HOST"],
		config["DB_PORT"],
		config["DB_NAME"],
		location.String(),
	)

	return connectionString
}

func NewDatabaseLogger(file *os.File) databaseLogger.Interface {
	return databaseLogger.New(
		log.New(io.MultiWriter(os.Stdout, file), "\r\n", log.LstdFlags),
		databaseLogger.Config{
			SlowThreshold:             250 * time.Millisecond,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  databaseLogger.Info,
			Colorful:                  true,
		},
	)

}
