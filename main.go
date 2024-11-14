package main

import (
	"commander/command"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	commander := command.NewCommander()
	server := &http.Server{
		Addr:    ":8080",
		Handler: handleRequests(commander),
	}
	log.Fatal(server.ListenAndServe())
}

func handleRequests(cmdr command.Commander) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/execute", handleCommand(cmdr))
	return mux
}

func handleCommand(cmdr command.Commander) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, fmt.Errorf("method not allowed").Error(), http.StatusMethodNotAllowed)
			return
		}

		r.ParseForm()

		command := r.FormValue("command")
		if command == "get-system-info" {
			sysInfo, err := cmdr.GetSystemInfo()
			if err != nil {
				http.Error(w, fmt.Errorf("internal error occurred during GetSystemInfo command").Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(sysInfo)
		} else if command == "ping" {
			host := r.FormValue("host")
			if host == "" {
				http.Error(w, fmt.Errorf("'host' must be given in request'").Error(), http.StatusBadRequest)
				return
			}

			pingRes, err := cmdr.Ping(host)
			if err != nil {
				fmt.Println(err)
				http.Error(w, fmt.Errorf("internal error occurred during Ping command").Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pingRes)
		} else {
			http.Error(w, fmt.Errorf("unknown 'command' requested").Error(), http.StatusBadRequest)
			return
		}
	}
}
