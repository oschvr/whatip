package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mbndr/figlet4go"
)

func IPHandler(w http.ResponseWriter, r *http.Request) {

	// Local var to store ip addr
	var realIp string

	// Get IP from headers or Request obj

	// Attempt to get it from X-Forwarded-For header
	if r.Header.Get("X-Forwarded-For") != "" {
		realIp = r.Header.Get("X-Forwarded-For")
		log.Printf("🔵 [INFO] X-Forwarded-For %s \n", realIp)

		// Attempt X-Real-Ip header
	} else if r.Header.Get("X-Real-Ip") != "" {
		realIp = string(net.ParseIP(r.Header.Get("X-Real-Ip")))
		log.Printf("🔵 [INFO] X-Real-Ip %s \n", realIp)

		// Use r.RemoteAddr to get it if none of above worked
	} else {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Fatalln("🔴 [ERROR] Could not get Remote Address", err)
		}
		realIp = ip
	}

	// Print in console
	log.Printf("🔵 [INFO] Call from %s \n", realIp)

	// Write header and 200 response
	w.WriteHeader(http.StatusOK)

	//Write back IP address
	w.Write([]byte(realIp))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	//Console log info
	log.Println("🔵 [INFO] whatip is healthy ✅")

	//Write OK
	w.Write([]byte("OK"))
}

func main() {

	// Create router
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthHandler)
	r.HandleFunc("/", IPHandler)
	port := ":8080"

	// Figlet ASCII banner
	ascii := figlet4go.NewAsciiRender()
	renderOpts := figlet4go.NewRenderOptions()
	renderOpts.FontColor = []figlet4go.Color{
		figlet4go.ColorCyan,
	}
	renderStr, _ := ascii.RenderOpts("whatip", renderOpts)
	fmt.Print(renderStr)
	fmt.Println("----------------")

	// Write to console
	log.Printf("🔵 [INFO] whatip running on port %s\n", port)

	// Run server
	if err := http.ListenAndServe(port, r); err != nil {
		log.Panicln("🔴 [ERROR] Error starting server", err)
	}
}
