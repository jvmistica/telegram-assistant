package record

import (
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	currentMsg string
	prevMsg    string
	msg        tgbotapi.MessageConfig
)

// Listen listens to the messages send to the bot and sends the appropriate response
func (r *RecordDB) Listen(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Read texts sent to the bot
		if update.Message.Text != "" {
			currentMsg = update.Message.Text
			msg = r.processMessage(prevMsg, update.Message.Chat.ID)
		}

		prevMsg = currentMsg
		msg.ParseMode = "Markdown"
		bot.Send(msg)
	}
}

// processMessage processes the message according to the commands and values provided
func (r *RecordDB) processMessage(prevMsg string, chatID int64) tgbotapi.MessageConfig {
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
