package main

import (
	"encoding/json"
	"net"
	"net/http"

	"log"
)

func health(w http.ResponseWriter, r *http.Request) {
	log.Println("%s called /health", r.Host)
	w.Header().Set("Content-Type", "application/json")
	c, _ := json.Marshal(map[string]bool{"status": true})
	w.Write(c)
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("%s called /", r.Host)
	w.Header().Set("Content-Type", "text/plain")
	ifaces, _ := net.Interfaces()
	var ip net.IP
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
		}
	}
	b, _ := ip.MarshalText()
	w.Write(b)
}

func main() {
	log.Println("Micro started")
	http.HandleFunc("/", hello)
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8000", nil)
}
