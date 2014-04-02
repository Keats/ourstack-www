package ourstack

import (
	"html/template"
	"net/http"
)

var tmpl = make(map[string]*template.Template)

func init() {
	tmpl["index"] = template.Must(template.ParseFiles("templates/index.html", "templates/base.html"))

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// "/" matchs every requests, but we only want the homepage
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	data := GetCompanies()

	err := tmpl["index"].ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
