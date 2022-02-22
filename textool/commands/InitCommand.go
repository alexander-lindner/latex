package commands

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/alexander-lindner/latex/textool/helper"
	"github.com/go-akka/configuration"
	"github.com/valyala/fasttemplate"
	"os"
	"strconv"
	"strings"
)

type AddCommand struct {
}

var addCommand AddCommand

// the questions to ask
var qs = []*survey.Question{
	{
		Name: "Documentclass",
		Prompt: &survey.Select{
			Message: "Basic: Choose a document class:",
			Options: []string{"scrbook", "scrreprt", "scrartcl2", "scrartcl", "scrlttr2", "book", "report", "article", "letter", "exam"},
			Default: "scrreprt",
		},
		Validate: survey.Required,
	},
	{
		Name: "Lang",
		Prompt: &survey.MultiSelect{
			Message: "Basic: Choose a document class:",
			Options: []string{ // add showlanguages to \usepackage[english,showlanguages]{babel} to see all available langs
				"english",
				"american",
				"nohyphenation",
				"german",
				"ngerman",
				"afrikaans",
				"ancientgreek",
				"ibycus",
				"arabic",
				"armenian",
				"basque",
				"belarusian",
				"bulgarian",
				"catalan",
				"pinyin",
				"churchslavonic",
				"coptic",
				"croatian",
				"czech",
				"danish",
				"dutch",
				"ukenglish",
				"british",
				"UKenglish",
				"usenglishmax",
				"esperanto",
				"estonian",
				"ethiopic",
				"amharic",
				"geez",
				"farsi",
				"persian",
				"finnish",
				"schoolfinnish",
				"french",
				"patois",
				"francais",
				"friulan",
				"galician",
				"georgian",
				"german",
				"ngerman",
				"swissgerman",
				"greek",
				"polygreek",
				"monogreek",
				"hungarian",
				"icelandic",
				"assamese",
				"bengali",
				"gujarati",
				"hindi",
				"kannada",
				"malayalam",
				"marathi",
				"oriya",
				"pali",
				"panjabi",
				"tamil",
				"telugu",
				"indonesian",
				"interlingua",
				"irish",
				"italian",
				"kurmanji",
				"classiclatin",
				"latin",
				"liturgicallatin",
				"latvian",
				"lithuanian",
				"macedonian",
				"mongolian",
				"mongolianlmc",
				"bokmal",
				"norwegian",
				"norsk",
				"nynorsk",
				"occitan",
				"piedmontese",
				"polish",
				"portuguese",
				"portuges",
				"romanian",
				"romansh",
				"russian",
				"sanskrit",
				"serbian",
				"serbianc",
				"slovak",
				"slovenian",
				"slovene",
				"spanish",
				"espanol",
				"swedish",
				"thai",
				"turkish",
				"turkmen",
				"ukrainian",
				"uppersorbian",
				"welsh",
			},
		},
		Validate: survey.Required,
	},
	{
		Name:     "Twocolumn",
		Prompt:   &survey.Confirm{Message: "Basic: Use two column layout?", Help: "See https://texblog.org/2013/02/13/latex-documentclass-options-illustrated/"},
		Validate: survey.Required,
	},
	{
		Name:     "Minted",
		Prompt:   &survey.Confirm{Message: "Extra: Add a listing engine?", Help: "Minted is a advanced possibility to highlight all kinds of code"},
		Validate: survey.Required,
	},
	{
		Name:     "Glossary",
		Prompt:   &survey.Confirm{Message: "Extra: Use glossary feature?"},
		Validate: survey.Required,
	},
	{
		Name:     "Bibliography",
		Prompt:   &survey.Confirm{Message: "Extra: Use a bibliography engine?"},
		Validate: survey.Required,
	},
	{
		Name:     "Example",
		Prompt:   &survey.Confirm{Message: "Extra: add some demo content?"},
		Validate: survey.Required,
	},
}

func (x *AddCommand) Execute(args []string) error {
	helper.Exists(options.Path,
		func() {

		},
		func() {
			_ = os.MkdirAll(options.Path, 0700)
		},
	)

	mainConfig := options.Path + "/.latex"
	helper.Exists(mainConfig,
		func() {

		},
		func() {
			answers := struct {
				Documentclass string
				Twocolumn     bool
				Minted        bool
				Glossary      bool
				Bibliography  bool
				Example       bool
				Lang          []string
			}{}

			err := survey.Ask(qs, &answers)
			if err != nil {
				fmt.Println(err.Error())
				panic(err)
			}

			templateEngine := fasttemplate.New(configFile, "{{", "}}")
			configFileContent := templateEngine.ExecuteString(map[string]interface{}{
				"minted":        strconv.FormatBool(answers.Minted),
				"glossary":      strconv.FormatBool(answers.Glossary),
				"bibliography":  strconv.FormatBool(answers.Bibliography),
				"lang":          strings.Join(answers.Lang[:], ","),
				"twocolumn":     strconv.FormatBool(answers.Twocolumn),
				"documentclass": answers.Documentclass,
				"tag":           "full",
			})

			content := []byte(configFileContent)
			err = os.WriteFile(mainConfig, content, 0644)
			if err != nil {
				panic(err)
			}
		},
	)

	latexmkrcFile := options.Path + "/latexmkrc"
	helper.Exists(latexmkrcFile,
		func() {

		},
		func() {
			content := []byte(latexmkrc)
			_ = os.WriteFile(latexmkrcFile, content, 0644)
		},
	)
	c, err := os.ReadFile(mainConfig)
	config := configuration.ParseString(string(c))

	mainTex := options.Path + "/main.tex"
	helper.Exists(mainTex,
		func() {

		},
		func() {

			mapping := map[string]interface{}{}
			if config.GetBoolean("features.glossary", false) {
				mapping["glossary"] = glossariesTex
			}
			if config.GetBoolean("features.minted", false) {
				mapping["minted"] = mintedTex
			}
			if config.GetBoolean("features.bibliography", false) {
				mapping["bibliography"] = bibliographyTex
			}

			mapping["lang"] = strings.Join(config.GetStringList("features.lang"), ",")
			mapping["documentclass"] = config.GetString("features.documentclass")

			if config.GetBoolean("features.twocolumn", false) {
				mapping["twocolumn"] = "twocolumn"
			} else {
				mapping["twocolumn"] = "onecolumn"
			}
			templateEngine := fasttemplate.New(MinimalLatex, "<<", ">>")
			configFileContent := templateEngine.ExecuteString(mapping)

			content := []byte(configFileContent)
			err = os.WriteFile(mainTex, content, 0644)
			if err != nil {
				panic(err)
			}
		},
	)
	if config.GetBoolean("features.bibliography", false) {
		bibliographyFile := options.Path + "/main.bib"
		helper.Exists(bibliographyFile,
			func() {

			},
			func() {
				content := []byte(biberTex)
				_ = os.WriteFile(bibliographyFile, content, 0644)
			},
		)
	}
	return nil
}

func init() {
	_, err := parser.AddCommand("init",
		"Initialise a latex project directory",
		"Creates a directory and adds a minimal Latex template to this directory",
		&addCommand,
	)
	if err != nil {
		return
	}
}
