package bot

import (
	"log"

	tele "gopkg.in/telebot.v3"
)

func (b *Bot) onStart(c tele.Context) error {
	chat := c.Sender()

	log.Printf("user: %s started bot", chat.Username)

	return c.Send(
		b.Text(c, "start"),
		b.Markup(c, "start"),
		tele.NoPreview,
	)
}
