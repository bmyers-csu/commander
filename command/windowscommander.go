package command

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

type windowscommander struct{}

func (wc *windowscommander) GetSystemInfo() (SystemInfo, error) {
	// get hostname
	hostname, err := os.Hostname()
	if err != nil {
		return SystemInfo{}, err
	}

	// get local IP
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localIPAddress := conn.LocalAddr().(*net.UDPAddr).IP.String()

	return SystemInfo{
		Hostname:  hostname,
		IPAddress: localIPAddress,
	}, nil
}

func (wc *windowscommander) Ping(host string) (PingResult, error) {
	pingOutput, err := exec.Command("ping", host, "-n", "1").Output()

	if err != nil || strings.Contains(string(pingOutput), "Request timed out.") {
		return PingResult{Successful: false}, fmt.Errorf("error occurred while attempting to ping %s", host)
	}

	wordBreakDown := strings.Fields(strings.Split(string(pingOutput), "\n")[2])
	var timeInMs int
	fmt.Sscanf(wordBreakDown[4], "time=%dms", &timeInMs)

	return PingResult{Successful: true, Time: (time.Duration(timeInMs) * time.Millisecond)}, nil
}
