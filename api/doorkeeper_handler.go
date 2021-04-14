package api

import (
	"github.com/getsentry/sentry-go"
	"github.com/sue445/condo3/model"
	"github.com/sue445/condo3/model/doorkeeper"
	"net/http"
	"time"
)

// DoorkeeperHandler returns handler of /api/doorkeeper
func (h *Handler) DoorkeeperHandler(w http.ResponseWriter, r *http.Request) {
	h.performAPI(w, r, func(groupName string) (*model.Group, error) {
		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetTag("group", groupName)
		})
		return doorkeeper.GetGroup(h.DoorkeeperAccessToken, groupName, time.Now())
	})
}
