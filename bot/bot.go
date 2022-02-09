package bot

import (
	"coddabot"
	"coddabot/database"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
	"gopkg.in/telebot.v3/middleware"
)

type tmpStore map[string]*database.Word

type Bot struct {
	*tele.Bot
	*layout.Layout
	db       *database.DB
	tmpStore tmpStore
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
		Bot:      b,
		Layout:   lt,
		db:       boot.DB,
		tmpStore: tmpStore{},
	}, nil
}

func (b *Bot) Start() {
	b.Use(middleware.Logger())

	b.Handle("/start", b.onStart)
	b.Handle("/add", b.onAdd, needUsername())
	b.Handle("/cancel", b.onCancel, needUsername())
	b.Handle(tele.OnText, b.onText, needUsername())
	b.Handle(tele.OnQuery, b.onQuery)

	b.Bot.Start()
}

func (b *Bot) selectWords(substring string) ([]database.Word, error) {
	return b.db.Words.GetWords(substring)
}
