package bot

import (
	tele "gopkg.in/telebot.v3"
)

func needUsername() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if username := c.Sender().Username; username == "" {
				return c.Send("Imposta uno username per parlare con il bot")
			}
			return next(c)
		}
	}
}
