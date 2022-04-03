package commands

import (
	"testing"
)

func TestCommandNormalUrl(t *testing.T) {
	options = Options{
		Path: "../test2",
	}
	cmd := BibCommand{Url: "https://en.wikipedia.org/wiki/Latex"}
	var args []string
	err := cmd.Execute(args)
	if err != nil {
		t.Errorf("Command executing went wrong.")
	}
	t.Log("Ok")
}
func TestCommandPdf(t *testing.T) {
	options = Options{
		Path: "../test2",
	}
	cmd := BibCommand{Url: "https://eprint.iacr.org/2016/086.pdf"}
	var args []string
	err := cmd.Execute(args)
	if err != nil {
		t.Errorf("Command executing went wrong.")
	}
	t.Log("Ok")
}
func TestCommandGithub(t *testing.T) {
	options = Options{
		Path: "../test2",
	}
	urls := []string{
		"https://github.com/intel/SGXDataCenterAttestationPrimitives",
		"https://github.com/intel/SGXDataCenterAttestationPrimitives/",
		"https://github.com/alexander-lindner/latex",
	}
	for _, url := range urls {
		t.Log("Testing " + url)
		cmd := BibCommand{Url: url}
		var args []string
		err := cmd.Execute(args)
		if err != nil {
			t.Errorf("Command executing went wrong.")
		}
	}

	t.Log("Finished Github tests")
}
