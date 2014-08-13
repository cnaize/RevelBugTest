package controllers

import (
	"github.com/revel/revel"
	"test/app/models"
	"test/app/routes"
)

//==============================================================================
// App
//================================================
type App struct {
	*revel.Controller
}

//==============================================================================
func (c App) Index() revel.Result {
	page := models.Page{}
	DB.Save(&page)

	revel.INFO.Println(page)

	return c.Redirect(routes.App.Show(page.Id))
}

//==============================================================================
func (c App) Show(id int64) revel.Result {
	var page models.Page
	DB.Find(&page, id)
	if page.Id == 0 {
		return c.NotFound("Page %v not found", id)
	}

	return c.Render(page)
}

//==============================================================================
