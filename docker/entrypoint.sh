#!/bin/bash
OUT=/data/out
OPTIONS=''
filename=""

if [[ "$1" == "watch" ]]; then
  OPTIONS='-pvc -f'
fi

if [ -f .latex ]; then
  echo "Config file found"
  filename=$(hocon -i .latex get texFile | tr -d '"' )
fi

BIBINPUTS=$OUT latexmk -quiet -outdir=$OUT -pdf -bibtex -shell-escape  -interaction=nonstopmode $OPTIONS "$filename"
