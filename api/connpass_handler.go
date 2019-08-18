package api

import (
	"github.com/gorilla/mux"
	"github.com/sue445/condo3/model/connpass"
	"net/http"
	"time"
)

// ConnpassHandler returns handler of /api/conpass
func (h *Handler) ConnpassHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	group, err := connpass.GetGroup(h.MemcachedConfig, vars["group"], time.Now())

	if err != nil {
		renderError(w, err)
		return
	}

	renderGroup(w, group, vars["format"])
}
