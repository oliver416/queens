package interfaces

type UserDB struct {
	// TODO: int ID is not abstract!!
	ID   int
	Name string
	Age  int
}

// TODO: there is some duplication between the structures User and UserRequest
type UserRequest struct {
	Name string
	Age  int
}

type UserRepository interface {
	AddUser(request UserRequest) UserDB
	GetUserByID(id any) UserDB
	UpdateUser(id any, request UserRequest)
	DeleteUser(id any)
	GetUsers() []UserDB
}
