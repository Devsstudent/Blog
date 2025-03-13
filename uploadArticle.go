package main

import ("github.com/gofiber/fiber/v2"
 "github.com/redis/go-redis/v9"
 "io"
 "context"
 "blog.local/interfaces/types"
 json "github.com/goccy/go-json"
)

func uploadArticle(c *fiber.Ctx) error {
  ctx := context.Background();
  rdb := redis.NewClient(&redis.Options{
    Addr: "localhost:9000",
    Password: "",
    DB: 0,
  });
  fileHeader, err := c.FormFile("file");
  if err != nil {
    return c.SendStatus(400);
  }
  title := c.FormValue("title");
  file, err := fileHeader.Open();
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).SendString("Failed to open file")
  }
  defer file.Close();
  str, err := io.ReadAll(file);
  if err != nil {
    return c.SendStatus(400);
  }
  article := new(types.IArticleText)
  article.Title = title;
  article.Content = string(str);
  article.Validated = false;
  jsonData, err := json.Marshal(article);
  if err:= rdb.Set(ctx, title, jsonData, 0).Err(); err != nil {
    return c.SendStatus(400);
  }

  return c.SendStatus(200);
}
