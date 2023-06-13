package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jvmistica/telegram-assistant/pkg/record"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	bot *tgbotapi.BotAPI
	db  *gorm.DB
	err error
)

func init() {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		log.Fatal("missing environment variable POSTGRES_HOST")
	}

	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		log.Fatal("missing environment variable POSTGRES_PORT")
	}

	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		log.Fatal("missing environment variable POSTGRES_USER")
	}

	password := os.Getenv("POSTGRES_PASS")
	if password == "" {
		log.Fatal("missing environment variable POSTGRES_PASS")
	}

	database := os.Getenv("POSTGRES_DB")
	if database == "" {
		log.Fatal("missing environment variable POSTGRES_DB")
	}

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("missing environment variable BOT_TOKEN")
	}

	// Connect to the database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, database, port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&record.Item{})

	bot, err = tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Listen to messages sent to Telegram bot
	config := tgbotapi.NewUpdate(0)
	config.Timeout = 60

	updates, err := bot.GetUpdatesChan(config)
	if err != nil {
		log.Fatal(err)
	}

	r := &record.RecordDB{DB: db}
	r.Listen(updates, bot)
}
