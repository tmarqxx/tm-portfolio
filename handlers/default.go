package handlers

import (
	"net/http"

	"tm-portfolio/views"

	"github.com/a-h/templ"
)

func NewDefaultMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", templ.Handler(views.IndexPage()))
	mux.Handle("/portfolio", templ.Handler(views.PortfolioPage()))
	mux.Handle("/contact", templ.Handler(views.ContactPage()))
	mux.HandleFunc("/menu", menuHandler)

	serveDir(mux, "/assets/", "assets")
	serveFile(mux, "/robots.txt", "public/robots.txt")

	return mux
}

func menuHandler(w http.ResponseWriter, r *http.Request) {
	views.MenuPage(r.Header.Get("Referer")).Render(r.Context(), w)
}
