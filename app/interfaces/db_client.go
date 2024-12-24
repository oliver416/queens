package interfaces

import "queens/app/use_cases"

type InMemoryDBClient struct {
	DB []use_cases.UserResponse
}

// TODO: UUID cannot be an ID in case of in-memory database, so the following
// function is uselesss
func (r *InMemoryDBClient) AnyToInt(value any) *int {
	number, ok := value.(int)

	if ok {
		return &number
	}

	return nil
}

func (r *InMemoryDBClient) CreateUser(
	request use_cases.UserRequest,
) use_cases.UserResponse {
	ID := len(r.DB)
	user := use_cases.UserResponse{
		ID:   ID,
		Name: request.Name,
		Age:  request.Age,
	}
	r.DB = append(r.DB, user)
	return user
}

func (r *InMemoryDBClient) GetUserByID(id any) use_cases.UserResponse {
	index := r.AnyToInt(id)

	if index != nil {
		return r.DB[*index]
	}

	return use_cases.UserResponse{}
}

func (r *InMemoryDBClient) UpdateUser(
	id any,
	request use_cases.UserRequest,
) use_cases.UserResponse {
	index := r.AnyToInt(id)

	if index == nil {
		return use_cases.UserResponse{}
	}

	user := &r.DB[*index]

	//TODO: how to avoid duplication??
	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Age != 0 {
		user.Age = request.Age
	}

	return *user
}

func (r *InMemoryDBClient) DeleteUser(id any) {
	index := r.AnyToInt(id)

	if index == nil {
		return
	}

	r.DB = append(r.DB[:*index], r.DB[*index+1:]...)
}

func (r *InMemoryDBClient) GetUsers() []use_cases.UserResponse {
	return r.DB
}
