package bot

import (
	tele "gopkg.in/telebot.v3"
)

func (b Bot) onQuery(c tele.Context) error {
	data := c.Data()
	if data == "" {
		data = "empty text"
	}

	results, err := b.selectWords(data)
	if err != nil {
		return err
	}

	var qResults = tele.Results{}
	for _, result := range results {
		qResults = append(qResults, &tele.ArticleResult{
			ResultBase:  tele.ResultBase{ID: string(result.ID)},
			Title:       result.Word,
			Text:        result.Word + ": " + result.Definition,
			HideURL:     true,
			Description: result.Word + ": " + result.Definition,
		})
	}

	return c.Answer(&tele.QueryResponse{
		Results:   qResults,
		CacheTime: 1000,
	})
}
