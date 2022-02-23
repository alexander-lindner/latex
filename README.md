# Latex docker image ![GitHub release (latest by SemVer)](https://img.shields.io/github/downloads/alexander-lindner/latex/v2.0.0-7/total?style=flat-square)

This repo contains a basic dockerfile for a full latex environment, and a helper tool for simplifying
the building and setup of a latex project

Features:
* all dependencies
* biber / biblatex
* minted
* pdflatex / xalatex
* makeglossaries
* svg support

## Installation

> You need docker installed.

### using go
```bash
go get github.com/alexander-lindner/latex/textool
```

### using precompiled binaries

```bash
wget https://github.com/alexander-lindner/latex/releases/download/v2.0.0-7/textool-1.17.x -O textool
chmod +x textool
mv textool /usr/local/bin/
```