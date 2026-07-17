package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerHealthCheck(t *testing.T) {
	// Cria uma requisição simulada para o nosso endpoint
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Cria um gravador de resposta para capturar o resultado
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandlerHealthCheck)

	// Executa o teste
	handler.ServeHTTP(rr, req)

	// Validação (Assertion) do Status Code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("O handler retornou um status inesperado: recebido %v, esperado %v", status, http.StatusOK)
	}
}
