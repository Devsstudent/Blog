package main

import (
	"context"
	"log"

	"blog.local/interfaces/types"
	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func validatePosting(c *fiber.Ctx) error {
	ctx := context.Background();
	rdb := redis.NewClient(&redis.Options{
	  Addr: "localhost:9000",
	  Password: "",
	  DB: 0,
	});

	envFile, _ := godotenv.Read(".env")

	envToken := envFile["VALIDATION_TOKEN"]

	token := c.Get("Authorization")
	if (token != envToken) {
		return c.SendStatus(401);
	}
	article := new(types.IArticleText)
    if err := c.BodyParser(article); err != nil {
      return c.SendStatus(400);
    }
	val, err := rdb.Get(ctx, article.Title).Result();
    if (err != nil) {
      return c.SendStatus(400);
    }
    var retrievedArticle types.IArticleText;
    if err := json.Unmarshal([]byte(val), &retrievedArticle); err != nil {
	    log.Fatalf("Error deserializing article: %v", err)
    }
	retrievedArticle.Validated = true;
	
	jsonData, err := json.Marshal(retrievedArticle);
	if err:= rdb.Set(ctx, article.Title, jsonData, 0).Err(); err != nil {
	  return c.SendStatus(400);
	}
  
	return c.SendStatus(200);
}