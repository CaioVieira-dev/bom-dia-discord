package bot

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Banco struct {
	nome        string
	id          string
	goodMorning time.Time
	closing     time.Time
}

var mockBanco = []Banco{}

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
	case "bom dia":
		nome := m.Author.Username
		id := m.Author.ID

		mockBanco = append(mockBanco, Banco{
			nome:        nome,
			id:          id,
			goodMorning: time.Now(),
		})

		fmt.Print(mockBanco)
		s.ChannelMessageSend(m.ChannelID, "Bom dia! Espero que você tenha um ótimo dia hoje! 🌞")
	case "encerrando":
		id := m.Author.ID

		// Encontrar e editar a entrada no mockBanco com o mesmo id
		found := false
		for index, banco := range mockBanco {
			if banco.id == id && !banco.goodMorning.IsZero() && banco.closing.IsZero() {
				mockBanco[index].closing = time.Now()
				found = true
				break
			}
		}

		if found {
			fmt.Printf("Encontrado :)")
			fmt.Print(mockBanco)

			s.ChannelMessageSend(m.ChannelID, "Até mais! Espero que você tenha um ótimo descanso! 🌙")
		} else {
			s.ChannelMessageSend(m.ChannelID, "Por que você está encerrando antes de dar bom dia?? ")
		}
	default:
		fmt.Println("Comando não reconhecido.")
	}
}
