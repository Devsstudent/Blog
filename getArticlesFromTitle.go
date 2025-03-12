package main

import ("github.com/gofiber/fiber/v2"
 "github.com/redis/go-redis/v9"
 "context"
"blog.local/interfaces/types"
)

func getArticleFromTitle(c *fiber.Ctx) error {
    ctx := context.Background();
    rdb := redis.NewClient(&redis.Options{
      Addr: "localhost:9000",
      Password: "",
      DB: 0,
    });
    query := c.Queries();
    title := query["title"];

    val, err := rdb.Get(ctx, title).Result();
    if (err != nil) {
      return c.SendStatus(400);
    }
    
    article := types.IArticleText{Content: val, Title: title, Html: convertMarkdownToHtml(val, title)};
    c.JSON(article);
    return c.SendStatus(200);
}
