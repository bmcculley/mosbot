package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"strings"
	"regexp"

	"pkg.cld19.com/mosbot/core"
	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
}

type mosbotRequest struct {
	Text string `json:"text"`
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	p := &mosbotRequest{Text: "doc 0"}
	json.NewDecoder(in).Decode(p)
	docParts := strings.Fields(p.Text)
	match := false
	if len(docParts) == 2 {
		match, _ = regexp.MatchString("^[0-9]+.[0-9]$", docParts[1])
	}
	genUrl := "Bad Request."
	if match {
		var err error
		genUrl, err = core.GenerateUrl(docParts[0], docParts[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	msg := struct {
		Text string `json:"text"`
		ResponseType string `json:"response_type"`
	}{
		Text: genUrl,
		ResponseType: "in_channel",
	}
	json.NewEncoder(out).Encode(&msg)
}