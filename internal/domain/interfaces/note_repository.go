package interfaces

import (
	"go--markdown-note-app/internal/domain/dto"
	"go--markdown-note-app/internal/domain/entity"
)

type NoteRepository interface {
	Get(int) (entity.Note, error)
	Create(dto.CreateNoteRequest) int
}
