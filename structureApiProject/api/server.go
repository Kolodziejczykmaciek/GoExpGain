package api

import (
	"GoExpGain/structureApiProject/storage"
	"encoding/json"
	"net/http"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}
func (s *Server) Start() error {
	http.HandleFunc("/user", s.handleGetUserByID)
	return http.ListenAndServe(s.listenAddr, nil)
}

//attach some handlers

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(10)

	json.NewEncoder(w).Encode(user)
}
