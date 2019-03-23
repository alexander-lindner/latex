# Latex docker image
[![MicroBadger Size](https://img.shields.io/microbadger/image-size/alexanderlindner/latex.svg?style=for-the-badge)](https://hub.docker.com/r/alexanderlindner/latex/)
[![Docker Build Status](https://img.shields.io/docker/build/alexanderlindner/latex.svg?style=for-the-badge)](https://hub.docker.com/r/alexanderlindner/latex/) 
[![MicroBadger Layers](https://img.shields.io/microbadger/layers/alexanderlindner/latex.svg?style=for-the-badge)](https://hub.docker.com/r/alexanderlindner/latex/) 
[![Docker Pulls](https://img.shields.io/docker/pulls/alexanderlindner/latex.svg?style=for-the-badge)](https://hub.docker.com/r/alexanderlindner/latex/)


This repo contains a basic dockerfile for a full latext environment and some examples how to use it.


## local

Assuming you've got the following structure:
```
. 
├── out
├── src
│    ├── main.tex
│    └── main.bib
├── docker-compose.yml
├── run.sh
└── run_in_docker.sh
```
You can use *docker-compose* for less configure work:
```yaml
version: '3.3'
services:
  build:
    image: alexanderlindner/latex:latest
    volumes:
      - ./:/data
    working_dir: /data/src
    command: "bash -c './run_in_docker.sh'"
    user: ${CURRENT_UID}

```

Place the `run_in_docker.sh` next to the compose file:
```bash
#!/usr/bin/env bash


for OUTPUT in $(find . -type d)
do
  mkdir -p ../out/${OUTPUT}
done


latex -output-directory=../out/ -shell-escape -interaction=nonstopmode main.tex
biber --output-directory=../out/ main
latex -output-directory=../out/ -shell-escape -interaction=nonstopmode main.tex
pdflatex -output-directory=../out/ -shell-escape -interaction=nonstopmode main.tex
```
and the run.sh too:
```bash
#!/usr/bin/env sh

export DATE=`date '+%Y-%m-%d_%H:%M:%S'`
export CURRENT_UID=$(id -u):$(id -g)
docker-compose up build
docker-compose down
```
Using `./run.sh` you will trigger the generation.

You will find the generated `main.pdf` in the `out` directory