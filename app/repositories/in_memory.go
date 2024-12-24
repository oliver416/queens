package repositories

// TODO: direct access to entities
import "queens/app/entities"
import "queens/app/use_cases"

type InMemoryRepository struct {
	// TODO: protect the DB???
	DB []entities.User
}

// TODO: UUID cannot be an ID in case of in-memory database, so the following
// function is uselesss
func (r *InMemoryRepository) AnyToInt(value any) *int {
	number, ok := value.(int)

	if ok {
		return &number
	}

	return nil
}

func (r *InMemoryRepository) CreateUser(
	request use_cases.UserRequest,
) entities.User {
	ID := len(r.DB)
	user := entities.User{
		ID:   ID,
		Name: request.Name,
		Age:  request.Age,
	}
	r.DB = append(r.DB, user)
	return user
}

func (r *InMemoryRepository) GetUserByID(id any) entities.User {
	index := r.AnyToInt(id)

	if index != nil {
		return r.DB[*index]
	}

	return entities.User{}
}

func (r *InMemoryRepository) UpdateUser(
	id any,
	request use_cases.UserRequest,
) entities.User {
	index := r.AnyToInt(id)

	if index == nil {
		return entities.User{}
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

func (r *InMemoryRepository) DeleteUser(id any) {
	index := r.AnyToInt(id)

	if index == nil {
		return
	}

	r.DB = append(r.DB[:*index], r.DB[*index+1:]...)
}

func (r *InMemoryRepository) GetUsers() []entities.User {
	return r.DB
}
