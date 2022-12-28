package client

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

func GetHtmlTitle(url string) (string, error) {
	// Note: requires public URL
	req, err := http.NewRequest(http.MethodGet, url, nil)
	
	if err != nil {
		return "", errors.New("Failed to setup request client.")
	}
	
	resp, err := Client.Do(req)
	
	if err != nil {
		return "", errors.New("Failed to Get URL.")
	}
	
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("HTTP Get status %s", resp.Status))
	}

	if title, err := getHtmlTitle(resp.Body); err == nil {
		return title, nil
	}

	return "", errors.New("Failed to get HTML title")
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}

func getHtmlTitle(r io.Reader) (string, error) {
	doc, err := html.Parse(r)

	if err != nil {
		return "", errors.New("Failed to parse html")
	}

	if title, ok := traverse(doc); ok {
		return title, nil
	}

	return "", errors.New("Failed to find title.")
}