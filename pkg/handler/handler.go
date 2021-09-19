package handler

import (
	"LinkSaver/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
		}

		links := api.Group(":id/links")
		{
			links.POST("/", h.createLink)
			links.GET("/", h.getAllLinks)
			links.GET("/:link_id", h.getLinkById)
			links.PUT("/:link_id", h.updateLink)
			links.DELETE("/:link_id", h.deleteLink)
		}
	}

	return router
}
