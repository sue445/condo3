// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	indexTmpl = template.Must(
		template.ParseFiles(filepath.Join("templates", "index.html")),
	)
)

func main() {
	// Load variables
	kms := &Kms{KeyringKeyName: os.Getenv("KMS_KEYRING_KEY_NAME")}

	doorkeeperAccessToken, err := kms.GetFromEnvOrKms("DOORKEEPER_ACCESS_TOKEN")
	if err != nil {
		panic(err)
	}

	memcachedServer, err := kms.GetFromEnvOrKms("MEMCACHED_SERVER")
	if err != nil {
		panic(err)
	}

	memcachedUsername, err := kms.GetFromEnvOrKms("MEMCACHED_USERNAME")
	if err != nil {
		panic(err)
	}

	memcachedPassword, err := kms.GetFromEnvOrKms("MEMCACHED_PASSWORD")
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

	if err := indexTmpl.Execute(w, nil); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
