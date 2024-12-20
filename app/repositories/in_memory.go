package repositories

import "queens/app/interfaces"

type InMemoryRepository struct {
	// TODO: protect the DB???
	DB []interfaces.UserDB
}

func (r *InMemoryRepository) AnyToInt(value any) *int {
	number, ok := value.(int)

	if ok {
		return &number
	}

	return nil
}

func (r *InMemoryRepository) AddUser(
	request interfaces.UserRequest,
) interfaces.UserDB {
	ID := len(r.DB)
	user := interfaces.UserDB{
		ID:   ID,
		Name: request.Name,
		Age:  request.Age,
	}
	r.DB = append(r.DB, user)
	return user
}

func (r *InMemoryRepository) GetUserByID(id any) interfaces.UserDB {
	index := r.AnyToInt(id)

	if index != nil {
		return r.DB[*index]
	}

	return interfaces.UserDB{}
}

func (r *InMemoryRepository) UpdateUser(
	id any,
	request interfaces.UserRequest,
) {
	index := r.AnyToInt(id)

	if index == nil {
		return
	}

	user := &r.DB[*index]

	//TODO: how to avoid duplication??
	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Age != 0 {
		user.Age = request.Age
	}
}

func (r *InMemoryRepository) DeleteUser(id any) {
	index := r.AnyToInt(id)

	if index == nil {
		return
	}

	r.DB = append(r.DB[:*index], r.DB[*index+1:]...)
}

func (r *InMemoryRepository) GetUsers() []interfaces.UserDB {
	return r.DB
}
