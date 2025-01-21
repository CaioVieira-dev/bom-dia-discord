package bot

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

// Registrar comandos de barra
func RegisterSlashCommands(s *discordgo.Session) error {
	if s == nil {
		return fmt.Errorf("a sessão do Discord não foi inicializada")
	}

	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "bom-dia",
			Description: "Responde com uma mensagem amigável de 'Bom dia!'",
		},
		{
			Name:        "encerrando",
			Description: "Responde com uma mensagem amigável de 'Bom descanso!'",
		},
		{
			Name:        "ping",
			Description: "Responde com Pong!",
		},
	}

	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			return fmt.Errorf("não foi possível criar o comando %s: %w", cmd.Name, err)
		}
	}

	fmt.Println("Slash commands registrados com sucesso!")
	return nil
}

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

	// Registre os comandos de barra (slash commands)
	err = RegisterSlashCommands(dg)
	if err != nil {
		return nil, fmt.Errorf("erro ao registrar comandos de barra: %w", err)
	}
	// Adicione o handler para interações de comandos
	dg.AddHandler(slashCommandHandler)

	return dg, nil
}
