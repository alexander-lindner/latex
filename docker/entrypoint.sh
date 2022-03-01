#!/usr/bin/env bash

OUT=/data/out
BIBINPUTS=$OUT latexmk -outdir=$OUT -pdf -shell-escape  -interaction=nonstopmode main.tex