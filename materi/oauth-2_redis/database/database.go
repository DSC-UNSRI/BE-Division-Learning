package database

import (
	"context"
	"fmt"
	"oauth-2_redis/models"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Rdb *redis.Client

func Init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FJakarta",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	if dsn == "" {
		panic("DB_DSN not set in .env file")
	}

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected")
}

func Redis() {
	ctx := context.Background()
	dbNum, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		dbNum = 0
	}

	Rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNum,
	})

	pong, err := Rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis connected:", pong)
}

func Migrate() {
	if err := DB.Debug().AutoMigrate(&models.User{}, &models.Event{}); err != nil {
		panic(err)
	}
	fmt.Println("Migrate Successfuly")
}
