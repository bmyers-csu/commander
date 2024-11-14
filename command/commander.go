package command

import "time"

type Commander interface {
	Ping(host string) (PingResult, error)
	GetSystemInfo() (SystemInfo, error)
}

type PingResult struct {
	Successful bool
	Time       time.Duration
}

type SystemInfo struct {
	Hostname  string
	IPAddress string
}

func NewCommander() Commander {
	return &windowscommander{}
}
