package commandsystem

import (
	"github.com/bwmarrin/discordgo"
)

type Commandhandler struct {
	commands []Command
}

func NewCommandhandler(s *discordgo.Session) *Commandhandler {
	user, _ := s.User("@me")
	guilds, _ := s.UserGuilds(100, "", "")

	for _, guild := range guilds {
		existingCommands, _ := s.ApplicationCommands(user.ID, guild.ID)
		for _, command := range existingCommands {
			s.ApplicationCommandDelete(user.ID, guild.ID, command.ID)
		}
	}

	return &Commandhandler{}
}

func (c *Commandhandler) RegisterCommand(s *discordgo.Session, command Command) {
	c.commands = append(c.commands, command)

	user, _ := s.User("@me")
	guilds, _ := s.UserGuilds(100, "", "")

	for _, guild := range guilds {
		s.ApplicationCommandCreate(user.ID, guild.ID, &discordgo.ApplicationCommand{
			Name:        command.Name(),
			Description: command.Description(),
		})
	}

}

func (c *Commandhandler) HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if (i.Type != discordgo.InteractionApplicationCommand) || (i.ApplicationCommandData().Name == "") {
		return
	}

	for _, command := range c.commands {
		if command.Name() == i.ApplicationCommandData().Name {
			command.Execute(s, i)
		}
	}
}
