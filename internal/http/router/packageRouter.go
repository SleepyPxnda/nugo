package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"server/internal/http/controller"
)

func PackageRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", controller.GetPackageHandler)
	r.Put("/", controller.CreatePackageHandler)
	r.Post("/{ID}/{VERSION}", controller.UpdatePackageHandler)
	r.Delete("/{ID}/{VERSION}", controller.DeletePackageHandler)
	return r
}
