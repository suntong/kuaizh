package main

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

func main() {
	lpdToken := os.Getenv("LPD_TOKEN")
	if lpdToken == "" {
		log.Fatal("LPD_TOKEN environment variable is not set. Please set it.")
	}

	ctx, cancel := chromedp.NewRemoteAllocator(context.TODO(),
		"wss://euwest.cloud.lightpanda.io/ws?token="+lpdToken, chromedp.NoModifyURL,
	)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var title string
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://lightpanda.io"),
		chromedp.Title(&title),
	); err != nil {
		log.Fatalf("Failed getting title of lightpanda.io: %v", err)
	}

	log.Println("Got title of:", title)
}
