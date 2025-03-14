package main

import (
	"fmt"
	"log"
	"net/http"
	"test-metrics-app/sysinfo"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		provider := sysinfo.NewResourcesProvider("/")

		totalMemory, err := provider.TotalMemory()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting total memory: %v", err), http.StatusInternalServerError)
			return
		}
		totalDisk, err := provider.TotalDisk()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting total disk: %v", err), http.StatusInternalServerError)
			return
		}
		cpuCores, err := provider.CPUCores()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting CPU cores: %v", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Total Memory: %d bytes\n", totalMemory)
		fmt.Fprintf(w, "Total Disk: %d bytes\n", totalDisk)
		fmt.Fprintf(w, "CPU Cores: %d\n", cpuCores)
	})

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
