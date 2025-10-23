package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// Repository represents a GitCode repository
type Repository struct {
	ID              int64    `json:"id"`
	Name            string   `json:"name"`
	FullName        string   `json:"full_name"`
	Description     string   `json:"description"`
	Private         bool     `json:"private"`
	Fork            bool     `json:"fork"`
	HTMLURL         string   `json:"html_url"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
	PushedAt        string   `json:"pushed_at"`
	StargazersCount int      `json:"stargazers_count"`
	WatchersCount   int      `json:"watchers_count"`
	ForksCount      int      `json:"forks_count"`
	Language        *string  `json:"language"`
	Archived        bool     `json:"archived"`
	Disabled        bool     `json:"disabled"`
	OpenIssuesCount int      `json:"open_issues_count"`
	License         *License `json:"license"`
	Topics          []string `json:"topics"`
	DefaultBranch   string   `json:"default_branch"`
}

// License represents repository license information
type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SPDXID string `json:"spdx_id"`
	URL    string `json:"url"`
}

// SearchResponse represents the API response structure
type SearchResponse struct {
	TotalCount        int          `json:"total_count"`
	IncompleteResults bool         `json:"incomplete_results"`
	Items             []Repository `json:"items"`
}

// SearchResult contains all collected repositories and metadata
type SearchResult struct {
	Query      string       `json:"query"`
	TotalCount int          `json:"total_count"`
	Items      []Repository `json:"items"`
}

// GitCodeClient handles API interactions with GitCode
type GitCodeClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewGitCodeClient creates a new GitCode client with default settings
func NewGitCodeClient() *GitCodeClient {
	return &GitCodeClient{
		BaseURL: "https://gitcode.com/api/v5",
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// SearchRepositories searches for repositories matching the query
// Returns all repositories across the specified number of pages
func (c *GitCodeClient) SearchRepositories(ctx context.Context, query string, maxPages int) (*SearchResult, error) {
	if query == "" {
		return nil, errors.New("query cannot be empty")
	}
	if maxPages <= 0 {
		return nil, errors.New("maxPages must be greater than 0")
	}

	var allRepos []Repository
	totalCount := 0

	for page := 1; page <= maxPages; page++ {
		resp, err := c.searchPage(ctx, query, page, 50) // max per_page=50
		if err != nil {
			// If we get an error on the first page, return it
			if page == 1 {
				return nil, fmt.Errorf("failed to fetch page %d: %w", page, err)
			}
			// For subsequent pages, log the error but continue with what we have
			log.Printf("Warning: failed to fetch page %d: %v", page, err)
			break
		}

		// Update total count from the first page response
		if page == 1 {
			totalCount = resp.TotalCount
		}

		// If no items returned, we've reached the end
		if len(resp.Items) == 0 {
			break
		}

		allRepos = append(allRepos, resp.Items...)

		// Respect rate limiting by adding a small delay between requests
		if page < maxPages {
			time.Sleep(100 * time.Millisecond)
		}
	}

	return &SearchResult{
		Query:      query,
		TotalCount: totalCount,
		Items:      allRepos,
	}, nil
}

// searchPage fetches a single page of search results
func (c *GitCodeClient) searchPage(ctx context.Context, query string, page, perPage int) (*SearchResponse, error) {
	// Build the request URL
	u, err := url.Parse(c.BaseURL + "/search/repositories")
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	q := u.Query()
	q.Set("q", query)
	q.Set("page", strconv.Itoa(page))
	q.Set("per_page", strconv.Itoa(perPage))
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set appropriate headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "gitcode-search-client/1.0")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var searchResp SearchResponse
	if err := json.Unmarshal(body, &searchResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &searchResp, nil
}

// RepositorySummary contains key information about a repository
type RepositorySummary struct {
	Name            string   `json:"name"`
	FullName        string   `json:"full_name"`
	Description     string   `json:"description"`
	URL             string   `json:"url"`
	Stars           int      `json:"stars"`
	Forks           int      `json:"forks"`
	Language        string   `json:"language"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
	IsPrivate       bool     `json:"is_private"`
	IsFork          bool     `json:"is_fork"`
	IsArchived      bool     `json:"is_archived"`
	Topics          []string `json:"topics"`
	License         string   `json:"license"`
	OpenIssuesCount int      `json:"open_issues_count"`
}

