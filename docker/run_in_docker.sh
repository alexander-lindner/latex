#!/usr/bin/env bash

for OUTPUT in $(find . -type d)
do
  mkdir -p ../out/${OUTPUT}
done


pdflatex -output-directory=../out/ -shell-escape -interaction=nonstopmode main.tex
biber --output-directory=../out/ main
makeglossaries -d ../out/ main
pdflatex -output-directory=../out/ -shell-escape -interaction=nonstopmode main.tex
pdflatex -output-directory=../out/ -shell-e