![Banner](.github/banner.png)
Relaxe while TeXtool take care of your tex mess...

# TeXtool ![GitHub release (latest by SemVer)](https://img.shields.io/github/downloads/alexander-lindner/latex/v2.1.4/total?style=flat-square)

TeXtool is a small and powerful cli tool to create and manage latex projects.
It utilises docker and latexmk.
It supports all common latex features and tools like *minted*. 
TeXtools uses a [docker image](https://github.com/alexander-lindner/latex/pkgs/container/latex) as a foundation and lets you customise the whole build, 
e.g. you can add additional, non latex tools. 

Using Github actions? Use our action: [latex-action](https://github.com/alexander-lindner/latex-action)





## Use it

See our [documentation](https://textool.alindner.org/)


### use the docker image without `textool`

Setup the latex project files by yourself.
Then use the image like below:
```base
docker run --rm -ti \
           --volume $(pwd):/data \ # the main.tex file needs to be at /data/main.tex
           --user $(id -u):$(id -g) \ # so you can remove the temp file in ./out/
           ghcr.io/alexander-lindner/latex:full \ # use :base with for the minimal environment
           watch # remove this for a normal compilation
```

## ToDos

* config command
* compile without .latex & ./Dockerfile file
* customize main.tex file name
* Retype documentation
* finish github actions repo
* docker image name fix