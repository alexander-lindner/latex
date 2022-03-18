package commands

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ledongthuc/pdf"
	"gopkg.in/yaml.v2"
)
import "github.com/google/go-github/v43/github"

type BibCommand struct {
	Url string `short:"u" long:"url" description:"The url of the web page you want to cite"`
}

var bibCommand BibCommand

func init() {
	_, err := parser.AddCommand("bib",
		"Adds a url to the bib file",
		"",
		&bibCommand,
	)
	if err != nil {
		log.Panic("Building the command parameter went wrong.", err)
	}
}

func (x *BibCommand) Execute(args []string) error {
	if len(options.Verbose) > 0 {
		log.SetLevel(log.TraceLevel)
		fmt.Println("Setting output to trace")
	}
	title := "Unknown"
	abstract := ""

	if strings.Contains(x.Url, "://github.com") {
		var re = regexp.MustCompile(`((git@|http(s)?:\/\/)([\w\.@]+)(\/|:))([\w,\-,\_]+)\/([\w,\-,\_]+)(.git){0,1}((\/){0,1})`)

		split := re.FindStringSubmatch(x.Url)
		t, a := renderGithub(split[6], split[7])

		title = t
		abstract = a
	} else {
		log.Println("Fetching url...")
		resp, err := http.Get(x.Url)
		if err != nil {
			log.Fatal("Couldn't fetch the content. Check the url and your internet connection.")
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Panic("An fatal error occurred during fetching the content of the provided url.")
			}
		}(resp.Body)
		content := resp.Body

		log.Debug("Received content type " + resp.Header.Get("Content-Type"))
		contentType := strings.Split(resp.Header.Get("Content-Type"), ";")[0]
		switch contentType {
		case "text/html":
			title, err = parseHtml(content)
			if err != nil {
				log.Fatal("Couldn't parse a html string. ", err)
			}
		case "application/pdf":
			readPdf(content)
		}
	}
	reg, err := regexp.Compile("[^a-zA-Z0-9_]+")
	if err != nil {
		log.Panic("Couldn't compile a regex string. ", err)
	}
	slug := strings.ReplaceAll(title, " ", "_")
	slug = reg.ReplaceAllString(slug, "")
	slug = strings.ReplaceAll(slug, "__", "_")
	currentTime := time.Now()
	fullDate := currentTime.Format("01-January") + "-" + strconv.Itoa(currentTime.Year())
	bibEntry := `
@online{ site:` + slug + `,
	title = { ` + title + ` },
	date = { ` + currentTime.Format("2006-01-02") + ` },
	url = { ` + x.Url + ` },
	abstract = { ` + abstract + ` },
	note = { [accessed ` + fullDate + `] }
}
`

	log.Info("This bibtex entry was successfully added to your main.bib file.")
	log.Info(bibEntry)
	bibliographyFile := options.Path + "/main.bib"
	f, err := os.OpenFile(bibliographyFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Couldn't open the file. ", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Panic("Couldn't close the file handler. ", err)
		}
	}(f)
	if _, err := f.WriteString(bibEntry); err != nil {
		log.Fatal("Couldn't write the entry to the file. ", err)
	}

	return nil
}
func parseHtml(content io.ReadCloser) (string, error) {

	doc, _ := html.Parse(content)
	node, err := getTitleFromBody(doc)
	if err != nil {
		return "", err
	}
	title := renderNode(node)

	return title, nil
}
func getTitleFromBody(doc *html.Node) (*html.Node, error) {
	var body *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "title" {
			body = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if body != nil {
		return body, nil
	}
	return nil, errors.New("missing <body> in the node tree")
}
func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	err := html.Render(w, n.FirstChild)
	if err != nil {
		return ""
	}
	return buf.String()
}

type T struct {
	Title      string `yaml:"title"`
	CffVersion string `yaml:"cff-version"`
	CffType    string `yaml:"type"`
	Authors    []struct {
		Names  string `yaml:"given-names"`
		Family string `yaml:"family-names"`
		Orcid  string `yaml:"orcid"`
	} `yaml:"authors"`
	Identifiers []struct {
		IdentifiersType string `yaml:"type"`
		Value           string `yaml:"value"`
		Description     string `yaml:"description"`
	} `yaml:"identifiers"`
	Code         string   `yaml:"repository-code"`
	Url          string   `yaml:"url"`
	Abstract     string   `yaml:"abstract"`
	Keywords     []string `yaml:"keywords"`
	License      string   `yaml:"license"`
	Version      string   `yaml:"version"`
	DateReleased string   `yaml:"date-released"`
}

func renderGithub(owner string, repoName string) (abstract string, slug string) {
	ctx := context.Background()
	client := github.NewClient(nil)

	get, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return
	}
	abstract = *get.Description
	contents, _, _, err := client.Repositories.GetContents(ctx, owner, repoName, "CITATION.cff", nil)
	if err != nil {
		return
	}
	slug = owner + " / " + repoName
	content, _ := contents.GetContent()

	t := T{}

	err = yaml.Unmarshal([]byte(content), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return t.Abstract, t.Title
}
func readPdf(content io.ReadCloser) (string, error) {
	path := options.Path + "/out/tmp.pdf"
	outFile, err := os.Create(path)
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			log.Panic("Couldn't create file. ", err)
		}
	}(outFile)
	_, err = io.Copy(outFile, content)
	if err != nil {
		log.Panic("Couldn't copy file content. ", err)
	}
	f, r, err := pdf.Open(path)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Panic("Couldn't close handler. ", err)
		}
	}(f)
	if err != nil {
		log.Fatal("Couldn't open pdf file. ", err)
	}
	fmt.Println(r.Page(0).V)
	return r.Outline().Title, nil
}
