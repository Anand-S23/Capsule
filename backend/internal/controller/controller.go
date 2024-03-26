package controller

import (
	"net/http"

	"github.com/Anand-S23/capsule/internal/store"
)

type Controller struct {
    store      *store.Store
    production bool
}

func NewController(store *store.Store, production bool) *Controller {
    return &Controller {
        store: store,
        production: production,
    }
}

func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) error {
    return WriteJSON(w, http.StatusOK, "Pong")
}

