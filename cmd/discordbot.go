/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/ei-sugimoto/gonews/internal/adapter"
	"github.com/spf13/cobra"
)

// discordbotCmd represents the discordbot command
var discordbotCmd = &cobra.Command{
	Use:   "discordbot",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(discordbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// discordbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// discordbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run() {
	discordClient, err := adapter.NewDiscordClient()
	if err != nil {
		log.Fatalf("failed to create discord client: %+v", err)
	}
	discordClient.Session.AddHandler(onMessageCreate)
	err = discordClient.Session.Open()
	if err != nil {
		log.Fatalf("failed to open discord session: %+v", err)
	}
	defer func() {
		discordClient.Session.Close()
		log.Println("discord session closed")
		os.Exit(0)
	}()
	fmt.Println("Bot is now running. Press CTRL+C to exit.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stop
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// 自分へのメンションが含まれているかを確認
	clientID, ok := os.LookupEnv("DISCORD_ID")
	u := m.Author
	if !ok {
		log.Println("DISCORD_ID is not set")
		return
	}
	botMention := "<@" + clientID + ">"
	if containsMention(m.Content, botMention) {

	} else {
		fmt.Printf("%20s %20s(%20s) > %s\n", m.ChannelID, u.Username, u.ID, m.Content)
	}

}

func sendMessage(s *discordgo.Session, channelID string, msg string) {
	_, err := s.ChannelMessageSend(channelID, msg)
	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}

func containsMention(content, mention string) bool {
	return strings.Contains(content, mention)
}
