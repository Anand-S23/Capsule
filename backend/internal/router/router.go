package router

import (
	"net/http"

	"github.com/Anand-S23/capsule/internal/controller"
	"github.com/Anand-S23/capsule/internal/middleware"
	"github.com/gorilla/handlers"
)

func NewRouter(c *controller.Controller) *http.ServeMux {
    router := http.NewServeMux()

    // Health Check
    router.HandleFunc("GET /ping", Fn(c.Ping))

    // Auth
    router.HandleFunc("POST /register", Fn(c.Register))
    router.HandleFunc("POST /login", Fn(c.Login))
    router.HandleFunc("POST /logout", Fn(c.Logout))
    router.HandleFunc("GET /getAuthUserID", middleware.Auth(Fn(c.GetAuthUser), c))

    // Connection
    router.HandleFunc("POST /connection/create", middleware.Auth(Fn(c.CreateConnection), c))
    router.HandleFunc("GET /connection/get/{id}", middleware.Auth(Fn(c.GetConnection), c))
    router.HandleFunc("GET /connection/getAll", middleware.Auth(Fn(c.GetAllConnections), c))
    router.HandleFunc("POST /connection/update/{id}", middleware.Auth(Fn(c.UpdateConnection), c))
    router.HandleFunc("POST /connection/delete/{id}", middleware.Auth(Fn(c.DeleteConnection), c))

    return router
}

func NewCorsRouter(router *http.ServeMux, allowedOrigin string) http.Handler {
    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:3000", allowedOrigin}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

    return corsHandler(router)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func Fn(fn apiFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := fn(w, r)
        if err != nil {
            controller.WriteJSON(w, http.StatusInternalServerError, controller.ErrMsg(err.Error()))
        }
    }
}

