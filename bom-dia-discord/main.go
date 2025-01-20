package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"bom-dia-discord/bot"
	"bom-dia-discord/server"

	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
		return
	}

	// Inicializar o bot do Discord
	discordBot, err := bot.NewBot()
	if err != nil {
		fmt.Println("Erro ao inicializar o bot do Discord:", err)
		return
	}

	// Iniciar o servidor HTTP
	go server.StartHTTPServer()

	// Mantenha o bot rodando até que seja interrompido
	fmt.Println("Bot está online. Pressione CTRL+C para sair.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Feche a conexão com o Discord
	discordBot.Close()
}
