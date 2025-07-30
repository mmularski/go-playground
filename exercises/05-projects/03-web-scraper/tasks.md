# Web Scraper Project - Tasks

## Task 1: Basic HTTP Client

Create a basic HTTP client that can fetch web pages.

**Requirements:**
- Implement a function to fetch HTML content from URLs
- Handle HTTP errors and timeouts
- Set appropriate user agent headers
- Implement proper error handling

## Task 2: HTML Parsing

Implement HTML parsing functionality using goquery.

**Requirements:**
- Parse HTML content into a structured format
- Extract text content from specific selectors
- Extract links and images from pages
- Handle parsing errors gracefully

## Task 3: Data Extraction

Implement data extraction from web pages.

**Requirements:**
- Extract titles, headings, and text content
- Extract links and their attributes
- Extract images and their sources
- Create structured data from extracted content

## Task 4: News Scraper

Build a news article scraper.

**Requirements:**
- Define a NewsItem struct with title, URL, summary, date
- Extract news articles from a news website
- Handle different article layouts
- Save extracted data to JSON format

## Task 5: Rate Limiting

Implement polite scraping with rate limiting.

**Requirements:**
- Create a rate limiter to control request frequency
- Add delays between requests
- Respect website's robots.txt
- Implement exponential backoff for failed requests

## Task 6: Error Handling and Retries

Implement robust error handling and retry mechanisms.

**Requirements:**
- Handle network timeouts and connection errors
- Implement retry logic with exponential backoff
- Handle HTTP status codes appropriately
- Log errors and retry attempts

## Task 7: Concurrent Scraping

Implement concurrent scraping with worker pools.

**Requirements:**
- Create a worker pool for concurrent requests
- Implement job distribution and result collection
- Handle rate limiting in concurrent environment
- Manage goroutine lifecycle properly

## Task 8: Data Storage

Implement data storage functionality.

**Requirements:**
- Save scraped data to JSON files
- Save data to CSV format
- Implement data deduplication
- Handle file I/O errors

## Task 9: Configuration Management

Implement configuration for the scraper.

**Requirements:**
- Load configuration from JSON files
- Support command line flags for configuration
- Implement default values for missing config
- Validate configuration parameters

## Task 10: Advanced Features

Implement advanced scraping features.

**Requirements:**
- Support for different content types (JSON, XML, RSS)
- Implement pagination handling
- Add search functionality
- Implement data filtering and sorting

## Bonus Challenges

### Challenge 1: E-commerce Scraper
- Scrape product information from e-commerce sites
- Extract prices, reviews, and product details
- Handle dynamic content loading
- Implement price tracking over time

### Challenge 2: Social Media Scraper
- Scrape social media posts and comments
- Handle authentication and session management
- Implement real-time data collection
- Respect platform rate limits

### Challenge 3: Image Scraper
- Download images from websites
- Implement image processing and resizing
- Handle different image formats
- Create image galleries and collections

### Challenge 4: API Integration
- Integrate with web APIs
- Handle API authentication
- Implement API rate limiting
- Create API wrapper libraries

## Testing Your Scraper

### Basic Usage Examples
```bash
# Scrape a single page
./scraper fetch https://example.com

# Scrape multiple pages
./scraper batch urls.txt

# Scrape with rate limiting
./scraper polite --delay 2s https://example.com

# Save to different formats
./scraper fetch --output json https://example.com
./scraper fetch --output csv https://example.com
```

### Testing Different Scenarios
```bash
# Test error handling
./scraper fetch https://invalid-url.com
./scraper fetch https://httpstat.us/404

# Test rate limiting
./scraper batch --workers 5 urls.txt

# Test data extraction
./scraper extract --selector "h1" https://example.com
./scraper extract --selector "img" https://example.com
```

## Success Criteria

Your web scraper should:
- ✅ Fetch web pages reliably with proper error handling
- ✅ Parse HTML content and extract structured data
- ✅ Implement rate limiting and polite scraping
- ✅ Handle concurrent requests efficiently
- ✅ Save data in multiple formats (JSON, CSV)
- ✅ Include comprehensive error handling and retries
- ✅ Support configuration management
- ✅ Handle different content types (HTML, JSON, XML)
- ✅ Include proper logging and monitoring
- ✅ Be well-documented with usage examples

## Project Structure

```
web-scraper/
├── main.go              # Entry point
├── scraper/
│   ├── scraper.go      # Main scraper logic
│   ├── client.go       # HTTP client
│   └── parser.go       # HTML parser
├── extractors/
│   ├── news.go         # News extractor
│   ├── products.go     # Product extractor
│   └── images.go       # Image extractor
├── storage/
│   ├── json.go         # JSON storage
│   └── csv.go          # CSV storage
├── config/
│   └── config.go       # Configuration
└── utils/
    ├── rate_limiter.go # Rate limiting
    └── retry.go        # Retry logic
```

## Dependencies

Add these to your go.mod:
```
require (
    github.com/PuerkitoBio/goquery v1.8.1
)
```

## Next Steps

Start with Task 1 and work through each task systematically. Each task builds upon the previous ones, so make sure to complete them in order.