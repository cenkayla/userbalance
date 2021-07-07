package apiserver

import (
	"net/http"

	"github.com/cenkayla/userbalance/internal/db"
	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	store  db.Store
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(store db.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("users/balance/{id}", s.GetUserBalance()).Methods("GET")
	s.router.HandleFunc("users/balance/add/{id}", s.IncreaseUserBalance()).Methods("PUT")
	s.router.HandleFunc("users/balance/reduce/{id}", s.DecreaseUserBalance()).Methods("PUT")
	s.router.HandleFunc("users/balance/transfer/{sender_id}", s.TransferUserBalance()).Methods("PUT")
}
