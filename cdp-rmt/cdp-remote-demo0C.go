package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

var pWaitElm = flag.String("w", "a", "Page element to wait for")

func main() {
	lpdToken := os.Getenv("LPD_TOKEN")
	if lpdToken == "" {
		log.Fatal("LPD_TOKEN environment variable is not set.")
	}

	flag.Parse()
	url := "https://example.com"
	if len(flag.Args()) >= 1 {
		url = flag.Args()[0]
	}
	fmt.Fprintf(os.Stderr, "Visiting '%s' (& wait for '%s') ...\n", url, *pWaitElm)

	// 1. ⏱️ Set the main timeout on the parent context
	// This context governs the entire operation (connection and tasks).
	parentCtx, parentCancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer parentCancel()

	// 2. Create the remote allocator context
	// Pass the timed-out parentCtx here.
	allocatorCtx, allocatorCancel := chromedp.NewRemoteAllocator(parentCtx,
		"wss://euwest.cloud.lightpanda.io/ws?token="+lpdToken, chromedp.NoModifyURL,
	)
	defer allocatorCancel()

	// 3. Create the standard browser context
	ctx, cancel := chromedp.NewContext(
		allocatorCtx,
		chromedp.WithLogf(log.Printf),   // Log all successful CDP commands
		chromedp.WithErrorf(log.Printf), // Log any errors, like timeout
		chromedp.WithDebugf(log.Printf), // Log detailed internal messages
	)
	defer cancel()

	var htmlContent string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		// Wait until the specific element is visible (ensures page is loaded)
		chromedp.WaitVisible(*pWaitElm, chromedp.ByQuery),
		// Retrieve the OuterHTML of the entire 'html' element, saving it to htmlContent
		chromedp.OuterHTML(`html`, &htmlContent, chromedp.ByQuery),
	); err != nil {
		log.Fatalf("Failed getting html source: %v", err)
	}

	fmt.Println(htmlContent)
}
