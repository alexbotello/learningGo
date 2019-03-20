package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	// Add a "addr" flag variable when running application ./simpleServer -addr=":3000" to run server on port 3000. Defaults to :8080
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse()

	// Initialize chat room and http handlers
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// Run the room in a seperate thread
	go r.run()

	// start the web server
	log.Println("Starting the web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
