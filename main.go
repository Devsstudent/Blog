package main

import ("github.com/gofiber/fiber/v2"
 "github.com/gofiber/fiber/v2/middleware/filesystem"
 "github.com/gofiber/fiber/v2/middleware/cors"
 "net/http"
 "blog.local/render"
 "fmt"
  "io/fs"
)

func main() {
  app := fiber.New();

  // API routes
  app.Get("/api/coucou", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  });
  index, err := fs.Sub(render.Dist, "dist")
	if err != nil {
		panic(err)
	}

	app.Use("/", filesystem.New(filesystem.Config{
		Root:   http.FS(index),
		Index:  "index.html",
		Browse: false,
	}))

  app.Use(cors.New());

    // Serve static files from React build
  app.Use("/", filesystem.New(filesystem.Config{
        Root:       http.FS(render.Dist),
        Index:      "index.html", 
        Browse:     true,
    }))

  fmt.Println("Listening on port 6000...");
  app.Listen(":8080");
}
