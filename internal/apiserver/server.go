package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/cenkayla/userbalance/internal/db"
	"github.com/cenkayla/userbalance/internal/model"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	store  db.Store
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(store db.Store) *Server {
	s := &Server{
		router: mux.NewRouter(),
		store:  store,
	}
	s.configureRouter()
	return s
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/users/balance/{id}", s.GetUserBalance()).Methods("GET")
	s.router.HandleFunc("/users/balance/add/{id}", s.IncreaseUserBalance()).Methods("PUT")
	s.router.HandleFunc("/users/balance/reduce/{id}", s.ReduceUserBalance()).Methods("PUT")
	s.router.HandleFunc("/users/balance/transfer/{sender_id}", s.TransferUserBalance()).Methods("PUT")
}

func (s *Server) GetUserBalance() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		u := model.User{}
		u.ID, _ = strconv.Atoi(mux.Vars(r)["id"])

		ur := s.store.User()
		var err error
		u.Balance, err = ur.GetBalanceById(u.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&u)
	}
}

func (s *Server) IncreaseUserBalance() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		u := model.User{}
		balance := r.URL.Query().Get("balance")
		if balance == "" {
			err := errors.New("parameter balance is empty")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		u.ID, _ = strconv.Atoi(mux.Vars(r)["id"])
		u.Balance, _ = strconv.ParseFloat(balance, 64)

		ur := s.store.User()
		err := ur.IncreaseBalance(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("message:Balance was increased.")
	}
}

func (s *Server) ReduceUserBalance() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		u := model.User{}
		balance := r.URL.Query().Get("balance")
		if balance == "" {
			err := errors.New("parameter balance is empty")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		u.ID, _ = strconv.Atoi(mux.Vars(r)["id"])
		u.Balance, _ = strconv.ParseFloat(balance, 64)

		ur := s.store.User()
		err := ur.ReduceBalance(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("message:Balance was reduced.")
	}
}

func (s *Server) TransferUserBalance() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		sender := model.User{}
		receiver := model.User{}

		receiverId := r.URL.Query().Get("receiver_id")
		if receiverId == "" {
			err := errors.New("parameter receiver_id is empty")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		balance := r.URL.Query().Get("balance")
		if balance == "" {
			err := errors.New("parameter balance is empty")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sender.ID, _ = strconv.Atoi(mux.Vars(r)["sender_id"])
		sender.Balance, _ = strconv.ParseFloat(balance, 64)
		receiver.ID, _ = strconv.Atoi(receiverId)
		receiver.Balance, _ = strconv.ParseFloat(balance, 64)

		ur := s.store.User()
		err := ur.TransferBalance(sender, receiver)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("message:Balance was transfered.")
	}
}
