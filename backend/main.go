package main

import ("github.com/gofiber/fiber/v2"
 "github.com/gofiber/fiber/v2/middleware/filesystem"
 "github.com/gofiber/fiber/v2/middleware/cors"
 "net/http"
)

func main() {
  app := fiber.New();

  // API routes
  app.Get("/api/coucou", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  });

  app.Use(cors.New())

    // Serve static files from React build
  app.Use("/", filesystem.New(filesystem.Config{
      Root: http.Dir("../render/dist"),
      Browse: true,
  }))


    // This is important for React Router
  app.Get("/*", func(c *fiber.Ctx) error {
      return c.SendFile("../render/dist/index.html")
   })
//  fmt.Println("Listening on port 6000...");
  app.Listen(":8080");
}
