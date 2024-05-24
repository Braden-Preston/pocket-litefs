package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

// Creating an struct to allow for dependency injection patterns
type Application struct {
	pocket *pocketbase.PocketBase
}

// Method on Appplication struct for getting a list of "products" with PocketBase Go API
func (app *Application) GetProducts(c echo.Context) error {

	info := apis.RequestInfo(c)

	// Get Records for "products" Collection
	products, err := app.pocket.Dao().FindRecordsByExpr("products")
	if err != nil {
		return apis.NewNotFoundError("", err)
	}

	// Check if "products" collection can be "viewed" with the rule set in the Pocketbase dashboard
	canAccess, err := app.pocket.Dao().CanAccessRecord(products[0], info, products[0].Collection().ViewRule)
	if !canAccess {
		return apis.NewForbiddenError("", err)
	}

	return c.JSON(http.StatusOK, products)
}

func main() {

	pocket := pocketbase.New()

	// Adding pocketbase "dependency" to Application
	app := &Application{
		pocket: pocket,
	}

	// Hook for adding in new routes before app starts serving data
	pocket.OnBeforeServe().Add(func(e *core.ServeEvent) error {

		// Useful if you are going to host the build output of a single page app
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))

		// Homepage will return a list of products as JSON data.
		// You could also return HTML with Templ and HTMX instead.
		e.Router.GET("/", app.GetProducts)
		//
		// ... add as many hypermedia routes as you want.

		return nil
	})

	if err := pocket.Start(); err != nil {
		log.Fatal(err)
	}
}
