package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Função handler para o evento de criação de mensagens
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Log para depuração
	fmt.Printf("Mensagem recebida: '%s' de %s#%s (ID: %s) no canal %s\n", m.Content, m.Author.Username, m.Author.Discriminator, m.Author.ID, m.ChannelID)
	fmt.Printf("Tipo de mensagem: %d\n", m.Type)
	fmt.Printf("Guild ID: %s\n", m.GuildID)
	fmt.Printf("Channel ID: %s\n", m.ChannelID)
	fmt.Printf("Message ID: %s\n", m.ID)
	fmt.Printf("Timestamp: %s\n", m.Timestamp)
	fmt.Printf("Edited Timestamp: %s\n", m.EditedTimestamp)
	fmt.Printf("Content Raw: %s\n", m.ContentWithMentionsReplaced())

	// Ignorar mensagens do próprio bot
	if m.Author.ID == s.State.User.ID {
		fmt.Println("Mensagem do próprio bot ignorada.")
		return
	}

	// Ignorar mensagens que não são de texto comum
	if m.Type != discordgo.MessageTypeDefault {
		fmt.Println("Mensagem não é do tipo texto comum.")
		return
	}

	// Verificar se o conteúdo da mensagem está vazio
	if m.Content == "" {
		fmt.Println("Conteúdo da mensagem está vazio.")
		return
	}

	// Responder a mensagens específicas
	switch m.Content {
	case "Olá":
		fmt.Println("Respondendo à mensagem 'Olá'.")
		s.ChannelMessageSend(m.ChannelID, "Olá! Posso te ajudar com algo?")
	case "Ping":
		fmt.Println("Respondendo à mensagem 'Ping'.")
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	default:
		fmt.Println("Comando não reconhecido.")
	}
}
