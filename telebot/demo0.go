package main

import (
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  "TOKEN_HERE",
		Poller: &tb.LongPoller{10 * time.Second},
	})

	if err != nil {
		return
	}

	b.Handle(tb.OnMessage, func(m *tb.Message) {
		b.Send(m.Sender, "hello world")
	})

	b.Start()
}
