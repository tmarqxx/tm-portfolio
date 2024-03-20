package handlers

import "net/http"

func serveDir(mux *http.ServeMux, pattern string, dir string) {
	mux.Handle(pattern, http.StripPrefix(pattern, http.FileServer(http.Dir(dir))))
}

func serveFile(mux *http.ServeMux, pattern string, filename string) {
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}