// ExtractSummary extracts key information from a repository
func ExtractSummary(repo Repository) RepositorySummary {
	language := "Unknown"
	if repo.Language != nil && *repo.Language != "" {
		language = *repo.Language
	}

	license := "None"
	if repo.License != nil && repo.License.Name != "" {
		license = repo.License.Name
	}

	// Clean up description
	description := strings.TrimSpace(repo.Description)
	if description == "" {
		description = "No description provided"
	}

	return RepositorySummary{
		Name:            repo.Name,
		FullName:        repo.FullName,
		Description:     description,
		URL:             repo.HTMLURL,
		Stars:           repo.StargazersCount,
		Forks:           repo.ForksCount,
		Language:        language,
		CreatedAt:       repo.CreatedAt,
		UpdatedAt:       repo.UpdatedAt,
		IsPrivate:       repo.Private,
		IsFork:          repo.Fork,
		IsArchived:      repo.Archived,
		Topics:          repo.Topics,
		License:         license,
		OpenIssuesCount: repo.OpenIssuesCount,
	}
}

// PrintSummary prints repository summaries in a readable format
func PrintSummary(summaries []RepositorySummary) {
	if len(summaries) == 0 {
		fmt.Println("No repositories found.")
		return
	}

	fmt.Printf("Found %d repositories:\n\n", len(summaries))
	for i, summary := range summaries {
		fmt.Printf("%d. %s\n", i+1, summary.FullName)
		fmt.Printf("   URL: %s\n", summary.URL)
		fmt.Printf("   Description: %s\n", summary.Description)
		fmt.Printf("   Language: %s | Stars: %d | Forks: %d\n",
			summary.Language, summary.Stars, summary.Forks)
		fmt.Printf("   Created: %s | Updated: %s\n", summary.CreatedAt, summary.UpdatedAt)
		fmt.Printf("   Private: %t | Fork: %t | Archived: %t\n",
			summary.IsPrivate, summary.IsFork, summary.IsArchived)
		if len(summary.Topics) > 0 {
			fmt.Printf("   Topics: %s\n", strings.Join(summary.Topics, ", "))
		}
		fmt.Printf("   License: %s | Open Issues: %d\n", summary.License, summary.OpenIssuesCount)
		fmt.Println(strings.Repeat("-", 50))
	}
}

func main() {
	// Parse command line arguments
	if len(os.Args) < 2 {
		log.Fatal("Usage: gitcode-search <search_query> [max_pages]")
	}

	query := os.Args[1]
	maxPages := 5 // default
	if len(os.Args) > 2 {
		var err error
		maxPages, err = strconv.Atoi(os.Args[2])
		if err != nil || maxPages <= 0 {
			log.Fatal("max_pages must be a positive integer")
		}
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// Initialize client and perform search
	client := NewGitCodeClient()
	result, err := client.SearchRepositories(ctx, query, maxPages)
	if err != nil {
		log.Fatalf("Search failed: %v", err)
	}

	// Output full JSON to stdout (can be redirected to file)
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Println(string(jsonData))

	// Extract and print key information to stderr for better UX
	fmt.Fprintln(os.Stderr, "\n=== KEY REPOSITORY INFORMATION ===")
	summaries := make([]RepositorySummary, len(result.Items))
	for i, repo := range result.Items {
		summaries[i] = ExtractSummary(repo)
	}
	PrintSummary(summaries)

	// Print summary statistics to stderr
	fmt.Fprintf(os.Stderr, "\nSearch completed:\n")
	fmt.Fprintf(os.Stderr, "- Query: %q\n", query)
	fmt.Fprintf(os.Stderr, "- Total repositories found: %d\n", result.TotalCount)
	fmt.Fprintf(os.Stderr, "- Repositories retrieved: %d\n", len(result.Items))
	fmt.Fprintf(os.Stderr, "- Pages processed: %d\n", min(maxPages, (len(result.Items)+99)/100))
}

// Helper function for Go versions < 1.21
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
