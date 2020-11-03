package conn

import (
	"log"
	"net/http"
	"path/filepath"
)

type HttpParams struct {
	Address string
	Prefix  string
	Root    string
}

// Init initializes the webserver and websocket connection
func Init(p *HttpParams) {
	var err error
	p.Root, err = filepath.Abs(p.Root)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("serving %s as %s on %s", p.Root, p.Prefix, p.Address)
	http.Handle(p.Prefix, http.StripPrefix(p.Prefix, http.FileServer(http.Dir(p.Root))))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.RemoteAddr + " " + r.Method + " " + r.URL.String())
		http.DefaultServeMux.ServeHTTP(w, r)
	})
	httpServer := http.Server{
		Addr:    p.Address,
		Handler: handler,
	}
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
