package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"
	"time"
)

func (b *Bot) messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	if m.Type == discordgo.MessageTypeGuildMemberJoin {
		return
	}

	if m.Content == "!conselheiros" {
		_ = s.ChannelTyping(m.ChannelID)
		b.marketData.RWMutex.RLock()
		_, err := s.ChannelMessageSend(m.ChannelID, strconv.Itoa(b.hasBothCount))
		b.marketData.RWMutex.RUnlock()

		if err != nil {
			log.Error().Msgf("error while sending conselheiros message: %s", err.Error())
		}
	} else if strings.HasPrefix(m.Content, "!opensea") || strings.HasPrefix(m.Content, "!o") {
		words := strings.Split(m.Content, " ")

		switch words[1] {
		case "conselho", "c":
			_ = s.ChannelTyping(m.ChannelID)
			b.openseaData.RWMutex.RLock()
			msg := discordgo.MessageSend{
				Embeds: []*discordgo.MessageEmbed{constructOPEmbed(*b.conselhoOPStats)},
			}
			b.openseaData.RWMutex.RUnlock()
			_, err := s.ChannelMessageSendComplex(m.ChannelID, &msg)

			if err != nil {
				log.Error().Msgf("error while sending opensea message: %s", err.Error())
			}

		case "brasao", "Brasão", "b":
			_ = s.ChannelTyping(m.ChannelID)
			b.openseaData.RWMutex.RLock()
			msg := discordgo.MessageSend{
				Embeds: []*discordgo.MessageEmbed{constructOPEmbed(*b.brasaoOPStats)},
			}
			b.openseaData.RWMutex.RUnlock()
			_, err := s.ChannelMessageSendComplex(m.ChannelID, &msg)

			if err != nil {
				log.Error().Msgf("error while sending opensea message: %s", err.Error())
			}
		}
	} else if strings.HasPrefix(m.Content, "!vendas") || strings.HasPrefix(m.Content, "!v") {
		words := strings.Split(m.Content, " ")

		switch words[1] {
		case "conselho", "c":
			_ = s.ChannelTyping(m.ChannelID)
			b.marketData.RWMutex.RLock()
			msg := discordgo.MessageSend{
				Embeds: []*discordgo.MessageEmbed{constructMarketEmbed(*b.conselhoStats)},
			}
			b.marketData.RWMutex.RUnlock()
			_, err := s.ChannelMessageSendComplex(m.ChannelID, &msg)

			if err != nil {
				log.Error().Msgf("error while sending market message: %s", err.Error())
			}

		case "brasao", "Brasão", "b":
			_ = s.ChannelTyping(m.ChannelID)
			b.marketData.RWMutex.RLock()
			msg := discordgo.MessageSend{
				Embeds: []*discordgo.MessageEmbed{constructMarketEmbed(*b.brasaoStats)},
			}
			b.marketData.RWMutex.RUnlock()
			_, err := s.ChannelMessageSendComplex(m.ChannelID, &msg)

			if err != nil {
				log.Error().Msgf("error while sending market message: %s", err.Error())
			}
		}
	}
}

