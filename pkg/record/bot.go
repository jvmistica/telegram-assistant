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
	d := &RecordDB{DB: db}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Read texts sent to the bot
		if update.Message.Text != "" {
			currentMsg = update.Message.Text
			msg = processMessage(prevMsg, d, update.Message.Chat.ID)
		}

		prevMsg = currentMsg
		msg.ParseMode = "Markdown"
		bot.Send(msg)
	}
}

func processMessage(prevMsg string, d *RecordDB, chatID int64) tgbotapi.MessageConfig {
	if prevMsg == "/additem" {
		result, err := Add(d, []string{currentMsg})
		if err != nil {
			return tgbotapi.NewMessage(chatID, fmt.Sprintf("%s", err))
		}
		return tgbotapi.NewMessage(chatID, result)
	}

	if prevMsg == "/deleteitem" {
		result, err := Delete(d, []string{currentMsg})
		if err != nil {
			return tgbotapi.NewMessage(chatID, fmt.Sprintf("%s", err))
		}
		return tgbotapi.NewMessage(chatID, result)
	}

	cmds, err := d.CheckCommand(currentMsg)
	if err != nil {
		return tgbotapi.NewMessage(chatID, fmt.Sprintf("%s", err))
	}
	return tgbotapi.NewMessage(chatID, cmds)
}
