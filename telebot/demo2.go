package main

import (
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"

	baselog "log"

	"github.com/LeKovr/go-base/log"
	"github.com/comail/colog"
)

// LogConfig defines logger flags
type LogConfig struct {
	Level string `long:"log_level"   default:"info"           description:"Log level [warn|info|debug]"`
}

// NewLog creates new logger
func NewLog(cfg LogConfig) (log.Logger, error) {

	lvl, err := colog.ParseLevel(cfg.Level)
	if err != nil {
		return nil, err
	}

	cl := colog.NewCoLog(os.Stderr, "", baselog.Lshortfile|baselog.Ldate|baselog.Ltime)
	cl.SetMinLevel(lvl)
	cl.SetDefaultLevel(lvl)
	lg := cl.NewLogger()
	return lg, nil
}

func main() {
	log, err := NewLog(LogConfig{"debug"})
	if err != nil {
		baselog.Fatal(err)
	}

	log.Printf("debug: Connecting to Telegram...")
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/hello", func(m *tb.Message) {
		log.Printf("%+v\n", m)
		b.Send(m.Sender, "hello back to the sender")
		b.Send(m.Chat, "hello back to the group")
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		// all the text messages that weren't
		// captured by existing handlers
		log.Printf("debug: Message: %+v", m)
		log.Printf("debug: Sender: %+v", m.Sender)
		log.Printf("debug: %s: %s", m.Chat.Title, m.Text)
	})

	b.Handle(tb.OnPhoto, func(m *tb.Message) {
		// photos only
		log.Printf("debug: Message: %+v", m)
		log.Printf("debug: Sender: %+v", m.Sender)
		log.Printf("debug: %s: %s", m.Chat.Title, m.Text)
	})

	b.Handle(tb.OnChannelPost, func(m *tb.Message) {
		// channel posts only
		log.Printf("debug: Message: %+v", m)
	})

	b.Handle(tb.OnQuery, func(q *tb.Query) {
		// incoming inline queries
		log.Printf("debug: Message: %+v", q)
	})

	b.Start()
}
