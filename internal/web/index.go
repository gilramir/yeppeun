package web

import (
	"log"
	"net/http"
)

func init() {
	// The index page handler. This is a bit special because it handles the index
	// page ("/") and any pages that don't match a registered route (serves as the
	// catch all handler).
	routes.register(route{
		Path: "/",
		Handler: logger(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "POST":
				err := templates.ExecuteTemplate(w, "pages/pretty.tmpl", map[string]interface{}{
					"Pretty": r.FormValue("dirty-json"),
				})

				if err != nil {
					log.Printf("web: error rendering pretty page; %s", err)
				}
			case "GET":
				if r.URL.Path != "/" {
					// NOTE: This is the catch all code path. We could do things here like
					// redirect old broken links, render a 404 page, etc.

					http.NotFound(w, r)
				} else {
					if err := templates.ExecuteTemplate(w, "pages/index.tmpl", nil); err != nil {
						log.Printf("web: error rendering index page; %s", err)
					}
				}
			default:
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			}
		}),
	})
}
