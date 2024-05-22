package controller

import (
	"encoding/json"
	"net/http"

	entity "root/entities"
)

type artistHandler struct {
	store map[string]entity.Artist
}

func (h *artistHandler) Get(w http.ResponseWriter, r *http.Request) {
	artists := make([]entity.Artist, len(h.store))

	i := 0
	for _, artist := range h.store {
		artists[i] = artist
		i++
	}

	jsonBytes, err := json.Marshal(artists)

	if err != nil {
		panic (err)
	}

	w.Write(jsonBytes)
}

var artistList = map[string]entity.Artist{
	"id1": {
		ID: 1,
		Name: "Van Gogh",
		Birthdate: "1853-03-30",
		Origin: "Zundert, Netherlands",
	},
	"id2": {
		ID: 2,
		Name: "Da Vinci",
		Birthdate: "1452-04-14",
		Origin: "Rome, Italy",
	},
}

func NewArtistsHandle() *artistHandler {
	return &artistHandler{
		store: artistList,
	}
}