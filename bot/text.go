package bot

import (
	"fmt"
	"regexp"

	"coddabot/database"
	tele "gopkg.in/telebot.v3"
)

const (
	YES = "y"
	NO  = "n"
)

func (b *Bot) onCancel(c tele.Context) error {
	username := c.Sender().Username

	_, containsKey := b.tmpStore[username]
	if !containsKey {
		return c.Send("Non ci sono operazioni da annullare")
	}

	delete(b.tmpStore, username)
	return c.Send("Okay! Ho cancellato la tua richiesta di aggiunta")
}

func (b *Bot) onText(c tele.Context) error {
	var (
		username = c.Sender().Username
		text     = c.Text()
	)

	value, containsKey := b.tmpStore[username]
	if !containsKey {
		return c.Send("Dimmi cosa devo fare")
	}

	if value == nil {
		word, err := parse(text)
		if err != nil {
			return c.Send("Non hai scritto correttamente il testo, gli spazi sono importanti!")
		}

		b.tmpStore[username] = word

		res := fmt.Sprintf("%s: %s\n\nConfermi di voler aggiungere? [y/n]", word.Word, word.Definition)
		return c.Send(res)
	}

	if text != YES && text != NO {
		res := fmt.Sprintf("%s: %s\n\nConfermi di voler aggiungere? [y/n]", value.Word, value.Definition)
		return c.Send(res)
	}

	if text == YES {
		err := b.db.Words.Create(value.Word, value.Definition, username)
		if err != nil {
			return c.Send("Ops! Operazione non riuscita, riprova dopo...")
		}

		delete(b.tmpStore, username)
		return c.Send("Okay! Ho aggiunto tutto!")
	}

	if text == NO {
		delete(b.tmpStore, username)
		return c.Send("Okay! Ho cancellato la tua richiesta di aggiunta")
	}

	return nil
}

func parse(text string) (*database.Word, error) {
	var results = regexp.MustCompile(` - `).Split(text, 2)

	if len(results) != 2 {
		return nil, fmt.Errorf("incorrect length")
	}

	if results[0] == "" || results[1] == "" {
		return nil, fmt.Errorf("empty field are not accepted")
	}

	return &database.Word{
		Word:       results[0],
		Definition: results[1],
	}, nil
}
