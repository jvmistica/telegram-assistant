package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
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

	db.AutoMigrate(&record.Item{})

	// Listen to messages sent to Telegram bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	d := &record.RecordDB{DB: db}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Read texts sent to the bot
		if update.Message.Text != "" {
			txt = update.Message.Text
			switch oldMsg {
			case "/additem":
				res, err := record.Add(d, []string{txt})
				if err != nil {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", err))
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, res)
				}
			case "/deleteitem":
				res, err := record.Delete(d, []string{txt})
				if err != nil {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", err))
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, res)
				}
			default:
				cmds, err := d.CheckCommand(txt)
				if err != nil {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", err))
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, cmds)
				}
			}
		}

		// Read documents sent to the bot
		if update.Message.Document != nil {
			url, err := bot.GetFileDirectURL(update.Message.Document.FileID)
			if err != nil {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", err))
			}

			res, err := http.Get(url)
			if err != nil {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", err))
			}

			defer res.Body.Close()
			rcsv := csv.NewReader(res.Body)
			contents, err := rcsv.ReadAll()
			if err != nil {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", err))
			}

			res2, err := record.Import(d, contents)
			if err != nil {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", err))
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, res2)
			}

		}

		oldMsg = txt
		msg.ParseMode = "Markdown"
		bot.Send(msg)
	}
}
