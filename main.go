package main

import ("github.com/gofiber/fiber/v2"
 "github.com/gofiber/fiber/v2/middleware/filesystem"
 "github.com/gofiber/fiber/v2/middleware/cors"
 "github.com/redis/go-redis/v9"
 "net/http"
 "blog.local/render"
 "fmt"
 "io/fs"
 "context"
)


// Il me faut 1 route pour poster un post dans la db
// 1 other pour les gets aussi

func main() {
  app := fiber.New();


  // API routes
  app.Get("/api/coucou", func(c *fiber.Ctx) error {
	ctx := context.Background();
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:9000",
		Password: "",
		DB: 0,
	});
	if err := rdb.Set(ctx, "key", "value", 0).Err(); err != nil {
		panic(err);
	}

	val, err := rdb.Get(ctx, "key").Result();
	if err != nil {
		panic(err);
	}
	fmt.Println("key", val);
    return c.SendString(val);
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

  fmt.Println("Listening on port 8080...");
  app.Listen(":8080");
}
