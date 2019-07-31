package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sue445/condo3/model/doorkeeper"
	"net/http"
	"time"
)

// DoorkeeperHandler returns handler of /api/doorkeeper
func DoorkeeperHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	group, err := doorkeeper.GetGroup(vars["group"], time.Now())

	if err != nil {
		w.WriteHeader(errorStatusCode(err))
		fmt.Fprint(w, err)
		return
	}

	renderGroup(w, group, vars["format"])
}
