package main

import ("github.com/gofiber/fiber/v2"
 "github.com/redis/go-redis/v9"
 "blog.local/interfaces/types"
 "context"
 json "github.com/goccy/go-json"
 "log"
)

func getAllArticles(c *fiber.Ctx) error {
  ctx := context.Background();
  rdb := redis.NewClient(&redis.Options{
    Addr: "localhost:9000",
    Password: "",
    DB: 0,
  });
  // This get all the key in the DB we are connected to
  keys, err := rdb.Keys(ctx, "*").Result();
  if err != nil {
    c.SendStatus(400);
  }
  //Checker la len de val maybe si 0 
  var allArticles []types.IArticleText;


  for _, key:= range(keys) {
    val, err := rdb.Get(ctx, key).Result();
    if (err != nil) {
      return c.SendStatus(400);
    }
    var retrievedArticle types.IArticleText;
    if err := json.Unmarshal([]byte(val), &retrievedArticle); err != nil {
	    log.Fatalf("Error deserializing article: %v", err)
    }
    
    retrievedArticle.Html = convertMarkdownToHtml(retrievedArticle.Content, retrievedArticle.Title);
    allArticles = append(allArticles, retrievedArticle);
  }
  c.JSON(allArticles);
  return c.SendStatus(200);
};
