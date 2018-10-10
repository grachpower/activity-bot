package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const BOT_TOKEN = "646189272:AAEFTGLNqqVXZc_RKCtwC5gJ7XlQBJR7XLA"

const TYPE_PATTERN = "ты "

func main() {
	bot, err := tgbotapi.NewBotAPI(BOT_TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// инициализируем канал, куда будут прилетать обновления от API
	ucfg := tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60

	updates, err := bot.GetUpdatesChan(ucfg)

	for {
		select {
		case update := <-updates:
			// Пользователь, который написал боту
			UserName := update.Message.From.UserName

			// ID чата/диалога.
			// Может быть идентификатором как чата с пользователем
			// (тогда он равен UserID) так и публичного чата/канала
			ChatID := update.Message.Chat.ID

			// Текст сообщения
			Text := update.Message.Text

			switch update.Message.Command() {
			case "init":
				msg := tgbotapi.NewMessage(ChatID, "пошли нахуй")
				bot.Send(msg)
			case "chatid":
				msg := tgbotapi.NewMessage(ChatID, "Chat id: "+strconv.FormatInt(ChatID, 10))
				bot.Send(msg)
			default:
				isAvailable := isAvailableMessage(Text)

				if !isAvailable {
					continue
				}

				fmt.Println("Text is available: " + strconv.FormatBool(isAvailable))

				log.Printf("[%s] %d %s", UserName, ChatID, Text)

				// Ответим пользователю его же сообщением
				reply := createReply(Text)
				// Созадаем сообщение
				msg := tgbotapi.NewMessage(ChatID, reply)
				msg.ReplyToMessageID = update.Message.MessageID
				// и отправляем его
				bot.Send(msg)
			}
		}

	}
}

func createReply(text string) string {
	itemIndex := strings.Index(strings.ToLower(text), TYPE_PATTERN)

	return "нет, " + text[itemIndex:len(text)]
}

func isAvailableMessage(text string) bool {
	return strings.Contains(strings.ToLower(text), TYPE_PATTERN)
}