func (b *Bot) updateBotInfo() {
	ticker := *time.NewTicker(time.Second * 30)
	i := 1
	for {
		<-ticker.C
		b.marketData.RWMutex.RLock()
		b.openseaData.RWMutex.RLock()

		nicknames := [2]string{fmt.Sprintf("⟠%.3f #Conselho", b.conselhoOPStats.FloorPrice),
			fmt.Sprintf("⟠%.3f #Brasão", b.brasaoOPStats.FloorPrice)}
		sts := [2]string{fmt.Sprintf(" 1d: %.f%% %s 7d: %.f%% %s ", getPercentage(b.conselhoStats.sales[0].price, b.conselhoOPStats.OneDayAveragePrice),
			getArrow(getPercentage(b.conselhoStats.sales[0].price, b.conselhoOPStats.OneDayAveragePrice)), getPercentage(b.conselhoStats.sales[0].price, b.conselhoOPStats.SevenDayAveragePrice),
			getArrow(getPercentage(b.conselhoStats.sales[0].price, b.conselhoOPStats.SevenDayAveragePrice))),
			fmt.Sprintf("1d: %.f%% %s 7d: %.f%% %s ", b.brasaoStats.oneDayPercent, getArrow(b.brasaoStats.oneDayPercent), b.brasaoStats.sevenDaysPercent, getArrow(b.brasaoStats.sevenDaysPercent))}

		ts := [2]float64{getPercentage(b.conselhoStats.sales[0].price, b.conselhoOPStats.SevenDayAveragePrice), b.brasaoStats.sevenDaysPercent}

		gS, err := b.UserGuilds(0, "", "")

		if err != nil {
			log.Error().Msgf("error while getting bot guilds: %s", err.Error())
			break
		}

		mod := i % 2
		nickname := nicknames[mod]
		status := sts[mod]

		for _, g := range gS {
			err = b.Session.GuildMemberNickname(g.ID, "@me", nickname)

			if err != nil {
				log.Error().Msgf("error while changing nickname: %s, guild id : %s", err.Error(), g.ID)
				break
			}

		}

		idle := 0
		err = b.Session.UpdateStatusComplex(
			discordgo.UpdateStatusData{
				IdleSince: &idle,
				Activities: []*discordgo.Activity{
					{
						Name: status,
						Type: discordgo.ActivityTypeWatching,
						URL:  "",
					},
				},
				AFK:    false,
				Status: getStatus(ts[mod]),
			},
		)

		if err != nil {
			log.Error().Msgf("error while changing status: %s", err.Error())
			break
		}

		i++
		b.marketData.RWMutex.RUnlock()
		b.openseaData.RWMutex.RUnlock()
	}
}

func getArrow(v float64) string {

	if v >= 0 {
		return "⬈"
	}
	return "⬊"
}

func constructOPEmbed(OPstats openseaStats) *discordgo.MessageEmbed {
	Embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x000080,
		Description: fmt.Sprintf("``` FP: ⟠%.3f   Donos: %v```\u200b ", OPstats.FloorPrice, OPstats.NumOwners),

		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Preço médio 7d",
				Value:  fmt.Sprintf("%.3f", OPstats.SevenDayAveragePrice),
				Inline: true,
			},
			{
				Name:   "Vendas 7d",
				Value:  fmt.Sprintf("%v", OPstats.SevenDaySales),
				Inline: true,
			},
			{
				Name:   "Volume 7d",
				Value:  fmt.Sprintf("%.3f", OPstats.SevenDayVolume),
				Inline: true,
			},
			{
				Name:   "Preço médio",
				Value:  fmt.Sprintf("%.3f", OPstats.AveragePrice),
				Inline: true,
			},
			{
				Name:   "Vendas total",
				Value:  fmt.Sprintf("%v", OPstats.TotalSales),
				Inline: true,
			},
			{
				Name:   "Volume total",
				Value:  fmt.Sprintf("%.3f", OPstats.TotalVolume),
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{Text: strings.TrimRight(time.Now().Format(time.RFC1123Z), "+0200")},
	}

	return Embed
}

func constructMarketEmbed(stats marketStats) *discordgo.MessageEmbed {
	Embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x000080,
		Description: fmt.Sprintf("``` ultima venda: %.3f  ```\u200b ", stats.lastSale),

		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Mudanca 1d",
				Value:  fmt.Sprintf("%.f%%", stats.oneDayPercent),
				Inline: true,
			},
			{
				Name:   "Mudanca 7d",
				Value:  fmt.Sprintf("%.f%%", stats.sevenDaysPercent),
				Inline: true,
			},
			{
				Name:   "Mudanca 30d",
				Value:  fmt.Sprintf("%.f%%", stats.thirtyDaysPercent),
				Inline: true,
			},
			{
				Name:   "Vendas 1d",
				Value:  fmt.Sprintf("%v", stats.oneDaySales),
				Inline: true,
			},
			{
				Name:   "Vendas 7d",
				Value:  fmt.Sprintf("%v", stats.sevenDaysSales),
				Inline: true,
			},
			{
				Name:   "Vendas 30d",
				Value:  fmt.Sprintf("%v", stats.thirtyDaysSales),
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{Text: strings.TrimRight(time.Now().Format(time.RFC1123Z), "+0200")},
	}

	return Embed
}

func getPercentage(new, old float64) float64 {
	if old == 0 {
		return 0
	}
	return (new - old) / old * 100
}

func getStatus(f float64) string {
	if f >= 0 {
		return "online"
	}
	return "dnd"
}
