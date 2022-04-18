---
label: textool init
icon: chevron-right
order: -10
---

The `textool init` command let you interactive create a latex project.
## Overview
```shell
Usage:
  textool [OPTIONS] init

Creates a directory and adds a minimal Latex template to this directory

Application Options:
  -v, --verbose   Verbose output
  -p, --path=     the name of the directory, which should be created

Help Options:
  -h, --help      Show this help message

```

## Additional information's

This command creates a `main.tex` file with additional demo content if chosen and a matching `main.bib`.
A `latexmkrc` file is created to configure `latexmk`.
All options that are chosen in this step are stored inside a [config file](../config.md) named `.latex`.
For customization a `Dockerfile` is created.