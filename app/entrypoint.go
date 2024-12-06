package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"net/http"
	docs "queens/app/docs"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// TODO: there is some duplication between the structures
type UserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var DB = []User{
	{ID: 0, Name: "User1", Age: 10},
	{ID: 1, Name: "User2", Age: 20},
	{ID: 2, Name: "User3", Age: 30},
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
// @Success 200 {array} User
// @Router /users [get]
func GetUsers(context *gin.Context) {
	context.JSON(http.StatusOK, DB)
}

// GetUser godoc
// @Summary Get user
// @Schemes
// @Description Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Router /users/{id} [get]
func GetUser(context *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "Internal server error"},
			)
		}
	}()

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User is not found"})
		return
	}

	user := DB[id]
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
// @Success 200 {object} User
// @Router /users [post]
func CreateUser(context *gin.Context) {
	var request UserRequest

	if err := context.BindJSON(&request); err != nil {
		// TODO: add logging
		context.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ID := len(DB)
	user := User{
		ID:   ID,
		Name: request.Name,
		Age:  request.Age,
	}
	DB = append(DB, user)
	context.JSON(http.StatusCreated, user)
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
	}

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.Run()
}
