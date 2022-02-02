package main

import (
	"fmt"
	"github.com/SherzodAbdullajonov/booking/package/config"
	"github.com/SherzodAbdullajonov/booking/package/handlers"
	"github.com/SherzodAbdullajonov/booking/package/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager

// main is the main function of the project
func main() {

	// changes this to true in Production
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
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
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
