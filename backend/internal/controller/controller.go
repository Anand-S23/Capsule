package controller

import (
	"net/http"

	"github.com/Anand-S23/capsule/internal/store"
	"github.com/gorilla/securecookie"
)

type Controller struct {
    store        *store.Store
    production   bool
    JwtSecretKey []byte
    CookieSecret *securecookie.SecureCookie
}

func NewController(store *store.Store, secretKey []byte, cookieHashKey []byte, cookieBlockKey []byte, production bool) *Controller {
    return &Controller {
        store: store,
        production: production,
        JwtSecretKey: secretKey,
        CookieSecret: securecookie.New(cookieHashKey, cookieBlockKey),
    }
}

func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) error {
    return WriteJSON(w, http.StatusOK, "Pong")
}

