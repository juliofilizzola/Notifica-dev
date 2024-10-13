package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func InteractiveMessage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var githubUsernames = make(map[string]string)
	if i.Type == discordgo.InteractionApplicationCommand {
		if i.ApplicationCommandData().Name == "setgithub" {

			githubUsername := i.ApplicationCommandData().Options[0].StringValue()
			userID := i.Member.User.ID

			githubUsernames[userID] = githubUsername

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("Nome de usuário do GitHub '%s' foi salvo!", githubUsername),
				},
			})

			if err != nil {
				fmt.Println("Erro ao responder interação:", err)
			}
		}
	}
}

func ReadMessage(s *discordgo.Session, _ *discordgo.Ready) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "setgithub",
			Description: "Informe seu nome de usuário do GitHub",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "username",
					Description: "Seu nome de usuário do GitHub",
					Required:    true,
				},
			},
		},
	}

	for _, v := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Fatalf("Erro ao criar comando de slash: %v", err)
		}
	}
}
