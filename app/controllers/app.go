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
	return c.Render()
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
func (c App) Update(id int64, commentsenabled bool) revel.Result {
	var page models.Page
	DB.Find(&page, id)
	if page.Id == 0 {
		return c.NotFound("Page %v not found", id)
	}

	upage := models.Page{CommentsEnabled: commentsenabled}
	DB.Model(&page).Update(upage)

	return c.Redirect(routes.App.Show(id))
}

//==============================================================================
