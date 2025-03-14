package main

import ("github.com/gofiber/fiber/v2"
 "github.com/gofiber/fiber/v2/middleware/filesystem"
 "net/http"
 "blog.local/render"
 "io/fs"
)

func routing(app *fiber.App) {

  api := app.Group("/api");

  index, err := fs.Sub(render.Dist, "dist")
	if err != nil {
		panic(err)
  };

	app.Use("/", filesystem.New(filesystem.Config{
		Root:   http.FS(index),
		Index:  "index.html",
		Browse: false,
	}))

  api.Post("/uploadArticle", uploadArticle);
  api.Get("/getAllArticles", getAllArticles);
  api.Post("/addArticleText", addArticleText);
  api.Get("/getArticleFromTitle", getArticleFromTitle);
  api.Get("/validatePosting", validatePosting);
  // Envoyer un email quand on ajoute ou post un articles pour demander la validation

  app.Use(func(c *fiber.Ctx) error {
		path := c.Path()
		if len(path) > 4 && path[:4] == "/api" {
			return c.SendStatus(404) // API not found
		}
		return c.SendFile("render/dist/index.html")
	})
}
