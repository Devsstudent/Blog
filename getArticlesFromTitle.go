package main

import ("github.com/gofiber/fiber/v2"
 "github.com/redis/go-redis/v9"
 "context"
 "blog.local/interfaces/types"
 json "github.com/goccy/go-json"
 "log"
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
    var retrievedArticle types.IArticleText;
    if err := json.Unmarshal([]byte(val), &retrievedArticle); err != nil {
	    log.Fatalf("Error deserializing article: %v", err)
    }
    
    retrievedArticle.Html = convertMarkdownToHtml(retrievedArticle.Content, retrievedArticle.Title);
    c.JSON(retrievedArticle);
    return c.SendStatus(200);
}
