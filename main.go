package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"goobot/commandsystem"
	"goobot/commandsystem/commands"
	"goobot/envparser"
)

func main() {
	envparser.ParseEnv()

	botToken := os.Getenv("BOT_TOKEN")

	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	user, _ := dg.User("@me")
	fmt.Printf("Bot is now running. Invite bot at https://discord.com/api/oauth2/authorize?client_id=%s&permissions=8&scope=bot\n", user.ID)

	registerCommands(dg)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func registerCommands(s *discordgo.Session) {
	commandHandler := commandsystem.NewCommandhandler(s)

	commandHandler.RegisterCommand(s, &commands.PingCommand{})
	commandHandler.RegisterCommand(s, &commands.BlaCommand{})

	s.AddHandler(commandHandler.HandleInteraction)
}
