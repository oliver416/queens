package use_cases

import "queens/app/entities"

type UserInteractor struct {
	Repo UserRepository
}

func (u *UserInteractor) CreateUser(
	id any,
	user entities.User,
) entities.User {
	u.Repo.AddUser(id, user)
	return user
}

func (u *UserInteractor) GetUser(id any) entities.User {
	user := u.Repo.GetUserByID(id)
	return entities.User{Name: user.Name, Age: user.Age}
}

func (u *UserInteractor) UpdateUser(
	id any,
	user entities.User,
) entities.User {
	u.Repo.UpdateUser(id, user)
	return user
}

func (u *UserInteractor) DeleteUser(id any) {
	u.Repo.DeleteUser(id)
}

func (u *UserInteractor) ListUsers() []entities.User {
	return u.Repo.GetUsers()
}
