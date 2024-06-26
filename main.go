package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"meu-api/internal/bluetooth"
)

func main() {
	go bluetooth.StartBluetoothScanner()

	http.HandleFunc("/api/devices", handleDevices)

	port := ":8080"
	fmt.Printf("Servidor iniciado em http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleDevices(w http.ResponseWriter, r *http.Request) {
	devices := bluetooth.GetDevices()

	jsonBytes, err := json.Marshal(devices)
	if err != nil {
		http.Error(w, "Erro ao serializar dispositivos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
