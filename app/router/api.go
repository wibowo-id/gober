package router

import (
	"github.com/bangadam/go-fiber-starter/app/module/article"
	"github.com/bangadam/go-fiber-starter/utils/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Router struct {
	App           fiber.Router
	Cfg           *config.Config
	ArticleRouter *article.ArticleRouter
}

func NewRouter(fiber *fiber.App, cfg *config.Config, articleRouter *article.ArticleRouter) *Router {
	return &Router{
		App:           fiber,
		Cfg:           cfg,
		ArticleRouter: articleRouter,
	}
}

// Register routes
func (r *Router) Register() {
	// Test Routes
	r.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	// Swagger Documentation
	r.App.Get("/swagger/*", swagger.HandlerDefault)

	// Register routes of modules
	r.ArticleRouter.RegisterArticleRoutes()
}
