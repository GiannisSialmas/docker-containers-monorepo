package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	EnvSlackWebhook  = "SLACK_WEBHOOK"
	EnvSlackIcon     = "SLACK_ICON"
	EnvSlackChannel  = "SLACK_CHANNEL"
	EnvSlackTitle    = "SLACK_TITLE"
	EnvSlackMessage  = "SLACK_MESSAGE"
	EnvSlackColor    = "SLACK_COLOR"
	EnvSlackUserName = "SLACK_USERNAME"
)

type Webhook struct {
	Text        string       `json:"text,omitempty"`
	UserName    string       `json:"username,omitempty"`
	IconURL     string       `json:"icon_url,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	UnfurlLinks bool         `json:"unfurl_links"`
	Attachments []Attachment `json:"attachments,omitmepty"`
}

type Attachment struct {
	Fallback string  `json:"fallback"`
	Pretext  string  `json:"pretext,omitempty"`
	Color    string  `json:"color,omitempty"`
	Fields   []Field `json:"fields,omitempty"`
}

type Field struct {
	Title string `json:"title"`
	Value string `json:"value,omitempty"`
	Short bool   `json:short"`
}

// Package scope variables
var endpoint string = os.Getenv("SLACK_WEBHOOK")

func main() {
	// Allow confirmation the task handling service is running.
	http.HandleFunc("/", indexHandler)

	log.Printf("Slack webhook is %s", endpoint)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
		log.Printf("Defaulting to port %s", port)
	}

	if endpoint == "" {
		fmt.Fprintln(os.Stderr, "Slack webhook URL is required")
		return
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(response, request)
		return
	}

	bodyJSONEncoded, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("ReadAll: %v", err)
		http.Error(response, "Internal Error", http.StatusInternalServerError)
		return
	}

	type Message struct {
		Id      int64  `json:"id"`
		Message string `json:"message"`
	}

	var body Message
	err = json.Unmarshal(bodyJSONEncoded, &body)
	if err != nil {
		http.Error(response, err.Error(), 400)
		return
	}

	log.Printf(body.Message)

	var text string = body.Message
	if text == "" {
		fmt.Fprintln(os.Stderr, "Message is required")
		http.Error(response, "Message is required", http.StatusBadRequest)
	}

	msg := Webhook{
		UserName: os.Getenv(EnvSlackUserName),
		IconURL:  os.Getenv(EnvSlackIcon),
		Channel:  os.Getenv(EnvSlackChannel),
		Attachments: []Attachment{
			{
				Fallback: envOr(EnvSlackMessage, "This space intentionally left blank"),
				Color:    os.Getenv(EnvSlackColor),
				Fields: []Field{
					{
						Title: os.Getenv(EnvSlackTitle),
						Value: envOr(EnvSlackMessage, text),
					},
				},
			},
		},
	}

	if err := send(endpoint, msg); err != nil {
		fmt.Fprintf(os.Stderr, "Error sending message: %s\n", err)
		os.Exit(2)
	}
}

func envOr(name, def string) string {
	if d, ok := os.LookupEnv(name); ok {
		return d
	}
	return def
}

func send(endpoint string, msg Webhook) error {
	enc, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	b := bytes.NewBuffer(enc)
	res, err := http.Post(endpoint, "application/json", b)
	if err != nil {
		return err
	}

	if res.StatusCode >= 299 {
		return fmt.Errorf("Error on message: %s\n", res.Status)
	}
	fmt.Println(res.Status)
	return nil
}
