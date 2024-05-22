package entities

type Art struct {
	ID 		int		`json:"id"`
	Name 	string 	`json:"name"`
	Artist	string 	`json:"artist"`
	ForSale	bool	`json:"forSale"`
	Price 	int		`json:"price"`
}