package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sue445/condo3/model/connpass"
	"net/http"
)

// ConnpassHandler returns handler of /api/conpass
func ConnpassHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	group, err := connpass.GetGroup(vars["group"])

	if err != nil {
		w.WriteHeader(errorStatusCode(err))
		fmt.Fprint(w, err)
		return
	}

	renderGroup(w, group, vars["format"])
}
