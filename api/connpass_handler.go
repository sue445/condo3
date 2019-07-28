package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sue445/condo3/model/connpass"
	"google.golang.org/appengine"
	"net/http"
)

// ConnpassHandler returns handler of /api/conpass
func ConnpassHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ctx := appengine.NewContext(r)

	group, err := connpass.GetGroup(ctx, vars["group"])

	if err != nil {
		w.WriteHeader(errorStatusCode(err))
		fmt.Fprint(w, err)
		return
	}

	renderGroup(w, group, vars["format"])
}
