package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sue445/condo3/api"
	"github.com/sue445/condo3/model"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	indexTmpl = readTemplate("index.html")
)

func main() {
	// Load variables
	kms := &Kms{KeyringKeyName: os.Getenv("KMS_KEYRING_KEY_NAME")}

	doorkeeperAccessToken, err := kms.GetFromEnvOrKms("DOORKEEPER_ACCESS_TOKEN", true)
	if err != nil {
		panic(err)
	}

	memcachedServer, err := kms.GetFromEnvOrKms("MEMCACHED_SERVER", true)
	if err != nil {
		panic(err)
	}

	memcachedUsername, err := kms.GetFromEnvOrKms("MEMCACHED_USERNAME", false)
	if err != nil {
		panic(err)
	}

	memcachedPassword, err := kms.GetFromEnvOrKms("MEMCACHED_PASSWORD", false)
	if err != nil {
		panic(err)
	}

	a := api.Handler{
		DoorkeeperAccessToken: doorkeeperAccessToken,
		MemcachedConfig: &model.MemcachedConfig{
			Server:   memcachedServer,
			Username: memcachedUsername,
			Password: memcachedPassword,
		},
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/connpass/{group}.{format}", a.ConnpassHandler)
	r.HandleFunc("/api/doorkeeper/{group}.{format}", a.DoorkeeperHandler)
	r.HandleFunc("/api/sandbox/{group}.{format}", a.SandboxHandler)
	r.HandleFunc("/", indexHandler)
	http.Handle("/", r)

	// Serve static files out of the public directory.
	// By configuring a static handler in app.yaml, App Engine serves all the
	// static content itself. As a result, the following two lines are in
	// effect for development only.
	public := http.StripPrefix("/public", http.FileServer(http.Dir("public")))
	http.Handle("/public/", public)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// indexHandler uses a template to create an index.html.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if os.Getenv("GAE_SERVICE") == "" {
		// Hot reloading for local
		indexTmpl = readTemplate("index.html")
	}

	vars := map[string]string{}

	if err := indexTmpl.Execute(w, vars); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func readTemplate(name string) *template.Template {
	return template.Must(template.ParseFiles(filepath.Join("templates", name)))
}
