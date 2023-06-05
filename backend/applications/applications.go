package applications

import (
	"encoding/json"
	"net/http"

	"github.com/elonsoc/ods/backend/service"
	chi "github.com/go-chi/chi/v5"
)

/*
	This file deals with all things applications
	- recieving the registration form
	- generating a project ID and an API key
	- storing the information in the database
	- querying the database for the user's applications
	- revoking an API key
	- refreshing an API key

	Database Design can be found: https://docs.google.com/document/d/1zEK9K7crTcCcE9bMKm87qRHCyqa5I0_vjuI9eqYSaLg/edit?usp=sharing
*/

// Defining the struct for the applications router
type ApplicationsRouter struct {
	chi.Router
	Svcs *service.Services
}

// NewApplicationsRouter is a function that returns a new applications router
func NewApplicationsRouter(a *ApplicationsRouter) *ApplicationsRouter {
	a.Svcs.Log.Info("Initializing applications router", nil)

	r := chi.NewRouter()
	r.Post("/", a.newApp)
	r.Get("/", a.myApps)
	// r.Put("/{applicationID}", a.applicationByIdHandler)

	a.Router = r
	a.Svcs.Log.Info("Applications router initialized", nil)
	return a
}

// The Application type defines the structure of an application.
type Application struct {
	AppName     string `json:"appName" db:"name"`
	AppID       string `json:"appID" db:"id"`
	Description string `json:"description" db:"description"`
	Owners      string `json:"owners"`
	ApiKey      string `json:"apiKey" db:"api_key"`
	IsValid     bool   `json:"isValid" db:"is_valid"`
}

// This function handles the registration form. It is called when the user submits a registration form.
// It will parse the form, generate a project ID and API key, and store the information in the database.
// It will then return the pertinent information to the user.
func (ar *ApplicationsRouter) newApp(w http.ResponseWriter, r *http.Request) {
	var err error
	// Create a new Application struct
	app := Application{}
	// parse the request body into the application variable
	app.AppName = r.FormValue("title")
	app.Description = r.FormValue("description")
	app.IsValid = true

	// Store the application in the database
	appId, err := ar.Svcs.Db.NewApp(app.AppName, app.Description)
	if err != nil {
		ar.Svcs.Log.Error(err.Error(), nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(appId))
}

// todo(@SheedGuy)
// This function returns a list of all the applications that exist right now.
// Once accessing the users email is figured out, this function will return a
// list of all the applications that the user owns.
func (ar *ApplicationsRouter) myApps(w http.ResponseWriter, r *http.Request) {
	// Query db for all apps
	rows, err := ar.Svcs.Db.GetApplications()
	if err != nil {
		ar.Svcs.Log.Error(err.Error(), nil)
		return
	}

	// Scan the rows into an array of Applications
	apps := []Application{}
	for rows.Next() {
		app := Application{}
		err = rows.Scan(&app.AppName, &app.AppID, &app.Description)
		if err != nil {
			ar.Svcs.Log.Error(err.Error(), nil)
			return
		}
		apps = append(apps, app)
	}

	// Encode the apps array into JSON and send it back to the client
	json.NewEncoder(w).Encode(apps)
}
