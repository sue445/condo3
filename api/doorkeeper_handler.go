package api

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/sue445/condo3/model"
	"github.com/sue445/condo3/model/doorkeeper"
	"net/http"
	"time"
)

// DoorkeeperHandler returns handler of /api/doorkeeper
func (h *Handler) DoorkeeperHandler(w http.ResponseWriter, r *http.Request) {
	h.performAPI(w, r, func(groupName string) (*model.Group, error) {
		span := sentry.StartSpan(r.Context(), "/api/doorkeeper/{group}", sentry.TransactionName(fmt.Sprintf("/api/doorkeeper/%s", groupName)))
		defer span.Finish()

		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetTag("group", groupName)
		})
		return doorkeeper.GetGroup(h.DoorkeeperAccessToken, groupName, time.Now())
	})
}
