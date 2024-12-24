package use_cases

import "queens/app/entities"

type UserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type DBClient interface {
	CreateUser(request UserRequest) entities.User
	GetUserByID(id any) entities.User
	UpdateUser(id any, request UserRequest) entities.User
	DeleteUser(id any)
	GetUsers() []entities.User
}

type UserInteractor struct {
	DBClient DBClient
}

func (u *UserInteractor) CreateUser(
	request UserRequest,
) entities.User {
	return u.DBClient.CreateUser(request)
}

func (u *UserInteractor) GetUserByID(id any) entities.User {
	return u.DBClient.GetUserByID(id)
}

func (u *UserInteractor) UpdateUser(
	id any,
	request UserRequest,
) entities.User {
	return u.DBClient.UpdateUser(id, request)
}

func (u *UserInteractor) DeleteUser(id any) {
	u.DBClient.DeleteUser(id)
}

func (u *UserInteractor) GetUsers() []entities.User {
	return u.DBClient.GetUsers()
}
