# Web Scraper Project - Overview

## What is Web Scraping?

Web scraping is the process of extracting data from websites. It involves fetching web pages, parsing HTML content, and extracting specific information.

## Web Scraping in Go

### Basic HTTP Client
```go
package main

import (
    "fmt"
    "io"
    "net/http"
)

func fetchURL(url string) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}
```

## HTML Parsing

### Using goquery
```go
import (
    "github.com/PuerkitoBio/goquery"
    "strings"
)

func parseHTML(html string) (*goquery.Document, error) {
    return goquery.NewDocumentFromReader(strings.NewReader(html))
}

func extractLinks(doc *goquery.Document) []string {
    var links []string
    doc.Find("a").Each(func(i int, s *goquery.Selection) {
        if href, exists := s.Attr("href"); exists {
            links = append(links, href)
        }
    })
    return links
}
```

## Data Extraction Patterns

### Extracting Text Content
```go
func extractText(doc *goquery.Document, selector string) []string {
    var texts []string
    doc.Find(selector).Each(func(i int, s *goquery.Selection) {
        texts = append(texts, strings.TrimSpace(s.Text()))
    })
    return texts
}

func extractTitle(doc *goquery.Document) string {
    return strings.TrimSpace(doc.Find("title").Text())
}
```

### Extracting Attributes
```go
func extractImages(doc *goquery.Document) []string {
    var images []string
    doc.Find("img").Each(func(i int, s *goquery.Selection) {
        if src, exists := s.Attr("src"); exists {
            images = append(images, src)
        }
    })
    return images
}
```

## Structured Data Extraction

### News Article Scraper
```go
type Article struct {
    Title       string `json:"title"`
    Author      string `json:"author"`
    PublishedAt string `json:"published_at"`
    Content     string `json:"content"`
    URL         string `json:"url"`
}

func scrapeArticle(url string) (*Article, error) {
    html, err := fetchURL(url)
    if err != nil {
        return nil, err
    }

    doc, err := parseHTML(html)
    if err != nil {
        return nil, err
    }

    article := &Article{
        URL: url,
    }

    // Extract title
    article.Title = extractTitle(doc)

    // Extract author (example selector)
    authorSel := doc.Find(".author, .byline, [rel='author']")
    if authorSel.Length() > 0 {
        article.Author = strings.TrimSpace(authorSel.First().Text())
    }

    // Extract content
    contentSel := doc.Find("article, .content, .post-content")
    if contentSel.Length() > 0 {
        article.Content = strings.TrimSpace(contentSel.First().Text())
    }

    return article, nil
}
```

## Handling Different Content Types

### JSON API Scraping
```go
import "encoding/json"

type Product struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
}

func scrapeJSONAPI(url string) ([]Product, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var products []Product
    decoder := json.NewDecoder(resp.Body)
    err = decoder.Decode(&products)

    return products, err
}
```

### XML Scraping
```go
import "encoding/xml"

type RSS struct {
    Channel Channel `xml:"channel"`
}

type Channel struct {
    Title       string   `xml:"title"`
    Description string   `xml:"description"`
    Items       []Item   `xml:"item"`
}

type Item struct {
    Title       string `xml:"title"`
    Description string `xml:"description"`
    Link        string `xml:"link"`
    PubDate     string `xml:"pubDate"`
}

func scrapeRSS(url string) (*RSS, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var rss RSS
    decoder := xml.NewDecoder(resp.Body)
    err = decoder.Decode(&rss)

    return &rss, err
}
```

## Rate Limiting and Politeness

### Rate Limiter
```go
import "time"

type RateLimiter struct {
    interval time.Duration
    last     time.Time
}

func NewRateLimiter(requestsPerSecond int) *RateLimiter {
    return &RateLimiter{
        interval: time.Second / time.Duration(requestsPerSecond),
    }
}

func (rl *RateLimiter) Wait() {
    elapsed := time.Since(rl.last)
    if elapsed < rl.interval {
        time.Sleep(rl.interval - elapsed)
    }
    rl.last = time.Now()
}
```

### Polite Scraper
```go
func politeScrape(urls []string, rateLimiter *RateLimiter) []string {
    var results []string

    for _, url := range urls {
        rateLimiter.Wait()

        html, err := fetchURL(url)
        if err != nil {
            log.Printf("Error fetching %s: %v", url, err)
            continue
        }

        results = append(results, html)
    }

    return results
}
```

## User Agent and Headers

### Custom HTTP Client
```go
func createClient() *http.Client {
    return &http.Client{
        Timeout: 30 * time.Second,
    }
}

func fetchWithHeaders(url string) (string, error) {
    client := createClient()

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return "", err
    }

    // Set user agent
    req.Header.Set("User-Agent", "MyBot/1.0 (contact@example.com)")

    // Set other headers
    req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
    req.Header.Set("Accept-Language", "en-US,en;q=0.5")
    req.Header.Set("Accept-Encoding", "gzip, deflate")
    req.Header.Set("Connection", "keep-alive")

    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}
```

