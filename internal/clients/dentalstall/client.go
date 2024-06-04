package dentalstall

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type DentalStallClient struct {
	Url string
}

type IDentalStallClient interface {
	GetPageData(pageNumber int, proxy string) (io.Reader, error)
}

func NewClient(url string) IDentalStallClient {
	return &DentalStallClient{
		Url: url,
	}
}

func (c *DentalStallClient) GetPageData(pageNumber int, proxy string) (io.Reader, error) {
	var u string
	if pageNumber == 1 {
		u = c.Url + "/shop"
	} else {
		u = c.Url + fmt.Sprintf("/shop/page/%d/", pageNumber)
	}
	var err error
	for i := 0; i < 3; i++ {
		client := &http.Client{}
		if proxy != "" {
			proxyURL, _ := url.Parse(proxy)
			client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
		}

		resp, err := client.Get(u)
		if err != nil {
			return nil, fmt.Errorf("couldn't get response %v", err)
		}
		if resp.StatusCode == 200 {
			return resp.Body, nil
		}
	}
	return nil, fmt.Errorf("couldn't get response %v", err)
}
