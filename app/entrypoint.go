package app

import "queens/app/controllers"

func Run() {
	controller := controllers.GinController{}
	controller.Run()
}
