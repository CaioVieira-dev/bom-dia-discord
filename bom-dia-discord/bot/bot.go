package bot

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

// NewBot cria uma nova sessão do Discord e retorna a sessão
func NewBot() (*discordgo.Session, error) {
	// Obtenha o token do bot a partir da variável de ambiente
	TOKEN := os.Getenv("DISCORD_BOT_TOKEN")
	if TOKEN == "" {
		return nil, fmt.Errorf("token do bot não encontrado; defina a variável de ambiente DISCORD_BOT_TOKEN")
	}

	// Crie uma nova sessão do Discord usando o token do bot
	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar sessão do discord: %w", err)
	}

	// Abra a conexão com o Discord
	err = dg.Open()
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com o Discord: %w", err)
	}
	// Adicione um handler para o evento de criação de mensagens
	dg.AddHandler(messageCreate)

	return dg, nil
}
