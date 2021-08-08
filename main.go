package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
)

const version string = "gomac"

func newRouter() *mux.Router {
	r := mux.NewRouter()

	// Attach handlers
	r.HandleFunc("/", defaultHandler).Methods("GET")
	r.HandleFunc("/cpu/temperature", CPUTemperatureHandler).Methods("GET")
	r.HandleFunc("/gpu/temperature", GPUTemperatureHandler).Methods("GET")

	return r
}

func main() {
	r := newRouter()

	http.ListenAndServe(":8000", r)
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
