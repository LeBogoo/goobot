package commands

import (
	"github.com/bwmarrin/discordgo"
)

type BlaCommand struct {
}

func (c *BlaCommand) Name() string {
	return "bla"
}

func (c *BlaCommand) Description() string {
	return "Bla!"
}

// Execute is the function that is called when the command is executed

func (c *BlaCommand) Execute(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Bla!",
		},
	})
}
