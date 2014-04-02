package ourstack

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = make(map[string]*template.Template)

func init() {
	tmpl["home"] = template.Must(template.ParseFiles("templates/home.html", "templates/base.html"))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// "/" matchs every requests, but we only want the homepage
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	data := GetCompanies()

	err := tmpl["home"].ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
