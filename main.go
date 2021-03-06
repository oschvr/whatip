package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
	"github.com/mbndr/figlet4go"
)

var (
	VERSION string
)

func IPHandler(w http.ResponseWriter, r *http.Request) {

	// Local var to store ip addr
	var realIp string

	// Get IP from headers or Request obj
	// Attempt to get it from X-Forwarded-For header
	if r.Header.Get("X-Forwarded-For") != "" {
		realIp = r.Header.Get("X-Forwarded-For")
		log.Printf("🔵 [INFO] x %s \n", realIp)

		// Attempt X-Real-Ip header
	} else if r.Header.Get("X-Real-Ip") != "" {
		realIp = r.Header.Get("X-Real-Ip")
		log.Printf("🔵 [INFO] X-Real-Ip %s \n", realIp)

		// Use r.RemoteAddr to get it if none of above worked
	} else {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		fmt.Printf("RemoteAddr: %s \n", ip)
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

func getVersion() string {
	e := exec.Command("make", "version")
	var out bytes.Buffer
	e.Stdout = &out
	err := e.Run()
	if err != nil {
		log.Fatalln("🔴 [ERROR] Could not version", err)
	}
	return out.String()
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
	fmt.Println(renderStr)

	// Print version
	if VERSION == "" {
		VERSION = "development"
	}
	fmt.Printf("Version: %s \n", VERSION)
	fmt.Println("----------------")

	// Write to console
	log.Printf("🔵 [INFO] whatip running on port %s\n", port)

	// Run server
	if err := http.ListenAndServe(port, r); err != nil {
		log.Panicln("🔴 [ERROR] Error starting server", err)
	}
}
