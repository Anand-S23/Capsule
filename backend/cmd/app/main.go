package main

import (
	"log"
	"net/http"

	"github.com/Anand-S23/capsule/internal/config"
	"github.com/Anand-S23/capsule/internal/controller"
	"github.com/Anand-S23/capsule/internal/router"
)

func main() {
    env, err := config.LoadEnv()
    if err != nil {
        log.Fatal(err)
    }

    controller := controller.NewController(env.PRODUCTION)

    baseRouter := router.NewRouter(controller)
    router := router.NewCorsRouter(baseRouter, env.FRONTEND_URI)

    log.Println("Capsule backend running on port: ", env.PORT);
    http.ListenAndServe(":" + env.PORT, router)
}

