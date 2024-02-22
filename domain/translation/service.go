package translation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Service handles the business logic of the  module
type Service struct {
	// Add any dependencies here
	repo *Repository
}

type TranslationRequest struct {
	Text        []string `json:"text"`
	Formality   string   `json:"formality"`
	Source_lang string   `json:"source_lang"`
	Target_lang string   `json:"target_lang"`
}

// NewService creates a new instance of TestService
func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// GetMessage returns a greeting message
func (s *Service) GetMessage() Model {
	return s.repo.GetMessage()
}

func (s *Service) GetTranslation(text string) (string, error) {

	txtArray := []string{text}

	const translateApi string = "https://api-free.deepl.com/v2/translate"
	body := TranslationRequest{
		Text:        txtArray,
		Formality:   "prefer_more",
		Source_lang: "EN",
		Target_lang: "JA",
	}
	bodyBytes, err := json.Marshal(&body)

	if err != nil {
		return "", err
	}
	reader := bytes.NewReader(bodyBytes)

	client := http.Client{}

	req, err := http.NewRequest("POST", translateApi, reader)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Authorization", "DeepL-Auth-Key 76e27b5e-e187-4b1b-5ab9-b8f7d13f6aaf:fx")
	req.Header.Add("Content-Type", "application/json")
	
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	respp, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response from %s: %s", translateApi, err)
		return "", err
	}

	return string(respp), err
}
