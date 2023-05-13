package commandsystem

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/bwmarrin/discordgo"
	"rogchap.com/v8go"
)

type CommandParser struct {
	commandFiles []string
}

func NewCommandParser() *CommandParser {
	return &CommandParser{
		commandFiles: []string{},
	}
}

func GetCommandFiles(folder string) []string {
	pwd, _ := os.Getwd()
	files, err := ioutil.ReadDir(path.Join(pwd, folder))

	if err != nil {
		panic(err)
	}

	filenames := []string{}
	for _, file := range files {
		filenames = append(filenames, path.Join(pwd, folder, file.Name()))
	}

	return filenames
}

func (cp *CommandParser) RegisterCommandFiles(dg *discordgo.Session) {
	vm := 	v8go.NewIsolate()

	commandFiles := GetCommandFiles("commands")

	cp.commandFiles = commandFiles
	for _, file := range commandFiles {

		commandFile, err := ReadCommandFile(file)

		if err != nil {
			panic(err)
		}

		command := NewJSCommand(vm, commandFile.Name, commandFile.Content)

		applicationCommand := command.GetCommandData()

		// loop through all guilds
		for _, guild := range dg.State.Guilds {
			dg.ApplicationCommandCreate(dg.State.User.ID, guild.ID, applicationCommand)
		}
		dg.AddHandler(command.Handler)
	}

}
