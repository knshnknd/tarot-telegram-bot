package main

import (
	"flag"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)


var (
	// глобальная переменная, в которой храним токен
	telegramBotToken    string
)

func init() {
	// меняем BOT_TOKEN на токен бота от BotFather, в строке принимаем на входе флаг -telegrambottoken
	flag.StringVar(&telegramBotToken, "telegrambottoken", "", "Telegram Bot Token")
	flag.Parse()

	// без флага не запускаем
	if telegramBotToken == "" {
		log.Print("-telegrambottoken is required")
		os.Exit(1)
	}
}


func main() {
	// используя токен, создаем новый инстанс бота
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	// пишем об этом в консоль
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг, создаем канал, в который будут прилетать новые сообщения
	updates := bot.GetUpdatesChan(u)

	// в канал updates прилетают структуры типа Update, вычитываем их и обрабатываем
	for update := range updates {
		if update.Message == nil {
			continue
		}

		logUpdate(update)

		reply := ""
		if update.Message.IsCommand() {
			reply = processCommand(update)
		}

		sendReplyToUpdate(update, reply, bot)
	}
}


func logUpdate(update tgbotapi.Update) {
	message := update.Message.Text
	userName := update.Message.From.UserName
	chatID := update.Message.Chat.ID
	chatTitle := update.Message.Chat.Title

	log.Printf("[%s] sent message: \"%s\" to chat: \"%s\"[%d]", userName, message, chatTitle, chatID)
}


func processCommand(update tgbotapi.Update) string {
	command := update.Message.Command()
	reply := ""

	switch command {
	case "start":
		reply = "Start message of the Tarot Bot!"
	}
	return reply
}

func sendReplyToUpdate(update tgbotapi.Update, reply string, bot *tgbotapi.BotAPI) {
	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, reply)
	bot.Send(msg)
}
