package use_cases

// TODO: there is no need entities package at all((

type UserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type DBClient interface {
	CreateUser(request UserRequest) UserResponse
	GetUserByID(id any) UserResponse
	UpdateUser(id any, request UserRequest) UserResponse
	DeleteUser(id any)
	GetUsers() []UserResponse
}

type UserInteractor struct {
	DBClient DBClient
}

func (u *UserInteractor) CreateUser(
	request UserRequest,
) UserResponse {
	return u.DBClient.CreateUser(request)
}

func (u *UserInteractor) GetUserByID(id any) UserResponse {
	return u.DBClient.GetUserByID(id)
}

func (u *UserInteractor) UpdateUser(
	id any,
	request UserRequest,
) UserResponse {
	return u.DBClient.UpdateUser(id, request)
}

func (u *UserInteractor) DeleteUser(id any) {
	u.DBClient.DeleteUser(id)
}

func (u *UserInteractor) GetUsers() []UserResponse {
	return u.DBClient.GetUsers()
}
