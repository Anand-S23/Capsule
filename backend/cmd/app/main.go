package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Anand-S23/capsule/internal/controller"
	"github.com/Anand-S23/capsule/internal/router"
	"github.com/Anand-S23/capsule/internal/store"
	"github.com/Anand-S23/capsule/pkg/config"
)

func main() {
    env, err := config.LoadEnv()
    if err != nil {
        log.Fatal(err)
    }

    db := store.InitDB(env.DB_URI, env.PRODUCTION)
    store := store.NewStore(
        store.NewPgUserRepo(db), 
        store.NewPgConnectionRepo(db),
        store.NewPgMeetingRepo(db),
        store.NewPgReminderRepo(db),
    )

    ctxTimeout := time.Second * 10
    ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
    defer cancel()

    controller := controller.NewController(
        ctx, store, env.JWT_SECRET, env.COOKIE_HASH_KEY, env.COOKIE_BLOCK_KEY, env.PRODUCTION)

    baseRouter := router.NewRouter(controller)
    router := router.NewCorsRouter(baseRouter, env.FE_URI)

    log.Println("Capsule backend running on port: ", env.PORT);
    http.ListenAndServe(":" + env.PORT, router)
}

