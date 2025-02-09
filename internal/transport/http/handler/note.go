package handler

import (
	domainDto "go--markdown-note-app/internal/domain/dto"
	"go--markdown-note-app/internal/domain/service"
	transportDto "go--markdown-note-app/internal/transport/http/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NoteHandler struct {
	noteService *service.NoteService
}

func NewNoteHandler(noteService *service.NoteService, router *gin.Engine, prefix string) *NoteHandler {
	handler := &NoteHandler{noteService: noteService}

	//router.GET("/tasks/:id", handler.Update)
	router.POST(prefix+"/notes/create", handler.Create)
	router.POST(prefix+"/notes/get", handler.Get)
	router.POST(prefix+"/notes/render", handler.Render)

	return handler
}

func (h *NoteHandler) Create(c *gin.Context) {
	var dto = transportDto.CreateNoteRequest{}
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	domainDto := domainDto.CreateNoteRequest{
		Text: dto.Text,
	}
	id := (*h.noteService).Create(domainDto)

	c.JSON(
		http.StatusOK,
		transportDto.IdResponse{ID: id},
	)
}

func (h *NoteHandler) Get(c *gin.Context) {
	var dto = transportDto.IdRequest{}
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	note, err := (*h.noteService).Get(dto.ID)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		transportDto.GetNoteResponse{
			ID:   note.ID,
			Text: note.Text,
		},
	)
}

func (h *NoteHandler) Render(c *gin.Context) {
	var dto = transportDto.IdRequest{}
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	md, err := (*h.noteService).Render(dto.ID)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", md.Contents)
}
