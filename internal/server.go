package internal

import "time"

type Server interface {
	Start(port int) error
	Shutdown(timeout time.Duration) error
}
