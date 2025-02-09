package dto

type CreateNoteRequest struct {
	Text string `json: "text", binding: "required"`
}
