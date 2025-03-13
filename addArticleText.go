package main

import (
	"context"

	"blog.local/interfaces/types"
	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func addArticleText(c *fiber.Ctx) error {
    article := new(types.IArticleText)
    if err := c.BodyParser(article); err != nil {
      return c.SendStatus(400);
    }
    ctx := context.Background();
	  rdb := redis.NewClient(&redis.Options{
	    Addr: "localhost:9000",
	    Password: "",
	    DB: 0,
	  });
    article.Validated = false;
    jsonData, err := json.Marshal(article);
    if (err != nil) {
      panic(err);
    }
    if err:= rdb.Set(ctx, article.Title, jsonData, 0).Err(); err != nil {
      panic(err)
    }
    c.JSON(article);
    sendEmail(*article);
    return c.SendStatus(200);
  };