## Data Storage

### Saving to JSON
```go
func saveToJSON(data interface{}, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    return encoder.Encode(data)
}
```

### Saving to CSV
```go
func saveToCSV(data [][]string, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    return writer.WriteAll(data)
}
```

## Error Handling and Retries

### Retry Mechanism
```go
func fetchWithRetry(url string, maxRetries int) (string, error) {
    var lastErr error

    for i := 0; i < maxRetries; i++ {
        html, err := fetchURL(url)
        if err == nil {
            return html, nil
        }

        lastErr = err
        log.Printf("Attempt %d failed: %v", i+1, err)

        // Wait before retry
        time.Sleep(time.Duration(i+1) * time.Second)
    }

    return "", fmt.Errorf("failed after %d attempts: %v", maxRetries, lastErr)
}
```

### Status Code Handling
```go
func fetchWithStatusCheck(url string) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}
```

## Concurrent Scraping

### Worker Pool
```go
type Scraper struct {
    client      *http.Client
    rateLimiter *RateLimiter
    workers     int
}

func NewScraper(workers int) *Scraper {
    return &Scraper{
        client:      createClient(),
        rateLimiter: NewRateLimiter(2), // 2 requests per second
        workers:     workers,
    }
}

func (s *Scraper) ScrapeConcurrently(urls []string) []string {
    jobs := make(chan string, len(urls))
    results := make(chan string, len(urls))

    // Start workers
    for i := 0; i < s.workers; i++ {
        go s.worker(jobs, results)
    }

    // Send jobs
    for _, url := range urls {
        jobs <- url
    }
    close(jobs)

    // Collect results
    var allResults []string
    for i := 0; i < len(urls); i++ {
        result := <-results
        if result != "" {
            allResults = append(allResults, result)
        }
    }

    return allResults
}

func (s *Scraper) worker(jobs <-chan string, results chan<- string) {
    for url := range jobs {
        s.rateLimiter.Wait()

        html, err := fetchURL(url)
        if err != nil {
            log.Printf("Error fetching %s: %v", url, err)
            results <- ""
            continue
        }

        results <- html
    }
}
```

## Example Complete Scraper

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
    "time"

    "github.com/PuerkitoBio/goquery"
)

type NewsItem struct {
    Title   string `json:"title"`
    URL     string `json:"url"`
    Summary string `json:"summary"`
    Date    string `json:"date"`
}

type NewsScraper struct {
    client *http.Client
}

func NewNewsScraper() *NewsScraper {
    return &NewsScraper{
        client: &http.Client{
            Timeout: 30 * time.Second,
        },
    }
}

func (ns *NewsScraper) ScrapeNews(url string) ([]NewsItem, error) {
    resp, err := ns.client.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return nil, err
    }

    var news []NewsItem

    // Example selectors (adjust based on actual website)
    doc.Find("article, .news-item, .post").Each(func(i int, s *goquery.Selection) {
        title := strings.TrimSpace(s.Find("h1, h2, h3, .title").Text())
        link, _ := s.Find("a").Attr("href")
        summary := strings.TrimSpace(s.Find(".summary, .excerpt").Text())
        date := strings.TrimSpace(s.Find(".date, time").Text())

        if title != "" {
            news = append(news, NewsItem{
                Title:   title,
                URL:     link,
                Summary: summary,
                Date:    date,
            })
        }
    })

    return news, nil
}

func main() {
    scraper := NewNewsScraper()

    news, err := scraper.ScrapeNews("https://example-news-site.com")
    if err != nil {
        log.Fatal(err)
    }

    // Save to JSON
    data, _ := json.MarshalIndent(news, "", "  ")
    os.WriteFile("news.json", data, 0644)

    fmt.Printf("Scraped %d news items\n", len(news))
}
```

## Best Practices

### 1. Respect robots.txt
```go
func checkRobotsTxt(baseURL string) bool {
    robotsURL := baseURL + "/robots.txt"
    resp, err := http.Get(robotsURL)
    if err != nil {
        return true // Assume allowed if can't check
    }
    defer resp.Body.Close()

    // Parse robots.txt content
    // Check if your user agent is allowed
    return true
}
```

### 2. Use Appropriate Delays
```go
func politeDelay() {
    time.Sleep(1 * time.Second) // Wait 1 second between requests
}
```

### 3. Handle Errors Gracefully
```go
func safeScrape(url string) (string, error) {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Panic recovered: %v", r)
        }
    }()

    return fetchURL(url)
}
```

### 4. Log Your Activities
```go
func logScraping(url string, success bool) {
    log.Printf("Scraping %s: %s", url, map[bool]string{true: "SUCCESS", false: "FAILED"}[success])
}
```

## Next Steps

Now that you understand web scraping, try the tasks in `tasks.md`!