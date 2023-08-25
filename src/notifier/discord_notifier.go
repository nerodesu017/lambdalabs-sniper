package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nerodesu017/lambdalabs-sniper/src/constants"
)

type Discord_notifier struct {
	Webhook_url string
}

type Embed struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Color       int          `json:"color"`
	Fields      []EmbedField `json:"fields"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}



func (disc_notifier Discord_notifier) Notify(gpus []*constants.GPU) error {
	description := ""
	
	for _, gpu := range gpus {
		description += gpu.Name + "\n\n"
	}

	embed := Embed{
		Title:       "Available GPUs:",
		Description: description,
		Color:       9498256, // Light green
		Fields: []EmbedField{
			{
				Name:   "LINKS",
				Value:  "[Launch Instance](https://cloud.lambdalabs.com/instances)",
				Inline: true,
			},
		},
	}

	payload := map[string]interface{}{
		"embeds": []Embed{embed},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error when marshalling: %v", err)
	}

	resp, err := http.Post(disc_notifier.Webhook_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("error when posting to webhook: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to send message; status: %v", resp.Status)
	} else {
		log.Printf("Webhook sent!\n")
	}

	return nil
}