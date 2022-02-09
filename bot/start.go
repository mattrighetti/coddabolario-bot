package bot

import (
	tele "gopkg.in/telebot.v3"
)

func (b *Bot) onStart(c tele.Context) error {
	return c.Send("Eja!")
}
