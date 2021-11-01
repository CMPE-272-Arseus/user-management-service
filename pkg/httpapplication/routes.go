package httpapplication

import "net/http"

func (s *server) SetupRoutes() {
	s.router.HandleFunc("/{customer_id}/{user_id}", s.httpHandlerGetUserAccessLevel()).Methods(http.MethodGet)
	s.router.HandleFunc("/{customer_id}/{user_id}", s.httpHandlerPatchUserAccessLevel()).Methods(http.MethodPatch)
	s.router.HandleFunc("/{customer_id}/{user_id}", s.httpHandlerDeleteUserAccessLevel()).Methods(http.MethodDelete)

	s.router.HandleFunc("/{customer_id}", s.httpHandlerListUsersAccessLevels()).Methods(http.MethodGet)
	s.router.HandleFunc("/{customer_id}", s.httpHandlerPostUserAccessLevel()).Methods(http.MethodPost)
}
