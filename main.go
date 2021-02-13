//API
package main

import (
	"github.com/Belyakoff/goservice/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os" 
	"os/signal"
 	"log"
 	"time"
 	"context"	
) 


func main(){

	l  := log.New(os.Stdout, " product-api ", log.LstdFlags)

	
	ph := handlers.NewProducts(l)
	

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.ListAll)

  	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products", ph.Update)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.Create)
	postRouter.Use(ph.MiddlewareValidateProduct)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.Delete) 

	os.Setenv("PORT", "9090")
	port := os.Getenv("PORT")

	s := &http.Server{
		Addr: ":"+port,
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	go func(){
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
		 		l.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Received terminate, graceful shutdown:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}