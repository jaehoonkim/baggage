package command

import (
	"fmt"
	"strings"

	"github.com/jaehoonkim/baggage/parser"
	"github.com/mitchellh/cli"
)

type NewCommand struct {
	Ui     cli.Ui
	Parser *parser.Parser
}

func (c *NewCommand) Help() string {
	helpText := `
	Usage: Baggage new name
	`
	return strings.TrimSpace(helpText)
}

func (c *NewCommand) Run(args []string) int {
	c.Parser = parser.NewParser("./baggage.xml")
	c.Parser.Parse()
	if len(args) == 0 {
		fmt.Println(c.Help())
		return 0
	}

	name := args[0]
	if c.Contains(name) == true {
		fmt.Printf("%s 이 존재합니다.", name)
	} else {
		c.NewProject(name)
	}

	return 0
}

func (c *NewCommand) Contains(name string) bool {
	Bag := c.Parser.Bag
	for _, prj := range Bag.Projects {
		if prj.Name == name {
			return true
		}
	}
	return false
}

func (c *NewCommand) NewProject(name string) {
	c.Parser.Bag.AddProject(name)
	c.Parser.Save("./baggage.xml")
}

func (c *NewCommand) Synopsis() string {
	return "NewCommand::Synopsis"
}
