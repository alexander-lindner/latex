#!/bin/bash
OUT=/data/out
filename=""


if [ -f .latex ]; then
  echo "Config file found"
  filename=$(hocon -i .latex get texFile | tr -d '"' )
fi

if [[ "$1" == "watch" ]]; then
    BIBINPUTS=$OUT latexmk -quiet -outdir=$OUT -pdf -bibtex -shell-escape -interaction=nonstopmode -pvc -f "$filename"
else
    BIBINPUTS=$OUT latexmk -quiet -outdir=$OUT -pdf -bibtex -shell-escape -interaction=nonstopmode "$filename"
fi