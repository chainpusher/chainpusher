package monitor

import "github.com/chainpusher/blockchain/service"

type Code uint32

const (
	MaxRetries = iota
)

type Status struct {
	code Code
}

func (s *Status) GetCode() Code {
	return s.code
}

func (s *Status) Error() string {
	var message string
	switch s.GetCode() {
	case MaxRetries:
		message = "Max retries reached"
	default:
		message = "Unknown error"
	}
	return message
}

func IsMaxRetries(err error) bool {
	if err == nil {
		return false
	}
	status, ok := err.(*service.Status)
	if !ok {
		return false
	}
	cause := status.GetCause()
	causeStatus, ok := cause.(*Status)
	if !ok {
		return false
	}
	return causeStatus.GetCode() == MaxRetries
}

func NewWatcherError(code Code) error {
	watcherStatus := &Status{
		code: code,
	}
	return service.NewError(service.Other, watcherStatus)
}
