package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/nekrophantom/fiber-url_shortener/models"
)

func UrlShorten(ctx *fiber.Ctx) error {
	var inputURL struct {
		URL string `json:"url"`
	}

	if err := ctx.BodyParser(&inputURL); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Call Bitly API to shorten the URL
	shortenedURL, err := shortenWithBitly(inputURL.URL)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to shorten URL"})
	}

	response := models.URLResponse{
		OriginalURL:  inputURL.URL,
		ShortenedURL: shortenedURL,
	}

	return ctx.JSON(response)
}

func shortenWithBitly(originalURL string) (string, error) {
	
	bitlyAPI := os.Getenv("BITLY_API")
	bitlyToken := os.Getenv("ACCESS_TOKEN")

	reqBody, err := json.Marshal(map[string]string{"long_url": originalURL})
	
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", bitlyAPI, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+bitlyToken)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	var responseBody map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return "", err
	}

	shortURL, ok := responseBody["id"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get short URL from Bitly API")
	}

	return shortURL, nil
}