package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/rakim0/femProject/internal/app"
	"github.com/rakim0/femProject/internal/routes"
)

func main() {
	app, err := app.NewApplication()

	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	if err != nil {
		panic(err)
	}
	defer app.DB.Close()
	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      r,
	}
	app.Logger.Printf("We are running on port %d", port)
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
