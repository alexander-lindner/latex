#!/usr/bin/env bash

OUT=/data/out
BIBINPUTS=$OUT latexmk -outdir=$OUT -auxdir=$OUT/aux -pdf -shell-escape  -interaction=nonstopmode main.tex