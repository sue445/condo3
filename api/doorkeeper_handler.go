package api

import (
	"github.com/sue445/condo3/model"
	"github.com/sue445/condo3/model/doorkeeper"
	"net/http"
	"time"
)

// DoorkeeperHandler returns handler of /api/doorkeeper
func (h *Handler) DoorkeeperHandler(w http.ResponseWriter, r *http.Request) {
	h.performAPI(w, r, func(groupName string) (*model.Group, error) {
		return doorkeeper.GetGroup(h.DoorkeeperAccessToken, groupName, time.Now())
	})
}
