package app

import (
	"queens/app/interfaces"
	"queens/app/use_cases"
)

func Run() {
	DB := []use_cases.DBUser{
		{ID: 0, Name: "User1", Age: 10},
		{ID: 1, Name: "User2", Age: 20},
		{ID: 2, Name: "User3", Age: 30},
	}

	db_client := interfaces.InMemoryDBClient{DB: DB}
	interactor := use_cases.UserInteractor{DBClient: &db_client}
	controller := interfaces.GinController{Interactor: &interactor}
	controller.Run()
}
