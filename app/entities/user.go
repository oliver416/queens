package entities

// TODO: remove JSON tags???
// TODO: ID can vary from one DB to another
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
