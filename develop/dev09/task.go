package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wget <URL>")
		return
	}
	url := os.Args[1]
	downloadSite(url)
}

func downloadSite(url string) {
	dir := sanitizeURL(url)
	os.MkdirAll(dir, os.ModePerm)

	fmt.Printf("Downloading %s...\n", url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading %s: %s\n", url, err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	htmlFile := filepath.Join(dir, "index.html")
	err = os.WriteFile(htmlFile, body, 0644)
	if err != nil {
		fmt.Printf("Error writing to file %s: %s\n", htmlFile, err)
		return
	}

	parseLinks(url, body)
}

func parseLinks(baseURL string, body []byte) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		fmt.Printf("Error parsing document: %s\n", err)
		return
	}

	doc.Find("a").Each(func(index int, item *goquery.Selection) {
		href, exists := item.Attr("href")
		if exists {
			fullURL := resolveURL(baseURL, href)
			if fullURL != "" {
				downloadSite(fullURL)
			}
		}
	})
}

func resolveURL(base, href string) string {
	if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") {
		return href
	}

	if strings.HasPrefix(href, "/") {
		baseURLParts := strings.Split(base, "/")
		return fmt.Sprintf("%s//%s%s", baseURLParts[0], baseURLParts[2], href)
	}
	return ""
}

func sanitizeURL(url string) string {
	sanitized := strings.ReplaceAll(url, "http://", "")
	sanitized = strings.ReplaceAll(sanitized, "https://", "")
	sanitized = strings.ReplaceAll(sanitized, "/", "_")
	return sanitized
}
