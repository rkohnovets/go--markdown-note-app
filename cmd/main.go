package main

import (
	"go--markdown-note-app/internal/domain/service"
	"go--markdown-note-app/internal/repository/in_memory/note"
	"go--markdown-note-app/internal/transport/http"
	"log"
)

func main() {
	repo := note.CreateRepository()
	ns := service.CreateNoteService(&repo)

	router := http.SetupRouter(&ns)

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("failed to run gin http server, error: %v", err.Error())
	}
}
