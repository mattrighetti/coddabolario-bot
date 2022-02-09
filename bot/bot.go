package bot

import (
	"coddabot"
	"coddabot/database"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
	"gopkg.in/telebot.v3/middleware"
)

type Bot struct {
	*tele.Bot
	*layout.Layout
	db *database.DB
}

func New(path string, boot coddabot.Bootstrap) (*Bot, error) {
	lt, err := layout.New(path)
	if err != nil {
		return nil, err
	}

	b, err := tele.NewBot(lt.Settings())
	if err != nil {
		return nil, err
	}

	return &Bot{
		Bot:    b,
		Layout: lt,
		db:     boot.DB,
	}, nil
}

func (b *Bot) Start() {
	b.Use(middleware.Logger())
	b.Use(b.Middleware("it"))

	b.Handle("/start", b.onStart)
	b.Handle(tele.OnQuery, b.onQuery)

	b.Bot.Start()
}

func (b *Bot) selectWords(substring string) ([]database.Word, error) {
	return b.db.Words.GetWords(substring)
}
