package app

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
)

func Scrape(endpoint string) (*goquery.Document, error) {
	var doc *goquery.Document

	// Request the HTML page.
	url := TargetUrl(endpoint)
	res, err := http.Get(url)
	if err != nil {
		return doc, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return doc, errors.New(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
	}

	// Load the HTML document
	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return doc, err
	}

	return doc, err
}

func TargetUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s", os.Getenv("TARGET_URL"), endpoint)
}
