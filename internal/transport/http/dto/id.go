package dto

type IdRequest struct {
	ID int `json: "id", binding: "required"`
}

type IdResponse struct {
	ID int `json: "id"`
}
