package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jvmistica/henchmaid/pkg/item"
	"github.com/jvmistica/henchmaid/pkg/record"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	db.AutoMigrate(item.ItemRecord{})

	// Listen to messages sent to Telegram bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	i := &item.Item{DB: db}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		txt = update.Message.Text
		switch oldMsg {
		case "/additem":
			res, err := record.Add(i, []string{txt})
			if err != nil {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", err))
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, res)
			}
		case "/deleteitem":
			res, err := record.Delete(i, []string{txt})
			if err != nil {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", err))
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, res)
			}
		default:
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
