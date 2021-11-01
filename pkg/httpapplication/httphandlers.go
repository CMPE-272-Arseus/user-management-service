package httpapplication

import (
	"net/http"
)

// func (s *server) respond(w http.ResponseWriter, data interface{}, status int) {
// 	w.WriteHeader(status)

// 	if data != nil {
// 		response, _ := json.Marshal(data)
// 		_, _ = w.Write(response)
// 	}
// }

func (s *server) httpHandlerGetUserAccessLevel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) httpHandlerPatchUserAccessLevel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) httpHandlerDeleteUserAccessLevel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) httpHandlerPostUserAccessLevel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) httpHandlerListUsersAccessLevels() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// func (s *server) httpHandlerCoinChange() http.HandlerFunc {
// log := s.Logger()

// return func(w http.ResponseWriter, r *http.Request) {
// d, _ := r.URL.Query()["denominations"]

// di, err := helpers.Atoiarray(d)
// if err != nil {
// 	log.Info("httpHandlerCoinChange: invalid denominations in query params", zap.Strings("denominations", d))
// 	s.respond(w, nil, http.StatusBadRequest)
// 	return
// }

// t, _ := r.URL.Query()["total"]
// if len(t) < 1 {
// 	log.Info("httpHandlerCoinChange: no total given in query params")
// 	s.respond(w, nil, http.StatusBadRequest)
// 	return
// }

// ti, err := strconv.Atoi(t[0])
// if err != nil {
// 	log.Info("httpHandlerCoinChange: invalid total in query params", zap.Strings("total", t))
// 	s.respond(w, nil, http.StatusBadRequest)
// 	return
// }

// c := solving.NewProblemsCatalogue()
// p := c.NewCoinChangeProblem(di, ti)

// solution, err := p.Solve()
// if err != nil {
// 	log.Info("httpHandlerCoinChange: solve failed", zap.Error(err))
// 	s.respond(w, nil, http.StatusBadRequest)
// 	return
// }

// s.respond(w, nil, http.StatusOK)
// }
// }
