#!/usr/bin/env bash

OUT=/data/out

if [[ "$1" == "watch" ]]; then
  OPTIONS='-pvc -f  -pdflatex="pdflatex -synctex=1 -interaction=nonstopmode"'
else
  OPTIONS=''
fi
BIBINPUTS=$OUT latexmk -quiet -outdir=$OUT -pdf -bibtex -shell-escape  -interaction=nonstopmode $OPTIONS main.tex
