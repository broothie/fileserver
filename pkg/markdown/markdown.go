package markdown

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/shurcooL/github_flavored_markdown"
)

var client = &http.Client{
	Timeout: 2 * time.Second,
}

type Payload struct {
	Text    string `json:"text"`
	Mode    string `json:"mode"`
	Context string `json:"context,omitempty"`
}

func generateMarkup(markdown []byte) []byte {
	// Start local rendering
	localChan := make(chan []byte)
	go func() { localChan <- github_flavored_markdown.Markdown(markdown) }()

	// Build request
	body, err := ffjson.Marshal(&Payload{
		Text: string(markdown),
		Mode: "markdown",
	})
	request, err := http.NewRequest(http.MethodPost, "https://api.github.com/markdown", bytes.NewReader(body))
	if err != nil {
		return <-localChan
	}

	// Make request
	response, err := client.Do(request)
	if err != nil {
		return <-localChan
	}

	// Read request body
	markup, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return <-localChan
	}

	return markup
}
