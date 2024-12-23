package app

//import (
//"github.com/gin-gonic/gin"
//swaggerFiles "github.com/swaggo/files"     // swagger embed files
//ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
//"net/http"
//docs "queens/app/docs"
//"queens/app/entities"
//"queens/app/interfaces"
//"strconv"
//)

//// TODO: there is some duplication between the structures User and UserRequest
//type UserRequest struct {
//Name string `json:"name"`
//Age  int    `json:"age"`
//}

//// TODO: remove this
//var DB = []entities.User{
//{ID: 0, Name: "User1", Age: 10},
//{ID: 1, Name: "User2", Age: 20},
//{ID: 2, Name: "User3", Age: 30},
//}

//var UserController struct {
//UserInteractor interfaces.UserInteractorInterface
//}

//func (u *UserController) ServerErrorHandler(context *gin.Context) {
//if r := recover(); r != nil {
//context.JSON(
//http.StatusInternalServerError,
//gin.H{"error": "Internal server error"},
//)
//}
//}

//// Ping godoc
//// @Summary Ping
//// @Schemes
//// @Description Test endpoint
//// @Tags test
//// @Accept json
//// @Produce json
//// @Success 200 {string} Ping
//// @Router /ping [get]
//func (u *UserController) handler(context *gin.Context) {
//context.JSON(
//http.StatusOK,
//gin.H{"status": "ok"},
//)
//}

//// GetUsers godoc
//// @Summary Get users
//// @Schemes
//// @Description Get all users
//// @Tags users
//// @Accept json
//// @Produce json
//// @Success 200 {array} entities.User
//// @Router /users [get]
//func (u *UserController) GetUsers(context *gin.Context) {
//context.JSON(http.StatusOK, DB)
//}

//// GetUser godoc
//// @Summary Get user
//// @Schemes
//// @Description Get user by ID
//// @Tags users
//// @Accept json
//// @Produce json
//// @Param id path int true "User ID"
//// @Success 200 {object} entities.User
//// @Failure 404
//// @Failure 500
//// @Router /users/{id} [get]
//func (u *UserController) GetUser(context *gin.Context) {
//defer u.ServerErrorHandler(context)

//id, err := strconv.Atoi(context.Param("id"))

//if err != nil {
//context.JSON(
//http.StatusNotFound,
//gin.H{"error": "User is not found"},
//)
//return
//}

//u.UserInteractor.GetUser(id)
//context.JSON(http.StatusOK, user)
//}

//// CreateUser godoc
//// @Summary Create user
//// @Schemes
//// @Description Create user
//// @Tags users
//// @Accept json
//// @Produce json
//// @Param request body UserRequest true "User data"
//// @Success 201 {object} entities.User
//// @Failure 400
//// @Failure 500
//// @Router /users [post]
//func (u *UserController) CreateUser(context *gin.Context) {
//var request UserRequest

//if err := context.BindJSON(&request); err != nil {
//// TODO: add logging
//// TODO: standardise error messages
//context.JSON(
//http.StatusBadRequest,
//gin.H{"error": "bad request"},
//)
//return
//}

//u.UserInteractor.CreateUser(id, request)
//context.JSON(http.StatusCreated, user)
//}

//// DeleteUser godoc
//// @Summary Delete user
//// @Schemes
//// @Description Delete user
//// @Tags users
//// @Accept json
//// @Produce json
//// @Param id path int true "User ID"
//// @Success 204
//// @Failure 400
//// @Failure 500
//// @Router /users/{id} [delete]
//func (u *UserController) DeleteUser(context *gin.Context) {
//defer u.ServerErrorHandler(context)

//id, err := strconv.Atoi(context.Param("id"))

//if err != nil {
//context.JSON(
//http.StatusBadRequest,
//gin.H{"error": "bad request"},
//)
//return
//}

//u.UserInteractor.DeleteUser(id)
//context.Status(http.StatusNoContent)
//}

//// UpdateUser godoc
//// @Summary Update user
//// @Schemes
//// @Description Update user
//// @Tags users
//// @Accept json
//// @Produce json
//// @Param id path int true "User ID"
//// @Param request body UserRequest true "User data"
//// @Success 200 {object} entities.User
//// @Failure 400
//// @Failure 500
//// @Router /users/{id} [patch]
//func (u *UserController) UpdateUser(context *gin.Context) {
//defer u.ServerErrorHandler(context)

//var request UserRequest

//if err := context.BindJSON(&request); err != nil {
//// TODO: add logging
//// TODO: standardise error messages
//context.JSON(
//http.StatusBadRequest,
//gin.H{"error": "bad request"},
//)
//return
//}

//ID, err := strconv.Atoi(context.Param("id"))

//if err != nil {
//context.JSON(
//http.StatusNotFound,
//gin.H{"error": "User has not found"},
//)
//return
//}

//u.UserInteractor.UpdateUser(id, request)
//context.JSON(http.StatusOK, user)
//}

//func (u *UserController) Run() {
//app := gin.Default()
//basePath := "/api/v1"
//docs.SwaggerInfo.BasePath = basePath

//group := app.Group(basePath)
//{
//group.GET("/ping", handler)
//group.GET("/users", GetUsers)
//group.GET("/users/:id", GetUser)
//group.POST("/users", CreateUser)
//group.DELETE("/users/:id", DeleteUser)
//group.PATCH("/users/:id", UpdateUser)
//}

//app.GET(
//"/swagger/*any",
//ginSwagger.WrapHandler(swaggerFiles.Handler),
//)
//app.Run()
//}
