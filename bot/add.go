package bot

import (
	"fmt"

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

	return c.Send(fmt.Sprintf("Invia la definizione in questo formato:\n\nparola - definizione"))
}
