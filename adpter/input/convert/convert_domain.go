package convert

import (
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/app/domain"
	"github.com/juliofilizzola/bot_discord/config/env"
)

func ConvertDomainGithub(githubDomain *domain.Github) discordgo.WebhookParams {
	embed := &discordgo.MessageEmbed{
		URL:         githubDomain.PullRequest.HtmlUrl,
		Type:        discordgo.EmbedTypeLink,
		Title:       githubDomain.PullRequest.Title,
		Description: githubDomain.PullRequest.Body,
		Timestamp:   time.Now().Format(`2006-01-02 15:04:05`),
		Color:       0,
		Footer: &discordgo.MessageEmbedFooter{
			Text:         githubDomain.Organization.Login,
			IconURL:      githubDomain.Organization.AvatarUrl,
			ProxyIconURL: "",
		},
		Image: &discordgo.MessageEmbedImage{
			URL:      githubDomain.Organization.AvatarUrl,
			ProxyURL: "",
			Width:    280,
			Height:   20,
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL:      githubDomain.Organization.AvatarUrl,
			ProxyURL: "",
			Width:    280,
			Height:   20,
		},
		Author: &discordgo.MessageEmbedAuthor{
			URL:          githubDomain.Sender.HtmlUrl,
			Name:         githubDomain.Sender.Login,
			IconURL:      githubDomain.Sender.AvatarUrl,
			ProxyIconURL: "",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Status:",
				Value:  githubDomain.PullRequest.State,
				Inline: false,
			},
			{
				Name: "Assinado:",
				Value: func() string {
					if len(githubDomain.PullRequest.Assignee.Login) == 0 {
						return "Não teve assinatura"
					}
					return githubDomain.PullRequest.Assignee.Login
				}(),
				Inline: false,
			},
			{
				Name:   "Codigo adicionado:",
				Value:  strconv.Itoa(githubDomain.PullRequest.Additions),
				Inline: true,
			},
			{
				Name:   "Codigo deletado",
				Value:  strconv.Itoa(githubDomain.PullRequest.Deletions),
				Inline: true,
			},
			{
				Name:   "Commits:",
				Value:  "[commits](" + githubDomain.PullRequest.HtmlUrl + "/commits)",
				Inline: false,
			},
		},
	}

	if len(githubDomain.PullRequest.RequestedReviewers) > 0 {
		for _, value := range githubDomain.PullRequest.RequestedReviewers {
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   "Review :",
				Value:  value.Login,
				Inline: false,
			})
		}
	} else {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "Review :",
			Value:  "sem review...",
			Inline: false,
		})
	}

	return discordgo.WebhookParams{
		Content:    "Nova PR no Repositorio: " + githubDomain.Repository.Name,
		Username:   env.Username,
		AvatarURL:  env.AvatarURL,
		TTS:        false,
		Files:      nil,
		Components: nil,
		Embeds:     []*discordgo.MessageEmbed{embed},
		AllowedMentions: &discordgo.MessageAllowedMentions{
			Parse: []discordgo.AllowedMentionType{
				discordgo.AllowedMentionTypeEveryone,
			},
			Roles:       nil,
			Users:       nil,
			RepliedUser: false,
		},
		Flags: 0,
	}
}