package dto

type GetNoteResponse struct {
	ID   int    `json: "id"`
	Text string `json: "text"`
}
