package handlers

import "net/http"

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/home", s.sayHello()).Methods("GET")
	s.router.HandleFunc("/sign-up", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/sign-in", s.handleSessionsCreate()).Methods("POST")
	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUser)
	private.HandleFunc("/whoami", s.handleWhoami()).Methods("GET")
}
