package main

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"pkg.cld19.com/mosbot/core"
	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
}

type mosbotRequest struct {
	DocType string `json:"type"`
	DocId   string `json:"id"`
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	p := &mosbotRequest{DocType: "doc", DocId: "0"}
	json.NewDecoder(in).Decode(p)
	log.Println(p)
	genUrl, err := core.GenerateUrl(p.DocType, p.DocId)
	if err != nil {
		log.Fatal(err)
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