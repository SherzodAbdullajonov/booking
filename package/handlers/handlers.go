package handlers

import (
	"github.com/SherzodAbdullajonov/booking/package/config"
	"github.com/SherzodAbdullajonov/booking/package/modules"
	"github.com/SherzodAbdullajonov/booking/package/render"
	"net/http"
)

const PortNumber = ":8080"

//Repo is the repository used by handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo Creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIpp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_Ip", remoteIpp)

	render.RenderTemplates(w, "home.page.html", &modules.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//Perform some logic
	stringMap := make(map[string]string)
	stringMap["Test"] = "Hello again."

	remoteIpp := m.App.Session.GetString(r.Context(), "remote_Ip")
	stringMap["remote_Ip"] = remoteIpp
	render.RenderTemplates(w, "about.page.html", &modules.TemplateData{
		StringMap: stringMap,
	})
}
