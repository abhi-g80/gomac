package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

const (
	version string = "gomac"
	port    string = ":8080"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()

	// Attach handlers
	r.HandleFunc("/", defaultHandler).Methods("GET")
	r.HandleFunc("/cpu/temperature", CPUTemperatureHandler).Methods("GET")
	r.HandleFunc("/gpu/temperature", GPUTemperatureHandler).Methods("GET")

	return r
}

func main() {
	l := log.New(os.Stdout, "gomac -> ", log.LstdFlags)
	r := newRouter()

	s := &http.Server{
		Addr:         port,
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	go func() {
		l.Println("Starting server on port", port)

		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	log.Printf("Received %s, gracefully shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()
	s.Shutdown(tc)
}

func getTemperature(object string) string {
	tcmd := "/usr/bin/powermetrics -s smc -i 1 | grep -m1 '%s.*temp' | awk '{print $4}'"
	cmd := fmt.Sprintf(tcmd, object)

	out, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	return string(out)
}

func getCPUTemp() string {
	return getTemperature("CPU")
}

func getGPUTemp() string {
	return getTemperature("GPU")
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", version)
}

func CPUTemperatureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", getCPUTemp())
}

func GPUTemperatureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", getGPUTemp())
}
