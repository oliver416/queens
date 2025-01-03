package interfaces

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"net/http"
	docs "queens/app/docs"
	"queens/app/use_cases"
	"strconv"
)

type User = use_cases.User
type DBUser = use_cases.DBUser

type UserRequest struct {
	User
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Interactor interface {
	CreateUser(user User) DBUser
	GetUserByID(id any) DBUser
	UpdateUser(id any, user User) DBUser
	DeleteUser(id any)
	GetUsers() []DBUser
}

type GinController struct {
	Interactor Interactor
}

func (g *GinController) ServerErrorHandler(context *gin.Context) {
	if r := recover(); r != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Internal server error"},
		)
	}
}

// Ping godoc
// @Summary Ping
// @Schemes
// @Description Test endpoint
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {string} Ping
// @Router /ping [get]
func (g *GinController) handler(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		gin.H{"status": "ok"},
	)
}

// GetUsers godoc
// @Summary Get users
// @Schemes
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} UserResponse
// @Router /users [get]
func (g *GinController) GetUsers(context *gin.Context) {
	users := g.Interactor.GetUsers()
	context.JSON(http.StatusOK, users)
}

// GetUser godoc
// @Summary Get user
// @Schemes
// @Description Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} UserResponse
// @Failure 404
// @Failure 500
// @Router /users/{id} [get]
func (g *GinController) GetUser(context *gin.Context) {
	defer g.ServerErrorHandler(context)

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{"error": "User is not found"},
		)
		return
	}

	user := g.Interactor.GetUserByID(id)
	context.JSON(http.StatusOK, user)
}

// CreateUser godoc
// @Summary Create user
// @Schemes
// @Description Create user
// @Tags users
// @Accept json
// @Produce json
// @Param request body UserRequest true "User data"
// @Success 201 {object} UserResponse
// @Failure 400
// @Failure 500
// @Router /users [post]
func (g *GinController) CreateUser(context *gin.Context) {
	var request UserRequest

	if err := context.BindJSON(&request); err != nil {
		// TODO: add logging
		// TODO: standardise error messages
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "bad request"},
		)
		return
	}

	// TODO: it looks like some sort of duplication
	user := User{
		Name: request.Name,
		Age:  request.Age,
	}
	response := g.Interactor.CreateUser(user)
	context.JSON(http.StatusCreated, response)
}

// DeleteUser godoc
// @Summary Delete user
// @Schemes
// @Description Delete user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /users/{id} [delete]
func (g *GinController) DeleteUser(context *gin.Context) {
	defer g.ServerErrorHandler(context)

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "bad request"},
		)
		return
	}

	g.Interactor.DeleteUser(id)
	context.Status(http.StatusNoContent)
}

// UpdateUser godoc
// @Summary Update user
// @Schemes
// @Description Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body UserRequest true "User data"
// @Success 200 {object} UserResponse
// @Failure 400
// @Failure 500
// @Router /users/{id} [patch]
func (g *GinController) UpdateUser(context *gin.Context) {
	defer g.ServerErrorHandler(context)

	var request UserRequest

	if err := context.BindJSON(&request); err != nil {
		// TODO: add logging
		// TODO: standardise error messages
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "bad request"},
		)
		return
	}

	ID, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{"error": "User has not found"},
		)
		return
	}

	// TODO: duplication in CreateUser
	user := User{
		Name: request.Name,
		Age:  request.Age,
	}
	response := g.Interactor.UpdateUser(ID, user)
	context.JSON(http.StatusOK, response)
}

func (g *GinController) Run() {
	app := gin.Default()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	group := app.Group(basePath)
	{
		group.GET("/ping", g.handler)
		group.GET("/users", g.GetUsers)
		group.GET("/users/:id", g.GetUser)
		group.POST("/users", g.CreateUser)
		group.DELETE("/users/:id", g.DeleteUser)
		group.PATCH("/users/:id", g.UpdateUser)
	}

	app.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)
	app.Run()
}
