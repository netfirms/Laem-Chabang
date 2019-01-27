package users

import (
	httperror "github.com/portainer/libhttp/error"
	"github.com/netfirms/Laem-Chabang/api"
	"github.com/netfirms/Laem-Chabang/api/http/security"

	"net/http"

	"github.com/gorilla/mux"
)

func hideFields(user *portainer.User) {
	user.Password = ""
}

// Handler is the HTTP handler used to handle user operations.
type Handler struct {
	*mux.Router
	UserService            portainer.UserService
	TeamService            portainer.TeamService
	TeamMembershipService  portainer.TeamMembershipService
	ResourceControlService portainer.ResourceControlService
	CryptoService          portainer.CryptoService
	SettingsService        portainer.SettingsService
}

// NewHandler creates a handler to manage user operations.
func NewHandler(bouncer *security.RequestBouncer, rateLimiter *security.RateLimiter) *Handler {
	h := &Handler{
		Router: mux.NewRouter(),
	}
	h.Handle("/users",
		bouncer.RestrictedAccess(httperror.LoggerHandler(h.userCreate))).Methods(http.MethodPost)
	h.Handle("/users",
		bouncer.RestrictedAccess(httperror.LoggerHandler(h.userList))).Methods(http.MethodGet)
	h.Handle("/users/{id}",
		bouncer.AdministratorAccess(httperror.LoggerHandler(h.userInspect))).Methods(http.MethodGet)
	h.Handle("/users/{id}",
		bouncer.AuthenticatedAccess(httperror.LoggerHandler(h.userUpdate))).Methods(http.MethodPut)
	h.Handle("/users/{id}",
		bouncer.AdministratorAccess(httperror.LoggerHandler(h.userDelete))).Methods(http.MethodDelete)
	h.Handle("/users/{id}/memberships",
		bouncer.AuthenticatedAccess(httperror.LoggerHandler(h.userMemberships))).Methods(http.MethodGet)
	h.Handle("/users/{id}/passwd",
		rateLimiter.LimitAccess(bouncer.AuthenticatedAccess(httperror.LoggerHandler(h.userUpdatePassword)))).Methods(http.MethodPut)
	h.Handle("/users/admin/check",
		bouncer.PublicAccess(httperror.LoggerHandler(h.adminCheck))).Methods(http.MethodGet)
	h.Handle("/users/admin/init",
		bouncer.PublicAccess(httperror.LoggerHandler(h.adminInit))).Methods(http.MethodPost)

	return h
}
