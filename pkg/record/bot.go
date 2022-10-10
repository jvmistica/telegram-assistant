package record

import (
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
)

var (
	currentMsg string
	prevMsg    string
	msg        tgbotapi.MessageConfig
)

func Listen(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, db *gorm.DB) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Read texts sent to the bot
		if update.Message.Text != "" {
			currentMsg = update.Message.Text
			msg = processMessage(prevMsg, db, update.Message.Chat.ID)
		}

		prevMsg = currentMsg
		msg.ParseMode = "Markdown"
		bot.Send(msg)
	}
}

func processMessage(prevMsg string, db *gorm.DB, chatID int64) tgbotapi.MessageConfig {
	r := &RecordDB{DB: db}
	if prevMsg == "/additem" {
		result, err := r.Add([]string{currentMsg})
		if err != nil {
			return tgbotapi.NewMessage(chatID, fmt.Sprintf("%s", err))
		}
		return tgbotapi.NewMessage(chatID, result)
	}

	if prevMsg == "/deleteitem" {
		result, err := r.Delete([]string{currentMsg})
		if err != nil {
			return tgbotapi.NewMessage(chatID, fmt.Sprintf("%s", err))
		}
		return tgbotapi.NewMessage(chatID, result)
	}

	cmds, err := r.CheckCommand(currentMsg)
	if err != nil {
		return tgbotapi.NewMessage(chatID, fmt.Sprintf("%s", err))
	}
	return tgbotapi.NewMessage(chatID, cmds)
}
