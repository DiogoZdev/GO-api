package entities

type Artist struct {
	ID 			int 	`json:"id"`
	Name 		string 	`json:"name"`
	Birthdate 	string 	`json:"birthdate"`
	Origin 		string 	`json:"origin"`
}