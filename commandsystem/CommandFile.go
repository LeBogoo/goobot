package commandsystem

import (
	"strings"

	"goobot/utils"
)

type CommandFile struct {
	Name    string
	Content string
}

func ReadCommandFile(filepath string) (CommandFile, error) {
	content, err := utils.ReadFile(filepath)
	if err != nil {
		return CommandFile{}, err
	}

	parts := strings.Split(filepath, "/")
	filename := parts[len(parts)-1]
	filename = strings.Split(filename, ".")[0]

	return CommandFile{
		Name:    filename,
		Content: content,
	}, nil
}
