package use_cases

import (
	// TODO: repositories access
	"queens/app/entities"
	"queens/app/repositories"
)

type UserInteractor struct {
	Repo repositories.InMemoryRepository
}

func (u *UserInteractor) CreateUser(
	request repositories.UserRequest,
) entities.User {
	return u.Repo.CreateUser(request)
}

func (u *UserInteractor) GetUserByID(id any) entities.User {
	return u.Repo.GetUserByID(id)
}

func (u *UserInteractor) UpdateUser(
	id any,
	request repositories.UserRequest,
) entities.User {
	return u.Repo.UpdateUser(id, request)
}

func (u *UserInteractor) DeleteUser(id any) {
	u.Repo.DeleteUser(id)
}

func (u *UserInteractor) GetUsers() []entities.User {
	return u.Repo.GetUsers()
}
