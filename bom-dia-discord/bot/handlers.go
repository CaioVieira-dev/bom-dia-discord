package bot

import (
	"fmt"
	"regexp"
	"strings"
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

	if m.GuildID == "" {
		//mensagem enviada em DM
		handleDM(s, m)
	} else {
		//mensagem enviada no canal
		handleChannelMessages(s, m)
	}

}

func handleChannelMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Carregar a localização "America/Sao_Paulo"
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("Erro ao carregar fuso horário:", err)
		return
	}

	// Responder a mensagens específicas
	switch {
	case matchString("(?i)^bom[ ]{0,1}dia", m.Content):
		nome := m.Author.Username
		id := m.Author.ID

		mockBanco = append(mockBanco, Banco{
			nome:        nome,
			id:          id,
			goodMorning: time.Now().In(location),
		})

		fmt.Print(mockBanco)
		err := s.MessageReactionAdd(m.ChannelID, m.ID, "🌞")
		if err != nil {
			fmt.Println("Erro ao adicionar reação:", err)
		}
	case matchString("(?i)^encerrando", m.Content):
		id := m.Author.ID

		// Encontrar e editar a entrada no mockBanco com o mesmo id
		found := false
		for index, banco := range mockBanco {
			if banco.id == id && !banco.goodMorning.IsZero() && banco.closing.IsZero() {
				mockBanco[index].closing = time.Now().In(location)
				found = true
				break
			}
		}

		if found {
			fmt.Printf("Encontrado :)")
			fmt.Print(mockBanco)
			err := s.MessageReactionAdd(m.ChannelID, m.ID, "🌚")
			if err != nil {
				fmt.Println("Erro ao adicionar reação:", err)
			}
		} else {
			err := s.MessageReactionAdd(m.ChannelID, m.ID, "❔")
			if err != nil {
				fmt.Println("Erro ao adicionar reação:", err)
			}
		}
	default:
		fmt.Println("Comando não reconhecido.")
	}
}

func handleDM(s *discordgo.Session, m *discordgo.MessageCreate) {
	switch {
	case matchString("(?i)^!ponto", m.Content):
		// Carregar a localização "America/Sao_Paulo"
		location, err := time.LoadLocation("America/Sao_Paulo")
		if err != nil {
			fmt.Println("Erro ao carregar fuso horário:", err)
			return
		}

		id := m.Author.ID
		tabela := "💾 **Tabela de Registros Ponto**\n```\n"                                                  // Começa como um bloco de código
		tabela += fmt.Sprintf("%-20s | %-10s | %-10s | %-10s \n", "Nome", "Bom dia", "Encerrando", "Total") // Cabeçalho formatado
		tabela += fmt.Sprintf("%s\n", "---------------------+------------+------------+------------")       // Separador

		// Encontrar e editar a entrada no mockBanco com o mesmo id
		found := false
		for index, banco := range mockBanco {
			if banco.id == id {
				row := mockBanco[index]
				goodMorning := row.goodMorning
				closing := row.closing

				diff := closing.Sub(goodMorning)

				tabela += fmt.Sprintf("%-20s | %02d:%02d:%02d   | %02d:%02d:%02d   | ",
					row.nome,
					goodMorning.In(location).Hour(),
					goodMorning.In(location).Minute(),
					goodMorning.In(location).Second(),
					closing.In(location).Hour(),
					closing.In(location).Minute(),
					closing.In(location).Second(),
				)

				if closing.IsZero() {
					tabela += "-\n"
				} else {
					tabela += fmt.Sprintf("%v\n",
						diff)
				}
				found = true
			}
		}

		tabela += "```"

		if found {
			s.ChannelMessageSend(m.ChannelID, tabela)
		} else {

		}
	default:
		fmt.Println("Comando não reconhecido.")
	}

}

func matchString(pattern, s string) bool {
	trimmedS := strings.TrimSpace(s)
	matched, err := regexp.MatchString(pattern, trimmedS)
	if err != nil {
		fmt.Println("Erro ao compilar regex:", err)
		return false
	}
	return matched
}
