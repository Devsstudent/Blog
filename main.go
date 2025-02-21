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
  json "github.com/goccy/go-json"
)

type ArticleText struct {
  Content string `json:"content"`;
  Title   string `json:"title"`;
}

// Encoding
//	p := Person{Name: "Alice", Age: 25}
//	data, _ := json.Marshal(p)
//	fmt.Println(string(data)) // Output: {"name":"Alice","age":25}
//
//	// Decoding
//	var p2 Person
//	_ = json.Unmarshal(data, &p2)
//	fmt.Println(p2.Name, p2.Age) // Output: Alice 25


func main() {
  app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,   // Use go-json for encoding
		JSONDecoder: json.Unmarshal, // Use go-json for decoding
	});

  app.Use(cors.New());

  api := app.Group("/api");

  // API routes

  api.Get("/coucou", func(c *fiber.Ctx) error {
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
	fmt.Println("keyo", val);
    return c.SendString(val);
  });


  api.Get("/getAllArticles", func(c *fiber.Ctx) error {
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

    var allArticles []ArticleText;

    for _, key:= range(keys) {
      val, err := rdb.Get(ctx, key).Result();
      if (err != nil) {
        return c.SendStatus(400);
      }
      allArticles = append(allArticles, ArticleText{Content: val, Title: key});
    }
    c.JSON(allArticles);
    return c.SendStatus(200);
  });


  api.Post("/addArticleText", func(c *fiber.Ctx) error {
    article := new(ArticleText)
    if err := c.BodyParser(article); err != nil {
      fmt.Println(err);
      return c.SendStatus(400);
    }
    fmt.Println(article.Content, "yes");
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
  });

  index, err := fs.Sub(render.Dist, "dist")
	if err != nil {
		panic(err)

  };

	app.Use("/", filesystem.New(filesystem.Config{
		Root:   http.FS(index),
		Index:  "index.html",
		Browse: false,
	}))

  app.Use(func(c *fiber.Ctx) error {
		path := c.Path()
		if len(path) > 4 && path[:4] == "/api" {
      fmt.Println("TESTTTTT");
			return c.SendStatus(404) // API not found
		}
		c.SendFile("render/dist/index.html")
    return c.SendStatus(200);
	})


  fmt.Println("Listening on port 8080...");
  app.Listen(":8080");
}
