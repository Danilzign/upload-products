package handler

import (
	"net/http"

	Image "github.com/danilzign/todo-app/pkg/image"
	"github.com/danilzign/todo-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		name := api.Group("/name")
		{
			name.POST("/", h.createProduct)
			name.GET("/", h.getAllProducts)
			name.GET("/:id", h.getProductById)
			name.PUT("/:id", h.updateProduct)
			name.DELETE("/:id", h.deleteProduct)
		}

		images := api.Group("/image")
		{
			images.Use(LiberalCORS)
			images.Static("/stat-img", "./image")
			images.Static("/static", "./webClient/dist")
			images.Static("/statics", "./webClient/dist/statics")
			images.StaticFile("/", "./webClient/dist/index.html")

			images.POST("/upload_image", Image.SimpleUploadImage)
		}
	}
	router.NoRoute(func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "./webClient/dist/index.html")
	})

	return router
}

func LiberalCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == "OPTIONS" {
		if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
			c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
		}
		c.AbortWithStatus(http.StatusOK)
	}
}
