#!/bin/bash

OUT=/data/out

if [[ "$1" == "watch" ]]; then
  OPTIONS='-pvc -f'
else
  OPTIONS=''
fi

filename=$(hocon -i .latex get fileName | tr -d '"' )
filename=${filename%.pdf}
export OUTPUT_PDF_NAME=${filename##*/}

BIBINPUTS=$OUT latexmk -quiet -outdir=$OUT -pdf -bibtex -shell-escape  -interaction=nonstopmode $OPTIONS main.tex
