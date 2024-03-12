package locations

import (
	buildings_v1 "github.com/elonsoc/ods/backend/locations/v1/buildings"
	"github.com/elonsoc/ods/backend/service"
	chi "github.com/go-chi/chi/v5"
)

// LocationsRouter is the router for the locations service
// It contains a pointer to the chi router and a pointer to the service struct
// that contains the various services that we will be using in the backend.
type LocationsRouter struct {
	chi.Router
	Svcs *service.Services
}

// NewLocationsRouter creates a new instance of the LocationsRouter struct
// and returns a pointer to it.
// This function takes a pointer to the LocationsRouter struct as an argument
// so that we can pass the service struct around by reference.
// By passing the service struct by reference, we can make changes to the
// struct and those changes will be reflected in the original struct.
// By returning a pointer to the struct, we can use the struct
// in the top level to mount the router defined here to the main router.
func NewLocationsRouter(l *LocationsRouter) *LocationsRouter {
	r := chi.NewRouter()
	l.Svcs.Log.Info("Initializing locations router", nil)
	// Here we're mounting a group of routers to the v1 route.
	// The reason why we're nesting a group of routers is because
	// we want to be able to version the API.
	// This way, we can make changes or improvements to the API without breaking
	// the existing API.
	// we could go further and define this in the v1/ directory but it is not
	// necessary for now.
	r.Mount("/v1", r.Group(func(r chi.Router) {
		r.Mount("/buildings", buildings_v1.NewBuildingsRouter(&buildings_v1.BuildingsRouter{Svcs: l.Svcs}).Router)
	}))

	l.Svcs.Log.Info("Finished initializing locations router", nil)
	return &LocationsRouter{
		Router: r,
	}
}
