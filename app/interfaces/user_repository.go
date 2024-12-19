package interfaces

import "queens/app/entities"

type UserRepository interface {
	// TODO: direct access to entities!!!
	AddUser(id any, user entities.User)
	GetUserByID(id any)
	UpdateUser(id any, user entities.User)
	DeleteUser(id any)
	GetUsers() []entities.User
}
