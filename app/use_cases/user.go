package use_cases

import "queens/app/entities"

type UserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Repository interface {
	CreateUser(request UserRequest) entities.User
	GetUserByID(id any) entities.User
	UpdateUser(id any, request UserRequest) entities.User
	DeleteUser(id any)
	GetUsers() []entities.User
}

type UserInteractor struct {
	Repo Repository
}

func (u *UserInteractor) CreateUser(
	request UserRequest,
) entities.User {
	return u.Repo.CreateUser(request)
}

func (u *UserInteractor) GetUserByID(id any) entities.User {
	return u.Repo.GetUserByID(id)
}

func (u *UserInteractor) UpdateUser(
	id any,
	request UserRequest,
) entities.User {
	return u.Repo.UpdateUser(id, request)
}

func (u *UserInteractor) DeleteUser(id any) {
	u.Repo.DeleteUser(id)
}

func (u *UserInteractor) GetUsers() []entities.User {
	return u.Repo.GetUsers()
}
