package word

import (
	service "github.com/Muravjeva/omp-bot/internal/service/english/word"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type WordCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

func NewWordCommander(bot *tgbotapi.BotAPI, service service.WordService) WordCommander {
	// ...
}
