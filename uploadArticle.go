package main

import ("github.com/gofiber/fiber/v2"
 "github.com/redis/go-redis/v9"
 "io"
 "context"
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
  if err:= rdb.Set(ctx, title, string(str), 0).Err(); err != nil {
    return c.SendStatus(400);
  }

  return c.SendStatus(200);
}
