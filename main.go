package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	oldMsg string
	txt    string
	msg    tgbotapi.MessageConfig
	err    error
	db     *gorm.DB
)

func main() {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASS")
	database := os.Getenv("POSTGRES_DB")

	// Connect to the database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, database, port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(Item{})

	// Listen to messages sent to Telegram bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	i := NewItems(db)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		txt = update.Message.Text
		if oldMsg == "/additem" {
			res, err := i.AddItem([]string{txt})
			if err != nil {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", err))
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, res)
			}
		} else if oldMsg == "/deleteitem" {
			res, err := i.DeleteItem([]string{txt})
			if err != nil {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", err))
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, res)
			}
		} else {
			cmds, err := i.CheckCommand(txt)
			if err != nil {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", err))
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, cmds)
			}
		}

		oldMsg = txt
		msg.ParseMode = "Markdown"
		bot.Send(msg)
	}
}
