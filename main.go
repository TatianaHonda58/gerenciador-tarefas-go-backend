package main

import (
	"encoding/json"
	"net/http"
)

// Estrutura simples de exemplo
type Resposta struct {
	Status  string `json:"status"`
	Mensagem string `json:"mensagem"`
}

func HandlerHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Resposta{Status: "OK", Mensagem: "API em Go funcionando!"})
}

func main() {
	http.HandleFunc("/health", HandlerHealthCheck)
	http.ListenAndServe(":8080", nil)
}