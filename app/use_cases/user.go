package use_cases

import "queens/app/entities"

type UserRequest struct {
	entities.User
	Name string
	Age  int
}

type DBUser struct {
	// TODO: use cases depend on User ID type int
	ID   int
	Name string
	Age  int
}

type DBClient interface {
	CreateUser(request UserRequest) DBUser
	GetUserByID(id any) DBUser
	UpdateUser(id any, request UserRequest) DBUser
	DeleteUser(id any)
	GetUsers() []DBUser
}

type UserInteractor struct {
	DBClient DBClient
}

func (u *UserInteractor) CreateUser(
	request UserRequest,
) DBUser {
	return u.DBClient.CreateUser(request)
}

func (u *UserInteractor) GetUserByID(id any) DBUser {
	return u.DBClient.GetUserByID(id)
}

func (u *UserInteractor) UpdateUser(
	id any,
	request UserRequest,
) DBUser {
	return u.DBClient.UpdateUser(id, request)
}

func (u *UserInteractor) DeleteUser(id any) {
	u.DBClient.DeleteUser(id)
}

func (u *UserInteractor) GetUsers() []DBUser {
	return u.DBClient.GetUsers()
}
