package note

import (
	"fmt"
	"go--markdown-note-app/internal/domain/dto"
	"go--markdown-note-app/internal/domain/entity"
	"go--markdown-note-app/internal/domain/interfaces"
	"sync"
)

type noteRepositoryImpl struct {
	items     map[int]noteRepo
	mu        sync.RWMutex
	index_ctr int
}

func CreateRepository() interfaces.NoteRepository {
	return &noteRepositoryImpl{
		items:     make(map[int]noteRepo),
		index_ctr: 0,
	}
}

type noteRepo struct {
	ID   int
	Text string
}

func convertNoteRepo(item noteRepo) entity.Note {
	return entity.Note{
		ID:   item.ID,
		Text: item.Text,
	}
}

func (repo *noteRepositoryImpl) Get(id int) (entity.Note, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	// TODO: not fmt.Errorf, but some custom error
	item, ok := repo.items[id]
	if !ok {
		return entity.Note{}, fmt.Errorf("not found")
	}

	return convertNoteRepo(item), nil
}

func (repo *noteRepositoryImpl) Create(dto dto.CreateNoteRequest) int {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.index_ctr++
	repo.items[repo.index_ctr] = noteRepo{
		ID:   repo.index_ctr,
		Text: dto.Text,
	}

	return repo.index_ctr
}
