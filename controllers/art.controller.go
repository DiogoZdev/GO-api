package controllers

import (
	"encoding/json"
	"net/http"
)

type Art struct {
	ID 		int		`json:"id"`
	Name 	string 	`json:"name"`
	Artist	string 	`json:"artist"`
	ForSale	bool	`json:"forSale"`
	Price 	int		`json:"price"`
}

type artHandlers struct {
	store map[string]Art
}

func (h *artHandlers) Get(w http.ResponseWriter, r *http.Request) {
	arts := make([]Art, len(h.store))

	i := 0
	for _, art := range h.store {
		arts[i] = art
		i++
	}

	jsonBytes, err := json.Marshal(arts)

	if err != nil {
		panic(err)
	}

	w.Write(jsonBytes)
}

var artsList = map[string]Art{
	"id1": {
		ID: 	 1,
		Name: 	 "Stary Night",
		Artist:	 "Van Gogh",
		ForSale: false,
		Price:   10000,
	},
	"id2": {
		ID: 	 2,
		Name: 	 "Mona Lisa",
		Artist:	 "Da Vinci",
		ForSale: true,
		Price:   225100,
	},
}

func NewArtHandlers() *artHandlers {
	return &artHandlers{
		store: artsList,
	}
}