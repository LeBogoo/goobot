package envparser

import (
	"os"
	"strings"

	"goobot/utils"
)


func ParseEnv() {
	content, err := utils.ReadFile(".env")
	if err != nil {
		return
	}
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		parts := strings.Split(line, "=")
		os.Setenv(parts[0], parts[1])	
	}
}