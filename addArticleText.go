package main

import ("github.com/gofiber/fiber/v2"
 "github.com/redis/go-redis/v9"
 "context"
  "blog.local/interfaces/types"
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
    if err:= rdb.Set(ctx, article.Title, article.Content, 0).Err(); err != nil {
      panic(err);
    }
    c.JSON(article);
    return c.SendStatus(200);
  };
