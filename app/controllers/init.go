package controllers

import (
	"github.com/revel/revel"
)

//==============================================================================
func init() {
	revel.OnAppStart(initApp)
}

//==============================================================================
func initApp() {
	revel.INFO.Println("Initializing application..")

	InitDB()

	revel.INFO.Println("Application initialized")
}

//==============================================================================
