# Web Scraping Tool using Go and FastAPI

This project is a web scraping tool developed in Go using the FastAPI framework. It scrapes product information from a specified website and stores the data in a local JSON database. It also supports proxy usage and limiting the number of pages to scrape.

## Project Structure

```plaintext
.
├── cmd
│   └── http
│       ├── __debug_bin2414403726
│       ├── database.json
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── cache
│   │   ├── cache.go
│   │   ├── redis.go
│   │   └── scrape.go
│   ├── clients
│   │   └── dentalstall
│   │       └── client.go
│   ├── domain
│   │   ├── entities
│   │   │   └── product.go
│   │   ├── request
│   │   │   └── scrape.go
│   │   └── response
│   │       └── scrape.go
│   ├── handler
│   │   ├── handler.go
│   │   └── scraper.go
│   ├── repo
│   │   ├── repo.go
│   │   └── scrape.go
│   └── service
│       ├── notify.go
│       ├── scraper.go
│       └── service.go
└── middleware
    └── authenticate.go
```


## Getting Started
### Prerequisites
  Go 1.16 or higher
  Redis (for caching)
## Installation
  ### Clone the repository:
    1. git clone https://github.com/yourusername/yourrepository.git
    2. cd yourrepository
### Install dependencies:
    1. go mod tidy
    2. Start Redis server:
      1. brew services start redis
### Running the Application
    go run cmd/http/main.go

### Make a scrape request:
Use curl to hit the scrape endpoint.
curl -X POST http://localhost:3000/scrape -H "Authorization: Bearer YOUR_STATIC_TOKEN" -d '{"pages": 5, "proxy": "http://your-proxy.com"}'
Configuration
Scraping Settings:
  Limit the number of pages to scrape by specifying the pages field in the request.
  Use a proxy by providing the proxy field in the request.
Authentication:
  The API uses a static token for authentication. Include the token in the Authorization header as Bearer YOUR_STATIC_TOKEN.
### Logging
The application logs scraping status to the console. Modify the logging mechanism in internal/service/notify.go if different notification methods are required.

### Caching
Caching is implemented to avoid redundant scraping of unchanged products. Modify caching logic in internal/cache as needed.


