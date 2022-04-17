---
label: textool config
icon: chevron-right
order: -50
---

TBD
```shell
Usage:
  textool [OPTIONS] config [config-OPTIONS]

Set and get configuration values from the cli, specially from scripts. The first argument is the path to the configuration value, the second argument is the value to set if it exists. If no argument is provided, the whole
configuration is printed.

Application Options:
  -v, --verbose      Verbose output
  -p, --path=        the name of the directory, which should be created

Help Options:
  -h, --help         Show this help message

[config command options]

    Configuration for set command:
      -r, --replace  When working with an array, clear it and add the new values
```