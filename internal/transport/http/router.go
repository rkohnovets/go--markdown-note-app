package http

import (
	"go--markdown-note-app/internal/domain/service"
	"go--markdown-note-app/internal/transport/http/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(ns *service.NoteService) *gin.Engine {
	router := gin.Default()
	handler.NewNoteHandler(ns, router, "")
	return router
}
