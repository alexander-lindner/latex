---
order: -20
icon: desktop-download
---

## Installation

!!!
You need docker installed and permission to use docker.
!!!

TeXtool is provided through Github's Release feature.
So, simply download the executable and put in the `PATH`.
+++ curl
```shell
curl -L -o /usr/local/bin/textool https://github.com/alexander-lindner/latex/releases/download/v2.1.4/textool-linux-amd64 && chmod  +x /usr/local/bin/textool
```
+++ wget
```shell
wget  https://github.com/alexander-lindner/latex/releases/download/v2.1.4/textool-linux-amd64 -O /usr/local/bin/textool && chmod  +x /usr/local/bin/textool
```
+++ dev
```shell
go install github.com/alexander-lindner/latex/textool@latest
```
+++


## Uninstallation

```shell
rm -f /usr/local/bin/textool
```