package web

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type ServerTask struct {
	server *http.Server
}

func (s *ServerTask) Start() error {
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (s *ServerTask) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)
	s.server = nil
	return err
}

func (s *ServerTask) Running() bool {
	return s.server != nil
}

func NewServerTask(host string, port int) *ServerTask {
	addr := fmt.Sprintf("%s:%d", host, port)
	server := &http.Server{Addr: addr}
	return &ServerTask{
		server: server,
	}
}
