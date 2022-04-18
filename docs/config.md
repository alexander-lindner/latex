---
order: -50
icon: gear
---

TeXtool stores a config file under `.latex` in the [HOCON](https://github.com/lightbend/config/blob/main/HOCON.md#hocon-human-optimized-config-object-notation) file format.

## Config file

```css basic config file with explanation
##############################################
#            textools config file            #
# https://github.com/alexander-lindner/latex #
##############################################
fileName: alex.pdf # final file name of the PDF file
texFile: main.tex # entrypoint for the latex project
docker: { # modify docker options
  file: Dockerfile # Path to the Dockerfile, which is a docker standard to create custom images
}
features: { # options set to run the init command. Changing these has no influence as they are only used during project creation.
  documentclass: scrreprt
  glossary: true
  minted: true
  bibliography: true
  twocolumn: true
  lang: [english, english]
  examples: true
}
```