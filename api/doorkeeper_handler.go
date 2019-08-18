package api

import (
	"github.com/gorilla/mux"
	"github.com/sue445/condo3/model/doorkeeper"
	"net/http"
	"time"
)

// DoorkeeperHandler returns handler of /api/doorkeeper
func (h *Handler) DoorkeeperHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	group, err := doorkeeper.GetGroup(h.DoorkeeperAccessToken, vars["group"], time.Now())

	if err != nil {
		renderError(w, err)
		return
	}

	renderGroup(w, group, vars["format"])
}
