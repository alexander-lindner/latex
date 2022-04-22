---
icon: mark-github 
order: -60
title: Github Actions
---

We provide a simple github action ([latex-action](https://github.com/alexander-lindner/latex-action)), which compiles a TeXtool
project (aka the [run command](usage/run.md)).

You need to setup docker before by yourself.
## Options

More infos can be found [here](https://github.com/alexander-lindner/latex-action/blob/master/action.yml).
### Inputs

* path: path to the project (aka the `-p` parameter)
* version: the TeXtool version to use. Default: 'v2.1.5.1'
### Outputs

* pdf: path to the pdf file

## Example
```yaml
steps:
  - uses: actions/checkout@v3

  - id: setup-docker
    name: Set up Docker Buildx
    uses: docker/setup-buildx-action@v1
    
  - id: latex
    uses: alexander-lindner/latex-action@v1.6
    with:
      path: 'thesis'

  - name: pdf path
    run: echo ${​{ steps.latex.outputs.pdf }}
    shell: bash
```