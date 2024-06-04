package request

type ScrapeRequest struct {
	PageLimit int    `json:"page_limit"`
	Proxy     string `json:"proxy"`
}
