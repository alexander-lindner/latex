---
label: textool config
icon: chevron-right
order: -50
---

A command to interact with the config file, mainly for scripts.

## Overview

```shell
Usage:
  textool [OPTIONS] config [config-OPTIONS] [config-PATH] [config-VALUE]

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
## Options
* PATH: a hocon path to select an element inside the config file.
Use `.` to separate different config keys.
* VALUE: if not set, the command will read and print the config value. 
If set, the given VALUE will be written to the given PATH
* -r, --replace:  When writing an array, clear it before and add the new values instead of appending.
See [example](#write-array-with-replace-options).


## Examples

###  Read
```json textool -p . config 
{
  fileName : demo.pdf
  texFile : main.tex
  docker : {
    file : Dockerfile
  }
  features : {
    documentclass : scrreprt
    glossary : true
    minted : true
    bibliography : true
    twocolumn : true
    lang : [ngerman]
    examples : true
  }
}
```

```json textool -p . config fileName
demo.pdf
```

```json textool -p . config features
{
  documentclass : scrreprt
  glossary : true
  minted : true
  bibliography : true
  twocolumn : true
  lang : [ngerman]
  examples : true
}
```

```json textool -p . config features.examples
true
```

```json textool -p . config features.lang
ngerman
```

### Write
```json textool -p . config features.lang english
{
  documentclass : scrreprt
  glossary : true
  minted : true
  bibliography : true
  twocolumn : true
  lang : [ngerman,english]
  examples : true
}
```
```json textool -p . config features.lang english
{
  documentclass : scrreprt
  glossary : true
  minted : true
  bibliography : true
  twocolumn : true
  lang : [ngerman,english,english]
  examples : true
}
```
#### Write array with replace options
```json textool -p . config -r features.lang english
{
  documentclass : scrreprt
  glossary : true
  minted : true
  bibliography : true
  twocolumn : true
  lang : [english]
  examples : true
}
```
## Additional information's

This command will not check for duplicates in an array.
When using the write feature, comments will be removed.