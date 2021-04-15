package api

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gorilla/mux"
	"github.com/sue445/condo3/model"
	"github.com/sue445/condo3/model/connpass"
	"net/http"
	"time"
)

// ConnpassHandler returns handler of /api/conpass
func (h *Handler) ConnpassHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	span := sentry.StartSpan(r.Context(), "/api/connpass/{group}.{format}", sentry.TransactionName(fmt.Sprintf("/api/connpass/%s.%s", vars["group"], vars["format"])))
	defer span.Finish()

	h.performAPI(w, r, func(groupName string) (*model.Group, error) {
		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetTag("group", groupName)
		})
		return connpass.GetGroup(h.MemcachedConfig, groupName, time.Now())
	})
}
