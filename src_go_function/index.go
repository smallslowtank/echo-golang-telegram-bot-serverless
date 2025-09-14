package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

func init() {
	// загрузить переменные из .env
	godotenv.Load()
}

func Handler(ctx context.Context, event *Event) (*Response, error) {

	// получить апдейт (JSON-строка) из API-Gateway
	rawUpdate := event.Body

	// распарсить апдейт в models.Update
	update := models.Update{}
	json.Unmarshal([]byte(rawUpdate), &update)

	// дефолтныйй роутер
	opts := []bot.Option{
		bot.WithDefaultHandler(handler_echo),
	}

	// получить токен
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("BOT_TOKEN is not set")
	}

	// создать бота
	b, err := bot.New(token, opts...)
	if nil != err {
		panic(err)
	}

	// передать апдейт в бота
	b.ProcessUpdate(ctx, &update)

	return &Response{
		StatusCode: 200,
		Body:       "ok",
	}, nil
}

// эхо роутер
func handler_echo(ctx context.Context, b *bot.Bot, update *models.Update) {
	// фильтр: ответ на не текстовое сообщение
	if update.Message.Text == "" {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "I only respond to text messages",
		})
	}
	// эхо
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
