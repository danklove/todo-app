package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/todo-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.singUp)
		auth.POST("/sing-in", h.singIn)
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

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				lists.GET("/", h.getAllItems)
				lists.GET("/:item_id", h.getItemById)
				lists.PUT("/:item_id", h.updateItem)
				lists.DELETE("/:item_id", h.deleteItem)
			}
		}
	}

	return router
}
