package commandsystem

import (
	"github.com/bwmarrin/discordgo"
)

type Command interface {
	Name() string
	Description() string
	Execute(s *discordgo.Session, i *discordgo.InteractionCreate)
}
