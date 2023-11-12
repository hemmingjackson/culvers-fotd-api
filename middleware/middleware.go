package middleware

import (
    "github.com/gocolly/colly"
	"strings"
	"net/http"
	"example/gin-service-web/exceptions"
	"example/gin-service-web/logger"
	"fmt"
	"go.uber.org/zap"
)

func ScrapeWebsite(url string) (*CulversData, error) {
	data := &CulversData{}
	c := colly.NewCollector()

    c.OnHTML(".ModuleRestaurantDetail-fotd h2 strong", func(e *colly.HTMLElement) {
        data.Flavor = strings.TrimSpace(e.Text)
    })

	c.OnHTML(".contain-restaurant-info h1", func(e *colly.HTMLElement) {
        data.Town = strings.TrimSpace(e.Text)
    })

	c.OnError(func(r *colly.Response, err error) {
		if r.StatusCode == http.StatusNotFound {
			err = exceptions.New(404, fmt.Sprintf("not found: %s", url))
			logger.Logger.Error("Scraping error", zap.Error(err)) // Log the error
		}
	})

    err := c.Visit(url)
    if err != nil {
		logger.Logger.Error("Scraping visit error", zap.Error(err))
        return nil, fmt.Errorf("failed to scrape Culver's website for town %s: %w", data.Town, err)
    }

    return data, nil
}
