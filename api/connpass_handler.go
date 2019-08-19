package api

import (
	"github.com/sue445/condo3/model"
	"github.com/sue445/condo3/model/connpass"
	"net/http"
	"time"
)

// ConnpassHandler returns handler of /api/conpass
func (h *Handler) ConnpassHandler(w http.ResponseWriter, r *http.Request) {
	performAPI(w, r, func(groupName string) (*model.Group, error) {
		return connpass.GetGroup(h.MemcachedConfig, groupName, time.Now())
	})
}
