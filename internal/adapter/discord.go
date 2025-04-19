package adapter

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/xerrors"
)

type discordClient struct {
	Session *discordgo.Session
	ID      string
}

func NewDiscordClient() (*discordClient, error) {
	token, ok := os.LookupEnv("DISCORD_TOKEN")
	if !ok {
		return nil, xerrors.New("DISCORD_TOKEN is not set")
	}
	id, ok := os.LookupEnv("DISCORD_ID")
	if !ok {
		return nil, xerrors.New("DISCORD_ID is not set")
	}
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, xerrors.New(err.Error())
	}

	return &discordClient{
		Session: discord,
		ID:      id,
	}, nil
}
