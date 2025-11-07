package api

import "database/sql"

type Server struct {
	addr string
	db   *sql.DB
}

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}
func (s *Server) Run() error {
	return nil
}
