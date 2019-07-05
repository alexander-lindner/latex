# Latex docker image
[![MicroBadger Size](https://img.shields.io/microbadger/image-size/alexanderlindner/latex.svg?style=for-the-badge)](https://hub.docker.com/r/alexanderlindner/latex/)
[![Docker Build Status](https://img.shields.io/docker/build/alexanderlindner/latex.svg?style=for-the-badge)](https://hub.docker.com/r/alexanderlindner/latex/) 
[![MicroBadger Layers](https://img.shields.io/microbadger/layers/alexanderlindner/latex.svg?style=for-the-badge)](https://hub.docker.com/r/alexanderlindner/latex/) 
[![Docker Pulls](https://img.shields.io/docker/pulls/alexanderlindner/latex.svg?style=for-the-badge)](https://hub.docker.com/r/alexanderlindner/latex/)


This repo contains a basic dockerfile for a full latext environment and some examples how to use it.

Features:
* all dependencies
* biber / biblatex
* minted
* pdflatex / xalatex
* makeglossaries

## Installation
You need docker installed.

### Use it (automatic)
Install curl and execute:
```bash
bash <(curl -s https://raw.githubusercontent.com/alexander-lindner/latex/master/installer/installer.sh)
```
Using `./run.sh` you will trigger the generation.

You will find the generated `main.pdf` in the `out` directory

### Use it (manual): docker

Assuming you've got the following structure:
```
. 
├── out
├── src
│    ├── main.tex
│    └── main.bib
├── run.sh
└── run_in_docker.sh (optional)
```
add the `run.sh` and make it executable:
```bash
#!/usr/bin/env bash

#run.sh 

docker run --rm -dti --volume $(pwd):/data --user $(id -u):$(id -g) -w="/data/src" alexanderlindner/latex:latest bash -c '/run_in_docker.sh'
```
Using `./run.sh` you will trigger the generation.

You will find the generated `main.pdf` in the `out` directory

### Use it (manual): docker-compose

Assuming you've got the following structure:
```
. 
├── out
├── src
│    ├── main.tex
│    └── main.bib
├── docker-compose.yml
├── run.sh
└── run_in_docker.sh (optional)
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
    command: "bash -c '/run_in_docker.sh'"
    user: ${CURRENT_UID}

```

Place the `run.sh` next to the compose file:

```bash
#!/usr/bin/env sh

export CURRENT_UID=$(id -u):$(id -g)
docker-compose up build
docker-compose down
```
Using `./run.sh` you will trigger the generation.

You will find the generated `main.pdf` in the `out` directory

### Use it (manual): additional stuff

If you want to use `minted`, add the following line to your `main.tex`:
```latex
\usepackage[outputdir=../out]{minted}
```

If you want to override the generation bash file (`run_in_docker.sh`), add this file to your root path:
```bash
#!/usr/bin/env bash


for OUTPUT in $(find . -type d)
do
  mkdir -p ../out/${OUTPUT}
done


pdflatex -output-directory=../out/ -shell-escape -interaction=nonstopmode main.tex
biber --output-directory=../out/ main
makeglossaries -d ../out/ main
pdflatex -output-directory=../out/ -shell-escape -interaction=nonstopmode main.tex
pdflatex -output-directory=../out/ -shell-escape -interaction=nonstopmode main.tex
```
and change the compose path from ` "bash -c '/run_in_docker.sh'"` to ` "bash -c '../run_in_docker.sh'"`.
