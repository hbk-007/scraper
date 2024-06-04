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


## Directory Breakdown
cmd/http: Entry point of the application.

main.go: Contains the main function to start the HTTP server.
database.json: Local JSON database to store scraped data.
internal/cache: Caching mechanism to prevent redundant scraping.

cache.go: Defines the cache interface.
redis.go: Redis implementation for caching.
scrape.go: Caching logic specific to scraping.
internal/clients/dentalstall: Client to interact with the target website.

client.go: Implements the HTTP client for scraping the website.
internal/domain: Domain entities and request/response structures.

entities: Defines the product entity.
product.go: Product struct definition.
request: Structures for handling scrape requests.
scrape.go: Scrape request struct.
response: Structures for handling scrape responses.
scrape.go: Scrape response struct.
internal/handler: HTTP request handlers.

handler.go: General HTTP handlers.
scraper.go: Handler specific to scraping requests.
internal/repo: Data repository for storing scraped data.

repo.go: Defines the repository interface.
scrape.go: Implementation of the repository for scraping data.
internal/service: Business logic and services.

notify.go: Service for sending notifications.
scraper.go: Scraping service.
service.go: General service definitions.
middleware: Middleware for HTTP request processing.

authenticate.go: Static token-based authentication middleware.
Getting Started
Prerequisites
Go 1.16 or higher
Redis (for caching)
Installation
Clone the repository:

sh
Copy code
git clone https://github.com/yourusername/yourrepository.git
cd yourrepository
Install dependencies:

sh
Copy code
go mod tidy
Start Redis server (if using Redis for caching):

sh
Copy code
brew services start redis
Running the Application
Run the application:

sh
Copy code
go run cmd/http/main.go
Make a scrape request:

Use curl to hit the scrape endpoint.

sh
Copy code
curl -X POST http://localhost:3000/scrape -H "Authorization: Bearer YOUR_STATIC_TOKEN" -d '{"pages": 5, "proxy": "http://your-proxy.com"}'
Configuration
Scraping Settings:

Limit the number of pages to scrape by specifying the pages field in the request.
Use a proxy by providing the proxy field in the request.
Authentication:

The API uses a static token for authentication. Include the token in the Authorization header as Bearer YOUR_STATIC_TOKEN.
Logging
The application logs scraping status to the console. Modify the logging mechanism in internal/service/notify.go if different notification methods are required.

Caching
Caching is implemented to avoid redundant scraping of unchanged products. Modify caching logic in internal/cache as needed.

Future Improvements
Implement more sophisticated data storage solutions (e.g., SQL or NoSQL databases).
Enhance error handling and retry mechanisms.
Implement more robust notification systems (e.g., email or messaging services).
Add more scraping settings and configurations.