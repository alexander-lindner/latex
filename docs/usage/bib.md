---
label: textool bib
icon: chevron-right
order: -60
---

Quickly add a bibliography entry to the bib file from the given url.

## Overview

```shell
Usage:
  textool [OPTIONS] bib [bib-OPTIONS]

Application Options:
  -v, --verbose   Verbose output
  -p, --path=     the name of the directory, which should be created

Help Options:
  -h, --help      Show this help message

[bib command options]
      -u, --url=  The url of the web page you want to cite
```

## Options

* -u, --url:  any kind of url which should be added

## examples

```ini textool -p . bib -u https://en.wikipedia.org/wiki/Latex
@online{ site:Latex_Wikipedia,
	title = { Latex - Wikipedia },
	date = { 2022-04-15 },
	url = { https://en.wikipedia.org/wiki/Latex },
	abstract = {  },
	note = { [accessed 15-April-2022] },
    author = {  }
}
```

```ini textool -p . bib -u https://eprint.iacr.org/2016/086.pdf 
@online{ site:httpseprintiacrorg2016086pdf,
	title = { https://eprint.iacr.org/2016/086.pdf },
	date = { 2022-04-15 },
	url = { https://eprint.iacr.org/2016/086.pdf },
	abstract = { IntelSGXExplainedVictorCostanandSrinivasDevadasvictor@costan.us,devadas@mit.eduComputerScienceandArtiﬁcialIntelligenceLaboratoryMassachusettsInstituteofTechnologyABSTRACTIntel’sSoftwareGuardExtensions(SGX)isasetofextensionstotheIntelarchitecturethataimsto },
	note = { [accessed 15-April-2022] },
    author = {  }
}
```

```ini textool -p . bib -u https://github.com/alexander-lindner/latex
@online{ site:alindners_latex_collection,
	title = { alindner's latex collection },
	date = { 2022-04-15 },
	url = { https://github.com/alexander-lindner/latex },
	abstract = { A collection of tools to provide easy latex support },
	note = { [accessed 15-April-2022] },
    author = { Lindner, Alexander }
}
```
## Additional information's
Supports the [Citation File Format](https://citation-file-format.github.io/) in Github repos.

This should only be used as a quick way of adding a file to the bib entry.
In a lot of cases, this command will do nonsens.
Therefore, we recommend using a tool which specialises in managing bibliography's like [Zotero](https://zotero.org) 
or the free online version [zoterobib](https://zbib.org/).