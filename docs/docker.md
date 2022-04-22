---
icon: container
order: -70
title: Docker image
---
We provide two docker images:
* [base](https://github.com/alexander-lindner/latex/blob/master/docker/Dockerfile):
  * `docker pull ghcr.io/alexander-lindner/latex:base` 
  * based on the latest `ubuntu` image
  * This image is used to build the other images.
  * It contains the base packages and the base configuration.
  * It is not meant to be used directly, instead [extend it](#custom-docker-image)
  * `~0.3 GB`
* [full](https://github.com/alexander-lindner/latex/blob/master/docker/Dockerfile.full):
  * `docker pull ghcr.io/alexander-lindner/latex:full` 
  * Contains **all** ctan.org latex packages
  * It is meant to be used directly
  * `~4.8GB`  

You can find the images on GitHub: [docker image](https://github.com/alexander-lindner/latex/pkgs/container/latex).
These images are internally used by TeXtool.

### Custom docker image


```Dockerfile
FROM ghcr.io/alexander-lindner/latex:base

RUN tlmgr install yourPackge && \
    rm -rf /usr/local/texlive/2021/texmf-dist/doc # to reduce the size of the image

RUN apt-get update && \
    apt-get install build-essential && \
    rm -rf /var/lib/apt/lists/* # to reduce the size of the image
```
### use the docker image without `textool`

Setup the latex project files by yourself.
Then use the image like below:
```bash
docker run --rm -ti \
           --volume $(pwd):/data \ # the main.tex file needs to be at /data/main.tex
           --user $(id -u):$(id -g) \ # so you can remove the tmp files in ./out/
           ghcr.io/alexander-lindner/latex:full \ 
           watch # remove this for a normal compilation
```