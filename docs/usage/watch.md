---
label: textool watch
icon: chevron-right
order: -30
---

Watches for changes in the project structure and re-compiles the project if needed.

## Overview
```shell
Usage:
  textool [OPTIONS] watch

Compiles a latex project and afterwards, watches for changes which triggers a recompilation

Application Options:
  -v, --verbose   Verbose output
  -p, --path=     the name of the directory, which should be created

Help Options:
  -h, --help      Show this help message
```

## Additional information's

This command will compile the project first and afterwards, it watches for changes in the project structure levering `latexmks` watch feature.

Use `[CTRL] + [C]` to stop watching.