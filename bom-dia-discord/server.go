package main

import (
	"fmt"
	"net/http"
)

// Função handler para a rota "/"
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// Função handler para a rota "/mensagem"
func mensagemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esta é a rota /mensagem!")
}

func main() {
	// Associa as funções handler às rotas
	http.HandleFunc("/", handler)
	http.HandleFunc("/mensagem", mensagemHandler)

	// Inicia o servidor na porta 8080
	fmt.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}
