---
label: textool run
icon: chevron-right
order: -20
---

Compiles the given project and outputs a PDF file as specified in the [config file](../config.md).  
## Overview

```shell
Usage:
  textool [OPTIONS] run

Compiles a given latex project by first building the docker image and then use it to compile the project.

Application Options:
  -v, --verbose   Verbose output
  -p, --path=     the name of the directory, which should be created

Help Options:
  -h, --help      Show this help message
```

## Additional information's

This command builds a docker image based on the given `Dockerfile` iff it was not already built.
It then starts this image, maps the current working project into it and runs `latexmk` to compile the PDF file.
If the compilation successfully finished, the container is removed and the final PDF file is moved to the working project and 
names as specified in the [config file](../config.md).  