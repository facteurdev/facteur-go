package facteur

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Facteur struct {
	APIKey string
}

type SendEmailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	HTML    string   `json:"html"`
	Text    string   `json:"text"`
}

func NewFacteur(apiKey string) *Facteur {
	return &Facteur{
		APIKey: apiKey,
	}
}

func (cli *Facteur) SendEmail(payload *SendEmailPayload) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://facteur.dev/api/v1/emails", &buf)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + cli.APIKey)

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode == 500 {
		return errors.New("Internal server error")
	}

	if resp.StatusCode != 200 {
		var body map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&body)
		if err != nil {
			return err
		}

		message := body["error"].(string)

		return errors.New(message)
	}


	return nil
}
