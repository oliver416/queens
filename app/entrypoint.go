package app

import (
	"queens/app/controllers"
	"queens/app/entities"
	"queens/app/repositories"
	"queens/app/use_cases"
)

func Run() {
	DB := []entities.User{
		{ID: 0, Name: "User1", Age: 10},
		{ID: 1, Name: "User2", Age: 20},
		{ID: 2, Name: "User3", Age: 30},
	}

	repository := repositories.InMemoryRepository{DB: DB}
	interactor := use_cases.UserInteractor{Repo: repository}
	controller := controllers.GinController{Interactor: interactor}
	controller.Run()
}
