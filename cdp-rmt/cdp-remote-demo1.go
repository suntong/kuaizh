package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/go-rod/rod/lib/launcher" // Import the launcher library
)

func main() {
	// 1. Connection Setup
	lpdToken := os.Getenv("LPD_TOKEN")
	if lpdToken == "" {
		log.Fatal("LPD_TOKEN environment variable is not set. Please set it.")
	}
    
    // Use the correct WSS scheme and port for the remote service.
	wsURL := fmt.Sprintf("wss://euwest.cloud.lightpanda.io:443/ws?token=%s", lpdToken)

	log.Printf("Connecting to Lightpanda Cloud: %s", wsURL)

    // Create a background context with a timeout
	baseCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// 2. Use rod/lib/launcher to parse the WSS URL correctly.
	// This generates the necessary options to connect directly via WebSocket.
	u := launcher.MustParseURL(wsURL)

    // 3. Create the chromedp context using the remote connection options.
    // We pass the URL in the format required by the With
	ctx, cancel := chromedp.NewContext(
		baseCtx,
		chromedp.WithErrorf(log.Printf), // Optional: Direct log errors to standard output
		chromedp.WithDebugf(log.Printf), // Optional: See connection details
		
        // Connect to the remote browser using the WSS URL parsed by launcher.
        chromedp.WithTargets(u.MustResolveWebSocketURL),
	)
	defer cancel()

	// 4. Run the automation task
	var title string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://example.com"),
		chromedp.Title(&title),
	)

	if err != nil {
		log.Fatalf("Error during automation: %v", err)
	}

	// 5. Output the result
	log.Printf("Successfully connected to Lightpanda Cloud.")
	log.Printf("Page Title: %s", title)
}