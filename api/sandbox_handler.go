package api

import (
	"github.com/sue445/condo3/model"
	"github.com/sue445/condo3/model/sandbox"
	"net/http"
)

// SandboxHandler returns handler of /api/sandbox
func (h *Handler) SandboxHandler(w http.ResponseWriter, r *http.Request) {
	performAPI(w, r, func(groupName string) (*model.Group, error) {
		return sandbox.GetGroup(groupName)
	})
}
