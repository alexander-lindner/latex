#!/bin/bash

docker build -t ghcr.io/alexander-lindner/latex:base ./docker --pull
docker build -t ghcr.io/alexander-lindner/latex:full ./docker/ -f ./docker/Dockerfile.full
