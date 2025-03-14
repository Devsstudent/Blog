package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"blog.local/interfaces/types"
	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)
func validatePosting(c *fiber.Ctx) error {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:9000",
		Password: "",
		DB:       0,
	})
	
	// Get the authorization token from the request query parameters
	authToken := c.Query("auth")
	
	// Load environment variables
	envFile, err := godotenv.Read(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	
	envToken := envFile["VALIDATION_TOKEN"]
	if authToken != envToken {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	
	// Get and decode the article data from the query parameter
	encodedData := c.Query("data")
	if encodedData == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Missing article data")
	}
	
	jsonData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid data encoding")
	}
	
	var article types.IArticleText
	if err := json.Unmarshal(jsonData, &article); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid article data")
	}
	
	// Get the article from Redis
	val, err := rdb.Get(ctx, article.Title).Result()
	if err != nil {
		log.Printf("Error retrieving article from Redis: %v", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	
	// Unmarshal the retrieved article
	var retrievedArticle types.IArticleText
	if err := json.Unmarshal([]byte(val), &retrievedArticle); err != nil {
		log.Printf("Error deserializing article: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	
	// Set article as validated
	retrievedArticle.Validated = true
	
	// Save the updated article back to Redis
	updatedJsonData, err := json.Marshal(retrievedArticle)
	if err != nil {
		log.Printf("Error serializing article: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	
	if err := rdb.Set(ctx, article.Title, updatedJsonData, 0).Err(); err != nil {
		log.Printf("Error saving article to Redis: %v", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	
	// Return a success page
	c.Set("Content-Type", "text/html")
	return c.SendString(fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Article Published</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					max-width: 800px;
					margin: 0 auto;
					padding: 20px;
					text-align: center;
				}
				.success {
					color: #4CAF50;
					font-size: 24px;
					margin: 20px 0;
				}
				.details {
					background-color: #f8f8f8;
					border-left: 4px solid #4CAF50;
					padding: 15px;
					text-align: left;
					margin: 20px 0;
				}
			</style>
		</head>
		<body>
			<h1 class="success">Article Successfully Published!</h1>
			<div class="details">
				<p><strong>Title:</strong> %s</p>
				<p><strong>Content:</strong></p>
				<p>%s</p>
			</div>
			<p>The article has been validated and published to your site.</p>
		</body>
		</html>
	`, retrievedArticle.Title, retrievedArticle.Content))
}