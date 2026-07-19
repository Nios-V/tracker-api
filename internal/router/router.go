package router

import (
	"github.com/Nios-V/tracker-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterCRUD[T any](group *gin.RouterGroup, h *handlers.GenericHandler[T]) {
	group.POST("", h.Create)
	group.GET("/:id", h.GetByID)
	group.GET("", h.GetAll)
	group.PUT("/:id", h.Update)
	group.DELETE("/:id", h.Delete)
}

func SetupRouter(h *handlers.Handlers) *gin.Engine {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	api := r.Group("/api/v1")
	RegisterCRUD(api.Group("/tags"), h.Tag)
	RegisterCRUD(api.Group("/media"), h.Media)
	RegisterCRUD(api.Group("/media-tags"), h.MediaTag)
	RegisterCRUD(api.Group("/experiences"), h.Experience)
	RegisterCRUD(api.Group("/sessions"), h.Session)
	RegisterCRUD(api.Group("/notes"), h.Note)
	RegisterCRUD(api.Group("/note-links"), h.NoteLink)
	return r
}
