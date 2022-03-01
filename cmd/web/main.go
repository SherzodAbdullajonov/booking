package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SherzodAbdullajonov/booking/package/config"
	"github.com/SherzodAbdullajonov/booking/package/handlers"
	"github.com/SherzodAbdullajonov/booking/package/render"
	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig
var Session *scs.SessionManager

// main is the main function of the project
func main() {

	// changes this to true in Production
	app.InProduction = false
	Session = scs.New()
	Session.Lifetime = 24 * time.Hour
	Session.Cookie.Persist = true
	Session.Cookie.SameSite = http.SameSiteLaxMode
	Session.Cookie.Secure = app.InProduction
	app.Session = Session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache", err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting aplication on port %s", handlers.PortNumber))
	//_ = http.ListenAndServe(handlers.PortNumber, nil)
	srv := &http.Server{
		Addr:    handlers.PortNumber,
		Handler: Routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
