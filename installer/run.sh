#!/usr/bin/env bash

docker run --rm -ti --volume $(pwd):/data --user $(id -u):$(id -g) -w="/data/src" alexanderlindner/latex:latest bash -c '/run_in_docker.sh'