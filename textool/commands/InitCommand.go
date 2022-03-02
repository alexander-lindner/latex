package commands

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/alexander-lindner/latex/textool/helper"
	"github.com/go-akka/configuration"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasttemplate"
	"os"
	"strconv"
	"strings"
)

type AddCommand struct {
}

var addCommand AddCommand

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
		Name:     "FileName",
		Prompt:   &survey.Input{Message: "Basic: How should the file be named after generation? Default: main.pdf"},
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
		Name:     "Docker",
		Prompt:   &survey.Confirm{Message: "Extra: use custom docker build image(yes) or use the provided full image(no)?"},
		Validate: survey.Required,
	},
	{
		Name:     "Example",
		Prompt:   &survey.Confirm{Message: "Extra: add some demo content?"},
		Validate: survey.Required,
	},
}

func init() {
	_, err := parser.AddCommand("init",
		"Initialise a latex project directory",
		"Creates a directory and adds a minimal Latex template to this directory",
		&addCommand,
	)
	if err != nil {
		log.Panic("Building the command parameter went wrong.", err)
	}
}

func (x *AddCommand) Execute(args []string) error {
	if !helper.PathExists(options.Path) {
		err := os.MkdirAll(options.Path, 0700)
		if err != nil {
			log.Fatal("Couldn't create a directory")
		}
	}

	mainConfig := options.Path + "/.latex"
	if !helper.PathExists(mainConfig) {
		log.Println("Creating ./.latex as main config file for this tool.")
		answers := struct {
			Documentclass string
			Twocolumn     bool
			Minted        bool
			Glossary      bool
			Bibliography  bool
			Example       bool
			Lang          []string
			Docker        bool
			FileName      string
		}{}

		err := survey.Ask(qs, &answers)
		if err != nil {
			log.Fatal("A error occurred during `survey`.", err)
		}
		dockerImage := "ghcr.io/alexander-lindner/latex:full"
		if answers.Docker {
			dockerImage = "local"
		}
		templateEngine := fasttemplate.New(configFile, "<<", ">>")
		configFileContent := templateEngine.ExecuteString(map[string]interface{}{
			"minted":        strconv.FormatBool(answers.Minted),
			"glossary":      strconv.FormatBool(answers.Glossary),
			"bibliography":  strconv.FormatBool(answers.Bibliography),
			"lang":          strings.Join(answers.Lang[:], ","),
			"twocolumn":     strconv.FormatBool(answers.Twocolumn),
			"documentclass": answers.Documentclass,
			"dockerImage":   dockerImage,
			"fileName":      answers.FileName,
		})

		content := []byte(configFileContent)
		err = os.WriteFile(mainConfig, content, 0644)
		if err != nil {
			log.Fatal("Couldn't write config file.", err)
		}
	}

	latexmkrcFile := options.Path + "/latexmkrc"
	if !helper.PathExists(latexmkrcFile) {
		log.Println("Creating ./latexmkrc which configures latexmk.")
		content := []byte(latexmkrc)
		err := os.WriteFile(latexmkrcFile, content, 0644)
		if err != nil {
			log.Fatal("Couldn't write latexmkrc file.", err)
		}
	}
	c, err := os.ReadFile(mainConfig)
	if err != nil {
		log.Fatal("Couldn't read config file.", err)
	}
	config := configuration.ParseString(string(c))

	mainTex := options.Path + "/main.tex"
	if !helper.PathExists(mainTex) {
		log.Println("Creating main tex file main.tex")
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
		mapping["content"] = `(Type your content here.)`
		templateEngine := fasttemplate.New(MinimalLatex, "<<", ">>")
		configFileContent := templateEngine.ExecuteString(mapping)

		content := []byte(configFileContent)
		err = os.WriteFile(mainTex, content, 0644)
		if err != nil {
			log.Fatal("Couldn't write main tex file.", err)
		}
	}

	if config.GetBoolean("features.bibliography", false) {
		log.Println("Creating bibliography file main.bib")
		bibliographyFile := options.Path + "/main.bib"
		if !helper.PathExists(bibliographyFile) {
			content := []byte(biberTex)
			err := os.WriteFile(bibliographyFile, content, 0644)
			if err != nil {
				log.Fatal("Couldn't write main bibliography file.", err)
			}
		}
	}
	if config.GetString("docker.image") == "local" {
		log.Println("You've decided to customize the build process. Therefore, a Dockerfile was created.")
		dockerfile := options.Path + "/Dockerfile"
		if !helper.PathExists(dockerfile) {
			pkgs := []string{"koma-script", "xetex", "xstring", "float", "fontspec"}

			for _, lang := range config.GetStringList("features.lang") {
				if lang == "ngerman" {
					lang = "german"
				}
				pkgs = append(pkgs, "hyphen-"+lang, "babel-"+lang)
			}

			if config.GetBoolean("features.minted", false) {
				pkgs = append(pkgs, "soul", "listings", "minted", "fvextra", "fancyvrb", "upquote", "lineno", "xcolor", "catchfile", "framed", "etoolbox")
			}
			if config.GetBoolean("features.glossary", false) {
				pkgs = append(pkgs, "glossaries", "mfirstuc", "etoolbox", "textcase", "xfor", "datatool", "tracklang", "xkeyval")
				for _, lang := range config.GetStringList("features.lang") {
					if lang == "ngerman" {
						lang = "german"
					}
					pkgs = append(pkgs, "glossaries-"+lang)
				}

			}
			if config.GetBoolean("features.bibliography", false) {
				pkgs = append(pkgs, "csquotes", "biber", "biblatex")
			}
			fmt.Println(strings.Join(pkgs, " "))
			templateEngine := fasttemplate.New(MinimalDockerFile, "{{", "}}")
			finalContent := templateEngine.ExecuteString(map[string]interface{}{
				"packages": strings.Join(pkgs, " "),
			})

			content := []byte(finalContent)
			err := os.WriteFile(dockerfile, content, 0644)
			if err != nil {
				log.Fatal("Couldn't write Dockerfile file.", err)
			}
		}
	}
	return nil
}
