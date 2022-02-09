package bot

import (
	tele "gopkg.in/telebot.v3"
)

func (b *Bot) onAdd(c tele.Context) error {
	if c.Chat().Type != tele.ChatPrivate {
		return c.Send("Questa funzionalità funziona solamente nella chat privata")
	}

	username := c.Sender().Username

	_, exists := b.tmpStore[username]
	if !exists {
		b.tmpStore[username] = nil
	}

	return c.Send(
		b.Text(c, "add-format"),
		b.Markup(c, "add-format"),
		tele.NoPreview,
	)
}
