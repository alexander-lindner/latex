package commands

import (
	"fmt"
	"github.com/alexander-lindner/latex/textool/helper"
	"github.com/go-akka/configuration/hocon"
	log "github.com/sirupsen/logrus"
	"strings"
)

type ConfigCommand struct {
}

var configCommand ConfigCommand

type ConfigSetCommand struct {
	Replace bool `short:"r" long:"replace" description:"When working with an array, clear it and add the new values"`
}

var configSetCommand ConfigSetCommand

type CustomElement struct {
	string string
}

func (t CustomElement) GetString() string {
	return t.string
}
func (t CustomElement) IsString() bool {
	return true
}
func (t CustomElement) IsArray() bool {
	return false
}
func (t CustomElement) GetArray() []*hocon.HoconValue {
	return make([]*hocon.HoconValue, 0)
}

type CustomArrayElement struct {
	el []*hocon.HoconValue
}

func (t CustomArrayElement) GetString() string {
	return ""
}
func (t CustomArrayElement) IsString() bool {
	return false
}
func (t CustomArrayElement) IsArray() bool {
	return true
}
func (t CustomArrayElement) GetArray() []*hocon.HoconValue {
	return t.el
}

func (x *ConfigCommand) Execute(args []string) error {
	mainConfig := options.Path + "/.latex"
	if !helper.PathExists(mainConfig) {
		log.Error("Config file not found")
		return nil
	}
	log.SetLevel(log.FatalLevel)
	config := helper.GetConfig(options.Path)

	if len(args) == 0 {
		fmt.Print(config.Root().String())
		return nil
	}
	configPath := args[0]
	if len(args) == 1 {
		if config.IsArray(configPath) {
			arr := config.GetStringList(configPath)
			fmt.Print(strings.Join(arr, ","))
		} else if config.IsObject(configPath) {
			fmt.Print(config.GetConfig(configPath).String())
		} else {
			fmt.Print(config.GetString(configPath))
		}
	} else {
		if config.IsArray(configPath) {
			currentElement := config.GetConfig(configPath).Root()
			currentArray := currentElement.GetArray()

			if configSetCommand.Replace {
				currentArray = make([]*hocon.HoconValue, 0)
			}
			for _, arg := range args[1:] {
				newValue := hocon.NewHoconValue()
				newValue.NewValue(CustomElement{arg})
				currentArray = append(currentArray, newValue)
			}

			currentElement.NewValue(CustomArrayElement{currentArray})
		} else {
			config.GetConfig(configPath).Root().NewValue(
				CustomElement{args[1]},
			)
		}
		helper.WriteConfig(config, options.Path)
	}

	return nil
}

func init() {
	g, err := parser.AddCommand("config",
		"Interact with the configuration",
		"Set and get configuration values from the cli, specially from scripts. "+
			"The first argument is the path to the configuration value, the second argument is the value to set if it exists. "+
			"If no argument is provided, the whole configuration is printed.",
		&configCommand,
	)
	if err != nil {
		log.Panic("Building the command parameter went wrong.", err)
	}
	_, err = g.AddGroup("Configuration for set command", "f", &configSetCommand)
	if err != nil {
		log.Panic("Building the command parameter went wrong.", err)
	}
}
