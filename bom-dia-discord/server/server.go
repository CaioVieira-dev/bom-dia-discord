package server

import (
	"fmt"
	"net/http"
)

// StartHTTPServer inicia o servidor HTTP
func StartHTTPServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Servidor HTTP está rodando!")
	})

	fmt.Println("Servidor HTTP está rodando na porta 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erro ao iniciar o servidor HTTP:", err)
	}
}
