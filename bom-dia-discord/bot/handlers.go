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

// Handler para Slash Commands
func slashCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case "bom-dia":
		nome := i.Member.User.Username
		id := i.Member.User.ID
		if nome == "" {
			nome = i.User.Username
		}
		if id == "" {
			id = i.User.ID
		}

		mockBanco = append(mockBanco, Banco{
			nome:        nome,
			id:          id,
			goodMorning: time.Now(),
		})

		fmt.Print(mockBanco)
		response := "Bom dia! Espero que você tenha um ótimo dia hoje! 🌞"
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: response,
			},
		})
	case "encerrando":
		id := i.Member.User.ID
		if id == "" {
			id = i.User.ID
		}

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
			response := "Até mais! Espero que você tenha um ótimo descanso! 🌙"
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: response,
				},
			})
		} else {
			response := "Por que você está encerrando antes de dar bom dia?? "
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: response,
				},
			})
		}

	case "ping":
		response := "Pong!"
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: response,
			},
		})
	default:
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Comando não reconhecido.",
			},
		})
	}
}

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
