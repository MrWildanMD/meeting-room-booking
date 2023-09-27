package config

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

var (
	TB *tele.Bot
	ChatID int64 = -4081629174
)

func RegisterBot(cfg *Config) {
	pref := tele.Settings{
		Token: cfg.TelebotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	TB = b
}