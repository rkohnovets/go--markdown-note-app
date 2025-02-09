package service

import (
	"go--markdown-note-app/internal/domain/dto"
	"go--markdown-note-app/internal/domain/entity"
	"go--markdown-note-app/internal/domain/interfaces"
	"go--markdown-note-app/pkg/utils"
	"log"
)

type NoteService interface {
	Create(dto.CreateNoteRequest) int
	Get(int) (entity.Note, error)
	Render(int) (dto.RenderMarkdownResponse, error)
}

type noteServiceImpl struct {
	repo *interfaces.NoteRepository
}

func CreateNoteService(repo *interfaces.NoteRepository) NoteService {
	if repo == nil {
		log.Fatalf("cannot create note repository")
	}

	return &noteServiceImpl{
		repo: repo,
	}
}

func (service *noteServiceImpl) Create(dto dto.CreateNoteRequest) int {
	id := (*service.repo).Create(dto)
	return id
}

func (service *noteServiceImpl) Get(id int) (entity.Note, error) {
	note, error := (*service.repo).Get(id)
	return note, error
}

func (service *noteServiceImpl) Render(id int) (dto.RenderMarkdownResponse, error) {
	note, err := service.Get(id)
	if err != nil {
		return dto.RenderMarkdownResponse{}, err
	}

	bytes := []byte(note.Text)
	mdBytes := utils.MarkdownToHTML(bytes)

	return dto.RenderMarkdownResponse{
		Contents: mdBytes,
	}, nil
}
