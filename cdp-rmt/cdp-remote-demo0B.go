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

	url := "https://example.com"
	var htmlContent string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
	// Wait until the body element is visible (ensures page is loaded)
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		// Retrieve the OuterHTML of the entire 'html' element, saving it to htmlContent
		chromedp.OuterHTML(`html`, &htmlContent, chromedp.ByQuery),
	); err != nil {
		log.Fatalf("Failed getting html source: %v", err)
	}

	log.Println(htmlContent)
}
