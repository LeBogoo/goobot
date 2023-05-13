package commandsystem

import (
	"encoding/json"

	"github.com/bwmarrin/discordgo"
	"rogchap.com/v8go"
)

type JSCommand struct {
	Name      string
	JSContent string
	JSCtx     *v8go.Context
}

func NewJSCommand(vm *v8go.Isolate, name string, JSContent string) *JSCommand {
	jsCtx := v8go.NewContext(vm)

	jsCtx.RunScript("const TYPES = { SubCommand: 1, SubCommandGroup: 2, String: 3, Integer: 4, Boolean: 5, User: 6, Channel: 7, Role: 8, Mentionable: 9, Number: 10, Attachment: 11 }", "argumentTypes.js")

	jsCtx.RunScript(JSContent, "JSCommand.js")

	return &JSCommand{
		Name:      name,
		JSContent: JSContent,
		JSCtx:     jsCtx,
	}
}

func (c *JSCommand) Handler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if (i.Type != discordgo.InteractionApplicationCommand) || (i.ApplicationCommandData().Name == "") {
		return
	}

	if c.Name == i.ApplicationCommandData().Name {
		c.Execute(s, i)
	}
}

func (c *JSCommand) Execute(s *discordgo.Session, i *discordgo.InteractionCreate) {

	// convert [{"name":"value1","type":10,"value":15},{"name":"value2","type":10,"value":2}] to { value1: 15, value2: 2 }
	var options = i.ApplicationCommandData().Options
	var optionsMap = make(map[string]interface{})
	for _, option := range options {
		optionsMap[option.Name] = option.Value
	}

	// json stringify options
	json, _ := json.Marshal(optionsMap)

	res, _ := c.JSCtx.RunScript("execute("+string(json)+")", "execute.js")
	result := res.String()

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: result,
		},
	})
}

func (c *JSCommand) GetCommandData() *discordgo.ApplicationCommand {
	optionsRes, optionsErr := c.JSCtx.RunScript("JSON.stringify(options)", "options.js")
	optionsJson := "[]"
	if optionsErr == nil {
		optionsJson = optionsRes.String()
	}

	options := []*discordgo.ApplicationCommandOption{}
	json.Unmarshal([]byte(optionsJson), &options)

	descriptionRes, descriptionErr := c.JSCtx.RunScript("description", "description.js")
	if descriptionErr != nil {
		panic("Description not found.")
	}

	description := descriptionRes.String()

	return &discordgo.ApplicationCommand{
		Name:        c.Name,
		Description: description,
		Options:     options,
	}
}
