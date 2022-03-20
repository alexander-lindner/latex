# Latex docker image ![GitHub release (latest by SemVer)](https://img.shields.io/github/downloads/alexander-lindner/latex/v2.1.1-3/total?style=flat-square)

This repo contains a cli tool `textool` and two [docker images](https://github.com/alexander-lindner/latex/pkgs/container/latex):

`textool` creates a latex project structure, compiles it, even auto compiles it after changes (aka watches) the project and adds more useful utils.


Docker images:
* a basic one with a minimal texlive installation
* a full one with texlive and all packages

The `:full` image takes care of all necessary features like:
* biber / biblatex
* minted
* pdflatex / xelatex / latexmk
* glossaries
* svg support

## Use it

> You need docker installed.
 
See [Usage](Usage.ipynb).

### With `textool`

Install it:


|                       using go                       |                                                                 using precompiled binaries                                                                  |
|:----------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------------------------------------------------------:|
| `go  get github.com/alexander-lindner/latex/textool ` | ` wget  https://github.com/alexander-lindner/latex/releases/download/v2.1.1-3/textool-amd64-linux -O textool && chmod  +x textool && mv  textool /usr/local/bin/` |

There are several commands available:
```
Usage:
  textool [OPTIONS] <command>

Application Options:
  -v, --verbose  Verbose output
  -p, --path=    the name of the directory, which should be created

Help Options:
  -h, --help     Show this help message

Available commands:
  bib    Adds a url to the .bib file
  clean  Cleans the working directory
  init   Initialise a latex project directory
  run    Compiles a latex project
  watch  Build and watches a latex project
```
#### Demo
[![asciicast](https://asciinema.org/a/475592.svg)](https://asciinema.org/a/475592)

### Manual

Setup the latex project files by yourself.
Then use the image like below:
```base
docker run --rm -ti \
           --volume $(pwd):/data \ # the main.tex file needs to be at /data/main.tex
           --user $(id -u):$(id -g) \ # so you can remove the temp file in ./out/
           ghcr.io/alexander-lindner/latex:full \ # use :base with for the minimal environment
           watch # remove this for a normal compilation
```