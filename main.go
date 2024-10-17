package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/banking-system/internal/cache"
	"github.com/banking-system/internal/config"
	"github.com/banking-system/internal/controller"
	"github.com/banking-system/internal/repository"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	cfg := config.Init()
	repository := repository.NewRepository(cfg.DBConfig)
	cache := cache.NewCache()

	controller := controller.Init(repository, cache)
	router := mux.NewRouter()
	router.HandleFunc("/transfer", controller.Transfer).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:         cfg.ServerAddress,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      cors.Default().Handler(router),
	}

	go func() {
		log.Printf("Starting server at port http://%v\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	var wait time.Duration = time.Second * 10
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	repository.Close()
	cache.Close()
	srv.Shutdown(ctx)
	log.Println("shutting service down")
	os.Exit(0)
}
