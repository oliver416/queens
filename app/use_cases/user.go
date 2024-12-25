package use_cases

import "queens/app/entities"

type User entities.User

type DBUser struct {
	// TODO: use cases depend on User ID type int
	User
	ID   int
	Name string
	Age  int
}

type DBClient interface {
	CreateUser(user User) DBUser
	GetUserByID(id any) DBUser
	UpdateUser(id any, user User) DBUser
	DeleteUser(id any)
	GetUsers() []DBUser
}

type UserInteractor struct {
	DBClient DBClient
}

func (u *UserInteractor) CreateUser(user User) DBUser {
	return u.DBClient.CreateUser(user)
}

func (u *UserInteractor) GetUserByID(id any) DBUser {
	return u.DBClient.GetUserByID(id)
}

func (u *UserInteractor) UpdateUser(id any, user User) DBUser {
	return u.DBClient.UpdateUser(id, user)
}

func (u *UserInteractor) DeleteUser(id any) {
	u.DBClient.DeleteUser(id)
}

func (u *UserInteractor) GetUsers() []DBUser {
	return u.DBClient.GetUsers()
}
