package commands

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/alexander-lindner/go-cff"
	"github.com/ledongthuc/pdf"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
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
	author := ""

	if strings.Contains(x.Url, "://github.com") {
		var re = regexp.MustCompile(`((git@|http(s)?:\/\/)([\w\.@]+)(\/|:))([\w,\-,\_]+)\/([\w,\-,\_]+)(.git){0,1}((\/){0,1})`)

		split := re.FindStringSubmatch(x.Url)
		t, a, authors := renderGithub(split[6], split[7])

		title = t
		abstract = a
		author = strings.Join(authors, " and ")
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
			title, abstract, err = renderPdf(content)
			if len(title) <= 0 {
				title = x.Url
			}
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
	fullDate := currentTime.Format("02-January") + "-" + strconv.Itoa(currentTime.Year())
	bibEntry := `
@online{ site:` + slug + `,
	title = { ` + title + ` },
	date = { ` + currentTime.Format("2006-01-02") + ` },
	url = { ` + x.Url + ` },
	abstract = { ` + abstract + ` },
	note = { [accessed ` + fullDate + `] },
    author = { ` + author + ` }
}
`
	writeToBibFile(bibEntry)
	log.Info("This bibtex entry was successfully added to your main.bib file.")
	return nil
}

func writeToBibFile(bibEntry string) {
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

func renderGithub(owner string, repoName string) (title string, abstract string, authors []string) {
	ctx := context.Background()
	client := github.NewClient(nil)

	data, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return
	}
	abstract = data.GetDescription()

	contents, _, _, err := client.Repositories.GetContents(ctx, owner, repoName, "CITATION.cff", nil)
	if err != nil {
		return data.GetFullName(), abstract, authors
	}

	content, _ := contents.GetContent()

	fileContent, err := cff.Parse(content)
	if err != nil {
		return
	}
	for _, item := range fileContent.Authors {
		if item.IsPerson {
			authors = append(authors, item.Person.Family+", "+item.Person.GivenNames)
		}
	}
	return fileContent.Title, fileContent.Abstract, authors
}
func renderPdf(content io.ReadCloser) (title string, abstract string, err error) {
	path := copyPdfToTmpFile(content)
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

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		log.Fatal("Couldn't render pdf file. ", err)
	}
	_, err = buf.ReadFrom(b)
	if err != nil {
		log.Fatal("Couldn't render pdf file. ", err)
	}
	abstract = firstNCharacters(buf.String(), 255)

	return r.Outline().Title, abstract, nil
}
func firstNCharacters(s string, n int) string {
	r := []rune(s)
	if len(r) > n {
		return string(r[:n])
	}
	return s
}
func copyPdfToTmpFile(content io.ReadCloser) string {
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
	return path
}
