package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sue445/condo3/model/sandbox"
	"net/http"
)

// SandboxHandler returns handler of /api/sandbox
func (h *Handler) SandboxHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	group, err := sandbox.GetGroup(vars["group"])

	if err != nil {
		w.WriteHeader(errorStatusCode(err))
		fmt.Fprint(w, err)
		return
	}

	renderGroup(w, group, vars["format"])
}
