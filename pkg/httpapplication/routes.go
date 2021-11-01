package httpapplication

import "net/http"

func (s *server) SetupRoutes() {
	s.router.HandleFunc("/", s.httpHandlerCoinChange()).Methods(http.MethodGet)
}
