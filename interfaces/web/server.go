package web

import (
	"context"
	"fmt"
	"github.com/chainpusher/chainpusher/interfaces/web/socket"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type ServerTask struct {
	server *http.Server
}

func (s *ServerTask) Start() error {
	logrus.Debugf("Starting server at %s", s.server.Addr)
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

	logrus.Debugf("Stopped server at %s", s.server.Addr)
	return err
}

func (s *ServerTask) Running() bool {
	return s.server != nil
}

func NewServerTask(host string, port int, processor MessageProcessor, clients *socket.Clients) *ServerTask {
	addr := fmt.Sprintf("%s:%d", host, port)

	socketHandler := &SocketHandler{clients, processor}
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", socketHandler.Handle)

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return &ServerTask{
		server: server,
	}
}
