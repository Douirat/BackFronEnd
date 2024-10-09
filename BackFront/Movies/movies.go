package movies

import (
// "fmt"
// "encoding/json"
)

// Declare a movie object the struct togather data related to the movies:
type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: title`
	Director *Director `json: "director"`
}

// Declare a director struct:
type Director struct {
	FirstName string `json: "firstName"`
	LastName  string `json: "lastName"`
}
