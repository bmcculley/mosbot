package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"pkg.cld19.com/mosbot/internal/mocks"
)

func init() {
	Client = &mocks.MockClient{}
}

func TestGetHtmlTitle(t *testing.T) {
	html := `<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>Hello World!</title>
</head>
<body>
	<h1>Just a test page</h1>
</body>
</html>`
	r := ioutil.NopCloser(bytes.NewReader([]byte(html)))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	got, err := GetHtmlTitle("https://support.oracle.com/knowledge/PeopleSoft%20Enterprise/1234567_1.html")
	if err != nil {
		t.Errorf("Got %q want nil", err)
	}

	want := "Hello World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}


func TestGetHtmlTitle2(t *testing.T) {
	html := `<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>Hello World!</title>
</head>
<body>
	<h1>Just a test page</h1>
</body>
</html>`
	r := ioutil.NopCloser(bytes.NewReader([]byte(html)))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 500,
			Body:       r,
		}, nil
	}

	got, err := GetHtmlTitle("https://support.oracle.com/knowledge/PeopleSoft%20Enterprise/1234567_1.html")
	if err == nil {
		t.Errorf("got %q wanted error", err)
	}

	want := ""

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
