package main

import (
	"bufio"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

type User struct {
	gorm.Model
	ChatID int64 `gorm:"unique_index"`
}

func main() {
	var err error
	bot, err = tgbotapi.NewBotAPI("6096077916:AAFu9hNesAGnYvhaIaOAnOXCCpsMU2bVgtU")
	if err != nil {
		// Abort if something is wrong
		log.Panic(err)
	}

	// Set this to true to log all interactions with telegram servers
	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Create a new cancellable background context. Calling `cancel()` leads to the cancellation of the context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// `updates` is a golang channel which receives telegram updates
	updates := bot.GetUpdatesChan(u)

	// Pass cancellable context to goroutine
	go ReceiveUpdates(ctx, updates)

	// Tell the user the bot is online
	log.Println("Start listening for updates. Press enter to stop")

	// Wait for a newline symbol, then cancel handling updates
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()

}
