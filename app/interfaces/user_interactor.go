package interfaces

// TODO: direct access to entities
import "queens/app/entities"

type UserInteractorInterface interface {
	CreateUser(id any, user entities.User) entities.User
	GetUser(id any) entities.User
	UpdateUser(id any, user entities.User) entities.User
	DeleteUser(id any)
	ListUsers() []entities.User
}
