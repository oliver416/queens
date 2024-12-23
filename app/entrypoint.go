package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"net/http"
	docs "queens/app/docs"
	"queens/app/entities"
	"queens/app/repositories"
	"strconv"
)

var DB = []entities.User{
	{ID: 0, Name: "User1", Age: 10},
	{ID: 1, Name: "User2", Age: 20},
	{ID: 2, Name: "User3", Age: 30},
}

var Repo = repositories.InMemoryRepository{DB: DB}

func ServerErrorHandler(context *gin.Context) {
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
func handler(context *gin.Context) {
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
// @Success 200 {array} entities.User
// @Router /users [get]
func GetUsers(context *gin.Context) {
	users := Repo.GetUsers()
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
// @Success 200 {object} entities.User
// @Failure 404
// @Failure 500
// @Router /users/{id} [get]
func GetUser(context *gin.Context) {
	defer ServerErrorHandler(context)

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{"error": "User is not found"},
		)
		return
	}

	user := Repo.GetUserByID(id)
	context.JSON(http.StatusOK, user)
}

// CreateUser godoc
// @Summary Create user
// @Schemes
// @Description Create user
// @Tags users
// @Accept json
// @Produce json
// @Param request body repositories.UserRequest true "User data"
// @Success 201 {object} entities.User
// @Failure 400
// @Failure 500
// @Router /users [post]
func CreateUser(context *gin.Context) {
	var request repositories.UserRequest

	if err := context.BindJSON(&request); err != nil {
		// TODO: add logging
		// TODO: standardise error messages
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "bad request"},
		)
		return
	}

	user := Repo.CreateUser(request)
	context.JSON(http.StatusCreated, user)
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
func DeleteUser(context *gin.Context) {
	defer ServerErrorHandler(context)

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "bad request"},
		)
		return
	}

	Repo.DeleteUser(id)
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
// @Param request body repositories.UserRequest true "User data"
// @Success 200 {object} entities.User
// @Failure 400
// @Failure 500
// @Router /users/{id} [patch]
func UpdateUser(context *gin.Context) {
	defer ServerErrorHandler(context)

	var request repositories.UserRequest

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

	user := Repo.UpdateUser(ID, request)
	context.JSON(http.StatusOK, user)
}

func Run() {
	app := gin.Default()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	group := app.Group(basePath)
	{
		group.GET("/ping", handler)
		group.GET("/users", GetUsers)
		group.GET("/users/:id", GetUser)
		group.POST("/users", CreateUser)
		group.DELETE("/users/:id", DeleteUser)
		group.PATCH("/users/:id", UpdateUser)
	}

	app.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)
	app.Run()
}
