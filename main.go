package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Items struct {
	db *gorm.DB
}

func NewItems(db *gorm.DB) *Items {
	return &Items{db: db}
}

func main() {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASS")
	database := os.Getenv("POSTGRES_DB")

	// Connect to the database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, database, port)
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Listen to messages sent to Telegram bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	var oldMsg, txt, tag string
	var msg tgbotapi.MessageConfig
	updates, err := bot.GetUpdatesChan(u)
	i := NewItems(db)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		txt = update.Message.Text
		if oldMsg == "/add item" {
			res, err := i.AddItem(txt)
			if err != nil {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("An error occured. %s", err))
			}
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, res)
		} else if oldMsg == "/delete item" {
			res := i.DeleteItem(txt)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, res)
		} else if oldMsg == "/edit item" {
			match := i.CheckItem(txt)
			if match {
				tag = "edit"
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("What do you want to change \"%s\" to?", txt))
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Item \"%s\" does not exist in the inventory.", txt))
			}
		} else if len(oldMsg) > 5 && oldMsg[:4] == "edit" {
			tag = ""
			i.EditItem(oldMsg[4:], txt)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Item \"%s\" has been changed to \"%s\".", oldMsg[4:], txt))
		} else {
			cmds, err := i.CheckCommand(txt)
			if err != nil {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("An error occured. %s", err))
			}
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, cmds)
		}

		oldMsg = tag + txt
		bot.Send(msg)
	}
}
