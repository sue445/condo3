package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func apiDoorkeeperHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "group: %v, format: %v", vars["group"], vars["format"])
}
