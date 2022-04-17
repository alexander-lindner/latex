---
order: -20
icon: rocket
title: Quick Start
---

# Create a sample project and compile it

Install `textool`:

```text curl -L -o /usr/local/bin/textool https://github.com/alexander-lindner/latex/releases/download/v2.1.4/textool-linux-amd64 && chmod +x /usr/local/bin/textool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   658  100   658    0     0   7152      0 --:--:-- --:--:-- --:--:--  7230
100 15.5M  100 15.5M    0     0  5908k      0  0:00:02  0:00:02 --:--:-- 6775k
```
Create your first project, this is an interactive command.

```text  textool -p demo-latex init
INFO[17:27:05] Current binary is the latest version 2.1.3   
INFO[17:27:05] Creating ./.latex as main config file for this tool. 
? Basic: Choose a document class: scrreprt
? Basic: Choose a document class: ngerman
? Basic: Use two column layout? Yes
? Basic: How should the file be named after generation? Default: main.pdf demo.pdf
? Extra: Add a listing engine? Yes
? Extra: Use glossary feature? Yes
? Extra: Use a bibliography engine? Yes
? Extra: use the base docker image(yes) [you need to customize it] or use the provided full image(no)? Yes
? Extra: add some demo content? Yes
INFO[17:27:37] Creating ./latexmkrc which configures latexmk. 
INFO[17:27:37] Creating main tex file main.tex              
INFO[17:27:37] Creating bibliography file main.bib          
INFO[17:27:37] The Dockerfile was created.
```

This creates a directory `demo-latex` and a couple of files:
```text
demo-latex
├── .latex
├── Dockerfile
├── latexmkrc
├── main.bib
└── main.tex

0 directories, 5 files
```
+++ .latex
Stores options of the `textool` executable.
For example, the final pdf file name is stored inside this file.
+++ Dockerfile
Your entrypoint to modify the used image.
If you have chosen the `:base` image, you can install latex packages or even install completly other tools.
The image is base on `ubuntu:latest`, so feel free to use `apt-get`.
+++ latexmkrc
Preconfigured config file for [latexmk](https://ctan.org/pkg/latexmk?lang=en).
+++ main.bib
Bibliography file for bibtex.
+++ main.tex
the main tex file, where you start the project.
+++
After you changed these files, let's compile the project.
TeXtool will download the docker image if not already presented and build a project specific image using the `Dockerfile`.
This image is later-on used to compile the project.
==- textool -p demo-latex run
```text
INFO[13:34:10] Current binary is the latest version 2.1.4   
INFO[13:34:10] Opening config file for  reading. Path:demo-latex/.latex 
INFO[13:34:10] Now pulling the base image: ghcr.io/alexander-lindner/latex:base 
INFO[13:34:10] Image is already available. Not pulling.     
INFO[13:34:10] It is necessary to build the file before using it 
Step 1/3 : FROM ghcr.io/alexander-lindner/latex:base
 ---> 4feb1d254540
Step 2/3 : RUN tlmgr update --self
 ---> Running in 6596d71e648c
tlmgr: package repository https://mirrors.rit.edu/CTAN/systems/texlive/tlnet (verified)
tlmgr: saving backups to /usr/local/texlive/2022/tlpkg/backups
tlmgr: no self-updates for tlmgr available
Removing intermediate container 6596d71e648c
 ---> ebe3b47210f0
Step 3/3 : RUN tlmgr install koma-script xetex xstring float fontspec abstract cleveref hyperref hyphen-german babel-german soul listings minted fvextra fancyvrb upquote lineno xcolor catchfile framed etoolbox glossaries mfirstuc etoolbox textcase xfor datatool tracklang xkeyval glossaries-german csquotes biber biblatex
 ---> Running in 69216d9b7360
tlmgr: package repository https://mirrors.rit.edu/CTAN/systems/texlive/tlnet (verified)
tlmgr install: package already present: hyperref
[1/42, ??:??/??:??] install: abstract [154k]
[2/42, 00:02/17:39] install: babel-german [505k]
[3/42, 00:04/08:13] install: biber.x86_64-linux [22990k]
[4/42, 00:10/00:34] install: biber [1202k]
[5/42, 00:12/00:39] install: biblatex [7375k]
[6/42, 00:16/00:40] install: catchfile [300k]
[7/42, 00:18/00:44] install: cleveref [482k]
[8/42, 00:20/00:49] install: csquotes [328k]
[9/42, 00:21/00:51] install: datatool [2777k]
[10/42, 00:25/00:56] install: etoolbox [241k]
[11/42, 00:26/00:58] install: euenc [153k]
[12/42, 00:28/01:02] install: fancyvrb [158k]
[13/42, 00:29/01:04] install: float [125k]
[14/42, 00:31/01:08] install: fontspec [1300k]
[15/42, 00:33/01:10] install: footmisc [529k]
[16/42, 00:35/01:13] install: fp [233k]
[17/42, 00:37/01:17] install: framed [242k]
[18/42, 00:38/01:18] install: fvextra [860k]
[19/42, 00:40/01:21] install: glossaries.x86_64-linux [1k]
[20/42, 00:41/01:23] install: glossaries [6449k]
[21/42, 00:45/01:18] install: glossaries-german [126k]
[22/42, 00:47/01:22] install: hyphen-german [218k]
[23/42, 00:48/01:23] install: koma-script [12659k]
[24/42, 00:52/01:11] install: lineno [739k]
[25/42, 00:54/01:12] install: listings [2517k]
[26/42, 00:58/01:15] install: logreq [7k]
[27/42, 00:58/01:15] install: mfirstuc [674k]
[28/42, 01:01/01:18] install: minted [859k]
[29/42, 01:03/01:19] install: soul [342k]
[30/42, 01:05/01:21] install: substr [20k]
[31/42, 01:06/01:23] install: textcase [196k]
[32/42, 01:08/01:25] install: tipa [5098k]
[33/42, 01:11/01:22] install: tracklang [1045k]
[34/42, 01:14/01:24] install: upquote [164k]
[35/42, 01:15/01:25] install: xcolor [997k]
[36/42, 01:18/01:27] install: xetex.x86_64-linux [7289k]
[37/42, 01:21/01:22] install: xetex [628k]
[38/42, 01:22/01:23] install: xetexconfig [1k]
[39/42, 01:23/01:24] install: xfor [107k]
[40/42, 01:25/01:26] install: xkeyval [439k]
[41/42, 01:27/01:27] install: xstring [671k]
[42/42, 01:28/01:28] install: xunicode [26k]
running mktexlsr ...
done running mktexlsr.
running updmap-sys ...
done running updmap-sys.
regenerating language.dat
regenerating language.def
regenerating language.dat.lua
regenerating fmtutil.cnf in /usr/local/texlive/2022/texmf-dist
running fmtutil-sys --byengine xetex --no-error-if-no-format --no-error-if-no-engine=luajithbtex,luajittex,mfluajit --status-file=/tmp/xsmWYZjJNT/neKBHzXJvs ...
  OK: xetex.fmt/xetex xelatex.fmt/xetex
running fmtutil-sys --byhyphen "/usr/local/texlive/2022/texmf-var/tex/generic/config/language.dat" --no-error-if-no-engine=luajithbtex,luajittex,mfluajit --status-file=/tmp/xsmWYZjJNT/neKBHzXJvs ...
  OK: xelatex.fmt/xetex dvilualatex.fmt/luatex latex.fmt/pdftex pdflatex.fmt/pdftex lualatex.fmt/luahbtex
running fmtutil-sys --byhyphen "/usr/local/texlive/2022/texmf-var/tex/generic/config/language.def" --no-error-if-no-engine=luajithbtex,luajittex,mfluajit --status-file=/tmp/xsmWYZjJNT/neKBHzXJvs ...
  OK: xetex.fmt/xetex luatex.fmt/luatex luahbtex.fmt/luahbtex pdftex.fmt/pdftex dviluatex.fmt/luatex pdfetex.fmt/pdftex etex.fmt/pdftex
running fmtutil-sys --byhyphen "/usr/local/texlive/2022/texmf-var/tex/generic/config/language.dat.lua" --no-error-if-no-engine=luajithbtex,luajittex,mfluajit --status-file=/tmp/xsmWYZjJNT/neKBHzXJvs ...
tlmgr: package log updated: /usr/local/texlive/2022/texmf-var/web2c/tlmgr.log
tlmgr: command log updated: /usr/local/texlive/2022/texmf-var/web2c/tlmgr-commands.log
Removing intermediate container 69216d9b7360
 ---> 6858eab40605
Successfully built 6858eab40605
Successfully tagged textool-tmp-xv/bnyoroi8lrb7vrggiiihfam:latest
INFO[13:36:03] Creating the docker container                
INFO[13:36:03] Starting the docker container                
Rc files read:
  latexmkrc
Latexmk: Run number 1 of rule 'pdflatex'
This is pdfTeX, Version 3.141592653-2.6-1.40.24 (TeX Live 2022) (preloaded format=pdflatex)
 \write18 enabled.
entering extended mode
/usr/local/bin/pygmentize
kpathsea: Running mktexpk --mfmode / --bdpi 600 --mag 1+0/600 --dpi 600 tctt1095
mktexpk: Running mf-nowin -progname=mf \mode:=ljfour; mag:=1+0/600; nonstopmode; input tctt1095
This is METAFONT, Version 2.71828182 (TeX Live 2022) (preloaded base=mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/tctt1095.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbase.mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/tctt.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txsymb.mf
 Ok (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccess.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txpseudo.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txaccent.mf
 Ok [0] [1] [2] [3] [4] [5] [6] [7] [8] [9] [10] [11] [12] [27] [29])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txgen.mf
 Ok [100] [109] [98] [99] [108])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txsymbol.mf
 Ok [13] [18] [21] [22] [23] [24] [25] [26] [28] [31] [32] [36] [39] [44]
[45] [46] [42] [47] [60] [61] [62] [77] [79] [87] [110] [91] [93] [94] [95]
[96] [126] [127] [128] [129] [130] [131] [132] [133] [134] [135] [136] [137]
[138] [139] [140] [141] [142] [143] [144] [145] [146] [147] [148] [149]
[150] [151] [152] [153] [154] [155] [156] [157] [158] [159] [160] [161]
[162] [163] [164] [165] [166] [167] [168] [169] [171] [172] [173] [174]
[175] [177] [176] [180] [181] [182] [183] [184] [187] [191] [214] [246])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txromod.mf
 Ok [48] [49] [50] [51] [52] [53] [54] [55] [56] [57])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txrsuper.mf
 Ok [185] [178] [179] [170] [186])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txrfract.mf
 Ok [188] [189] [190]) ) ) )
Font metrics written on tctt1095.tfm.
Output written on tctt1095.600gf (128 characters, 21204 bytes).
Transcript written on tctt1095.log.
mktexpk: /data/.texlive2022/texmf-var/fonts/pk/ljfour/jknappen/ec/tctt1095.600pk: successfully generated.
kpathsea: Running mktexpk --mfmode / --bdpi 600 --mag 1+0/600 --dpi 600 ecrm0600
mktexpk: Running mf-nowin -progname=mf \mode:=ljfour; mag:=1+0/600; nonstopmode; input ecrm0600
This is METAFONT, Version 2.71828182 (TeX Live 2022) (preloaded base=mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecrm0600.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbase.mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecrm.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exroman.mf
 Ok (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccess.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expseudo.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exruwest.mf
 Ok [192] [193] [194] [195] [196] [197] [198] [199] [200] [201] [202] [203]
[204] [205] [206] [207] [208] [209] [210] [211] [212] [213] [214] [215]
[216] [217] [218] [219] [220] [221] [222] [223])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlwest.mf
 Ok [224] [225] [226] [227] [228] [229] [230] [231] [232] [233] [234] [235]
[236] [237] [238] [239] [240] [241] [242] [243] [244] [245] [246] [247]
[248] [249] [250] [251] [252] [253] [254] [255])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrueast.mf
 Ok [128] [129] [130] [131] [132] [133] [134] [135] [136] [137] [138] [139]
[140] [141] [142] [143] [144] [145] [146] [147] [148] [149] [150] [151]
[152] [153] [154] [155] [156] [157])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrleast.mf
 Ok [158] [160] [161] [162] [163] [164] [165] [166] [167] [168] [169] [170]
[171] [172] [173] [174] [175] [176] [177] [178] [179] [180] [181] [182]
[183] [184] [185] [186] [187] [188])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbraces.mf
 Ok [94] [126] [23] [40] [41] [60] [124] [62] [91] [93] [92] [123] [125]
[95] [127] [32])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expunct.mf
 Ok [14] [15] [19] [20] [13] [18] [33] [39] [42] [43] [44] [46] [47] [58]
[59] [61] [96] [189] [17] [45] [16] [21] [22])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccent.mf
 Ok [0] [1] [2] [3] [4] [5] [6] [7] [8] [9] [10] [11] [12])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exsign.mf
 Ok [24] [34] [35] [36] [37] [64] [191] [159])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlig.mf
 Ok [25] [26] [28] [27] [29] [30] [31])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exromp.mf
 Ok [38] [63] [190])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrulett.mf
 Ok [65] [66] [67] [68] [69] [70] [71] [72] [73] [74] [75] [76] [77] [78]
[79] [80] [81] [82] [83] [84] [85] [86] [87] [88] [89] [90])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrllett.mf
 Ok [97] [98] [99] [100] [101] [102] [103] [104] [105] [106] [107] [108]
[109] [110] [111] [112] [113] [114] [115] [116] [117] [118] [119] [120]
[121] [122])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrdigit.mf
 Ok [48] [49] [50] [51] [52] [53] [54] [55] [56] [57])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrligtb.mf
 Ok) ) ) )
Font metrics written on ecrm0600.tfm.
Output written on ecrm0600.600gf (256 characters, 32624 bytes).
Transcript written on ecrm0600.log.
mktexpk: /data/.texlive2022/texmf-var/fonts/pk/ljfour/jknappen/ec/ecrm0600.600pk: successfully generated.
kpathsea: Running mktexpk --mfmode / --bdpi 600 --mag 1+0/600 --dpi 600 ectt1095
mktexpk: Running mf-nowin -progname=mf \mode:=ljfour; mag:=1+0/600; nonstopmode; input ectt1095
This is METAFONT, Version 2.71828182 (TeX Live 2022) (preloaded base=mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ectt1095.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbase.mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ectt.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exroman.mf
 Ok (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccess.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expseudo.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exruwest.mf
 Ok [192] [193] [194] [195] [196] [197] [198] [199] [200] [201] [202] [203]
[204] [205] [206] [207] [208] [209] [210] [211] [212] [213] [214] [215]
[216] [217] [218] [219] [220] [221] [222] [223])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlwest.mf
 Ok [224] [225] [226] [227] [228] [229] [230] [231] [232] [233] [234] [235]
[236] [237] [238] [239] [240] [241] [242] [243] [244] [245] [246] [247]
[248] [249] [250] [251] [252] [253] [254] [255])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrueast.mf
 Ok [128] [129] [130] [131] [132] [133] [134] [135] [136] [137] [138] [139]
[140] [141] [142] [143] [144] [145] [146] [147] [148] [149] [150] [151]
[152] [153] [154] [155] [156] [157])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrleast.mf
 Ok [158] [160] [161] [162] [163] [164] [165] [166] [167] [168] [169] [170]
[171] [172] [173] [174] [175] [176] [177] [178] [179] [180] [181] [182]
[183] [184] [185] [186] [187] [188])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbraces.mf
 Ok [94] [126] [23] [40] [41] [60] [124] [62] [91] [93] [92] [123] [125]
[95] [127] [32])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expunct.mf
 Ok [14] [15] [19] [20] [13] [18] [33] [39] [42] [43] [44] [46] [47] [58]
[59] [61] [96] [189] [17] [45] [16] [21] [22])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccent.mf
 Ok [0] [1] [2] [3] [4] [5] [6] [7] [8] [9] [10] [11] [12])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exsign.mf
 Ok [24] [34] [35] [36] [37] [64] [191] [159])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlig.mf
 Ok [25] [26] [28] [27] [29] [30] [31])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exromp.mf
 Ok [38] [63] [190])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrulett.mf
 Ok [65] [66] [67] [68] [69] [70] [71] [72] [73] [74] [75] [76] [77] [78]
[79] [80] [81] [82] [83] [84] [85] [86] [87] [88] [89] [90])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrllett.mf
 Ok [97] [98] [99] [100] [101] [102] [103] [104] [105] [106] [107] [108]
[109] [110] [111] [112] [113] [114] [115] [116] [117] [118] [119] [120]
[121] [122])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrdigit.mf
 Ok [48] [49] [50] [51] [52] [53] [54] [55] [56] [57])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exmligtb.mf
 Ok) ) ) )
Font metrics written on ectt1095.tfm.
Output written on ectt1095.600gf (256 characters, 47796 bytes).
Transcript written on ectt1095.log.
mktexpk: /data/.texlive2022/texmf-var/fonts/pk/ljfour/jknappen/ec/ectt1095.600pk: successfully generated.
kpathsea: Running mktexpk --mfmode / --bdpi 600 --mag 1+0/600 --dpi 600 tcrm1095
mktexpk: Running mf-nowin -progname=mf \mode:=ljfour; mag:=1+0/600; nonstopmode; input tcrm1095
This is METAFONT, Version 2.71828182 (TeX Live 2022) (preloaded base=mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/tcrm1095.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbase.mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/tcrm.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txsymb.mf
 Ok (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccess.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txpseudo.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txaccent.mf
 Ok [0] [1] [2] [3] [4] [5] [6] [7] [8] [9] [10] [11] [12] [27] [29])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txgen.mf
 Ok [100] [109] [98] [99] [108])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txsymbol.mf
 Ok [13] [18] [21] [22] [23] [24] [25] [26] [28] [31] [32] [36] [39] [44]
[45] [46] [42] [47] [60] [61] [62] [77] [79] [87] [110] [91] [93] [94] [95]
[96] [126] [127] [128] [129] [130] [131] [132] [133] [134] [135] [136] [137]
[138] [139] [140] [141] [142] [143] [144] [145] [146] [147] [148] [149]
[150] [151] [152] [153] [154] [155] [156] [157] [158] [159] [160] [161]
[162] [163] [164] [165] [166] [167] [168] [169] [171] [172] [173] [174]
[175] [177] [176] [180] [181] [182] [183] [184] [187] [191] [214] [246])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txromod.mf
 Ok [48] [49] [50] [51] [52] [53] [54] [55] [56] [57])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txrsuper.mf
 Ok [185] [178] [179] [170] [186])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txrfract.mf
 Ok [188] [189] [190]) ) ) )
(some charht values had to be adjusted by as much as 0.06952pt)
Font metrics written on tcrm1095.tfm.
Output written on tcrm1095.600gf (128 characters, 25592 bytes).
Transcript written on tcrm1095.log.
mktexpk: /data/.texlive2022/texmf-var/fonts/pk/ljfour/jknappen/ec/tcrm1095.600pk: successfully generated.
kpathsea: Running mktexpk --mfmode / --bdpi 600 --mag 1+0/600 --dpi 600 ecsx1440
mktexpk: Running mf-nowin -progname=mf \mode:=ljfour; mag:=1+0/600; nonstopmode; input ecsx1440
This is METAFONT, Version 2.71828182 (TeX Live 2022) (preloaded base=mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecsx1440.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbase.mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecsx.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exroman.mf
 Ok (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccess.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expseudo.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exruwest.mf
 Ok [192] [193] [194] [195] [196] [197] [198] [199] [200] [201] [202] [203]
[204] [205] [206] [207] [208] [209] [210] [211] [212] [213] [214] [215]
[216] [217] [218] [219] [220] [221] [222] [223])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlwest.mf
 Ok [224] [225] [226] [227] [228] [229] [230] [231] [232] [233] [234] [235]
[236] [237] [238] [239] [240] [241] [242] [243] [244] [245] [246] [247]
[248] [249] [250] [251] [252] [253] [254] [255])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrueast.mf
 Ok [128] [129] [130] [131] [132] [133] [134] [135] [136] [137] [138] [139]
[140] [141] [142] [143] [144] [145] [146] [147] [148] [149] [150] [151]
[152] [153] [154] [155] [156] [157])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrleast.mf
 Ok [158] [160] [161] [162] [163] [164] [165] [166] [167] [168] [169] [170]
[171] [172] [173] [174] [175] [176] [177] [178] [179] [180] [181] [182]
[183] [184] [185] [186] [187] [188])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbraces.mf
 Ok [94] [126] [23] [40] [41] [60] [124] [62] [91] [93] [92] [123] [125]
[95] [127] [32])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expunct.mf
 Ok [14] [15] [19] [20] [13] [18] [33] [39] [42] [43] [44] [46] [47] [58]
[59] [61] [96] [189] [17] [45] [16] [21] [22])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccent.mf
 Ok [0] [1] [2] [3] [4] [5] [6] [7] [8] [9] [10] [11] [12])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exsign.mf
 Ok [24] [34] [35] [36] [37] [64] [191] [159])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlig.mf
 Ok [25] [26] [28] [27] [29] [30] [31])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exromp.mf
 Ok [38] [63] [190])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrulett.mf
 Ok [65] [66] [67] [68] [69] [70] [71] [72] [73] [74] [75] [76] [77] [78]
[79] [80] [81] [82] [83] [84] [85] [86] [87] [88] [89] [90])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrllett.mf
 Ok [97] [98] [99] [100] [101] [102] [103] [104] [105] [106] [107] [108]
[109] [110] [111] [112] [113] [114] [115] [116] [117] [118] [119] [120]
[121] [122])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrdigit.mf
 Ok [48] [49] [50] [51] [52] [53] [54] [55] [56] [57])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrligtb.mf
 Ok) ) ) )
(some charht values had to be adjusted by as much as 0.07274pt)
Font metrics written on ecsx1440.tfm.
Output written on ecsx1440.600gf (256 characters, 64452 bytes).
Transcript written on ecsx1440.log.
mktexpk: /data/.texlive2022/texmf-var/fonts/pk/ljfour/jknappen/ec/ecsx1440.600pk: successfully generated.
kpathsea: Running mktexpk --mfmode / --bdpi 600 --mag 1+0/600 --dpi 600 ecrm1095
mktexpk: Running mf-nowin -progname=mf \mode:=ljfour; mag:=1+0/600; nonstopmode; input ecrm1095
This is METAFONT, Version 2.71828182 (TeX Live 2022) (preloaded base=mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecrm1095.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbase.mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecrm.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exroman.mf
 Ok (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccess.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expseudo.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exruwest.mf
 Ok [192] [193] [194] [195] [196] [197] [198] [199] [200] [201] [202] [203]
[204] [205] [206] [207] [208] [209] [210] [211] [212] [213] [214] [215]
[216] [217] [218] [219] [220] [221] [222] [223])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlwest.mf
 Ok [224] [225] [226] [227] [228] [229] [230] [231] [232] [233] [234] [235]
[236] [237] [238] [239] [240] [241] [242] [243] [244] [245] [246] [247]
[248] [249] [250] [251] [252] [253] [254] [255])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrueast.mf
 Ok [128] [129] [130] [131] [132] [133] [134] [135] [136] [137] [138] [139]
[140] [141] [142] [143] [144] [145] [146] [147] [148] [149] [150] [151]
[152] [153] [154] [155] [156] [157])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrleast.mf
 Ok [158] [160] [161] [162] [163] [164] [165] [166] [167] [168] [169] [170]
[171] [172] [173] [174] [175] [176] [177] [178] [179] [180] [181] [182]
[183] [184] [185] [186] [187] [188])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbraces.mf
 Ok [94] [126] [23] [40] [41] [60] [124] [62] [91] [93] [92] [123] [125]
[95] [127] [32])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expunct.mf
 Ok [14] [15] [19] [20] [13] [18] [33] [39] [42] [43] [44] [46] [47] [58]
[59] [61] [96] [189] [17] [45] [16] [21] [22])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccent.mf
 Ok [0] [1] [2] [3] [4] [5] [6] [7] [8] [9] [10] [11] [12])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exsign.mf
 Ok [24] [34] [35] [36] [37] [64] [191] [159])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlig.mf
 Ok [25] [26] [28] [27] [29] [30] [31])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exromp.mf
 Ok [38] [63] [190])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrulett.mf
 Ok [65] [66] [67] [68] [69] [70] [71] [72] [73] [74] [75] [76] [77] [78]
[79] [80] [81] [82] [83] [84] [85] [86] [87] [88] [89] [90])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrllett.mf
 Ok [97] [98] [99] [100] [101] [102] [103] [104] [105] [106] [107] [108]
[109] [110] [111] [112] [113] [114] [115] [116] [117] [118] [119] [120]
[121] [122])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrdigit.mf
 Ok [48] [49] [50] [51] [52] [53] [54] [55] [56] [57])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrligtb.mf
 Ok) ) ) )
(some charht values had to be adjusted by as much as 0.07599pt)
Font metrics written on ecrm1095.tfm.
Output written on ecrm1095.600gf (256 characters, 55424 bytes).
Transcript written on ecrm1095.log.
mktexpk: /data/.texlive2022/texmf-var/fonts/pk/ljfour/jknappen/ec/ecrm1095.600pk: successfully generated.
kpathsea: Running mktexpk --mfmode / --bdpi 600 --mag 1+0/600 --dpi 600 ecbx1095
mktexpk: Running mf-nowin -progname=mf \mode:=ljfour; mag:=1+0/600; nonstopmode; input ecbx1095
This is METAFONT, Version 2.71828182 (TeX Live 2022) (preloaded base=mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecbx1095.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbase.mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecbx.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exroman.mf
 Ok (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccess.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expseudo.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exruwest.mf
 Ok [192] [193] [194] [195] [196] [197] [198] [199] [200] [201] [202] [203]
[204] [205] [206] [207] [208] [209] [210] [211] [212] [213] [214] [215]
[216] [217] [218] [219] [220] [221] [222] [223])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlwest.mf
 Ok [224] [225] [226] [227] [228] [229] [230] [231] [232] [233] [234] [235]
[236] [237] [238] [239] [240] [241] [242] [243] [244] [245] [246] [247]
[248] [249] [250] [251] [252] [253] [254] [255])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrueast.mf
 Ok [128] [129] [130] [131] [132] [133] [134] [135] [136] [137] [138] [139]
[140] [141] [142] [143] [144] [145] [146] [147] [148] [149] [150] [151]
[152] [153] [154] [155] [156] [157])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrleast.mf
 Ok [158] [160] [161] [162] [163] [164] [165] [166] [167] [168] [169] [170]
[171] [172] [173] [174] [175] [176] [177] [178] [179] [180] [181] [182]
[183] [184] [185] [186] [187] [188])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbraces.mf
 Ok [94] [126] [23] [40] [41] [60] [124] [62] [91] [93] [92] [123] [125]
[95] [127] [32])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expunct.mf
 Ok [14] [15] [19] [20] [13] [18] [33] [39] [42] [43] [44] [46] [47] [58]
[59] [61] [96] [189] [17] [45] [16] [21] [22])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccent.mf
 Ok [0] [1] [2] [3] [4] [5] [6] [7] [8] [9] [10] [11] [12])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exsign.mf
 Ok [24] [34] [35] [36] [37] [64] [191] [159])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlig.mf
 Ok [25] [26] [28] [27] [29] [30] [31])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exromp.mf
 Ok [38] [63] [190])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrulett.mf
 Ok [65] [66] [67] [68] [69] [70] [71] [72] [73] [74] [75] [76] [77] [78]
[79] [80] [81] [82] [83] [84] [85] [86] [87] [88] [89] [90])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrllett.mf
 Ok [97] [98] [99] [100] [101] [102] [103] [104] [105] [106] [107] [108]
[109] [110] [111] [112] [113] [114] [115] [116] [117] [118] [119] [120]
[121] [122])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrdigit.mf
 Ok [48] [49] [50] [51] [52] [53] [54] [55] [56] [57])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrligtb.mf
 Ok) ) ) )
Font metrics written on ecbx1095.tfm.
Output written on ecbx1095.600gf (256 characters, 54848 bytes).
Transcript written on ecbx1095.log.
mktexpk: /data/.texlive2022/texmf-var/fonts/pk/ljfour/jknappen/ec/ecbx1095.600pk: successfully generated.
kpathsea: Running mktexpk --mfmode / --bdpi 600 --mag 1+0/600 --dpi 600 ecrm1440
mktexpk: Running mf-nowin -progname=mf \mode:=ljfour; mag:=1+0/600; nonstopmode; input ecrm1440
This is METAFONT, Version 2.71828182 (TeX Live 2022) (preloaded base=mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecrm1440.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbase.mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecrm.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exroman.mf
 Ok (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccess.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expseudo.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exruwest.mf
 Ok [192] [193] [194] [195] [196] [197] [198] [199] [200] [201] [202] [203]
[204] [205] [206] [207] [208] [209] [210] [211] [212] [213] [214] [215]
[216] [217] [218] [219] [220] [221] [222] [223])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlwest.mf
 Ok [224] [225] [226] [227] [228] [229] [230] [231] [232] [233] [234] [235]
[236] [237] [238] [239] [240] [241] [242] [243] [244] [245] [246] [247]
[248] [249] [250] [251] [252] [253] [254] [255])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrueast.mf
 Ok [128] [129] [130] [131] [132] [133] [134] [135] [136] [137] [138] [139]
[140] [141] [142] [143] [144] [145] [146] [147] [148] [149] [150] [151]
[152] [153] [154] [155] [156] [157])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrleast.mf
 Ok [158] [160] [161] [162] [163] [164] [165] [166] [167] [168] [169] [170]
[171] [172] [173] [174] [175] [176] [177] [178] [179] [180] [181] [182]
[183] [184] [185] [186] [187] [188])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbraces.mf
 Ok [94] [126] [23] [40] [41] [60] [124] [62] [91] [93] [92] [123] [125]
[95] [127] [32])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expunct.mf
 Ok [14] [15] [19] [20] [13] [18] [33] [39] [42] [43] [44] [46] [47] [58]
[59] [61] [96] [189] [17] [45] [16] [21] [22])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccent.mf
 Ok [0] [1] [2] [3] [4] [5] [6] [7] [8] [9] [10] [11] [12])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exsign.mf
 Ok [24] [34] [35] [36] [37] [64] [191] [159])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlig.mf
 Ok [25] [26] [28] [27] [29] [30] [31])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exromp.mf
 Ok [38] [63] [190])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrulett.mf
 Ok [65] [66] [67] [68] [69] [70] [71] [72] [73] [74] [75] [76] [77] [78]
[79] [80] [81] [82] [83] [84] [85] [86] [87] [88] [89] [90])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrllett.mf
 Ok [97] [98] [99] [100] [101] [102] [103] [104] [105] [106] [107] [108]
[109] [110] [111] [112] [113] [114] [115] [116] [117] [118] [119] [120]
[121] [122])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrdigit.mf
 Ok [48] [49] [50] [51] [52] [53] [54] [55] [56] [57])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrligtb.mf
 Ok) ) ) )
(some charht values had to be adjusted by as much as 0.08084pt)
Font metrics written on ecrm1440.tfm.
Output written on ecrm1440.600gf (256 characters, 71064 bytes).
Transcript written on ecrm1440.log.
mktexpk: /data/.texlive2022/texmf-var/fonts/pk/ljfour/jknappen/ec/ecrm1440.600pk: successfully generated.
kpathsea: Running mktexpk --mfmode / --bdpi 600 --mag 1+0/600 --dpi 600 ecsx1200
mktexpk: Running mf-nowin -progname=mf \mode:=ljfour; mag:=1+0/600; nonstopmode; input ecsx1200
This is METAFONT, Version 2.71828182 (TeX Live 2022) (preloaded base=mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecsx1200.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbase.mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecsx.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exroman.mf
 Ok (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccess.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expseudo.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exruwest.mf
 Ok [192] [193] [194] [195] [196] [197] [198] [199] [200] [201] [202] [203]
[204] [205] [206] [207] [208] [209] [210] [211] [212] [213] [214] [215]
[216] [217] [218] [219] [220] [221] [222] [223])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlwest.mf
 Ok [224] [225] [226] [227] [228] [229] [230] [231] [232] [233] [234] [235]
[236] [237] [238] [239] [240] [241] [242] [243] [244] [245] [246] [247]
[248] [249] [250] [251] [252] [253] [254] [255])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrueast.mf
 Ok [128] [129] [130] [131] [132] [133] [134] [135] [136] [137] [138] [139]
[140] [141] [142] [143] [144] [145] [146] [147] [148] [149] [150] [151]
[152] [153] [154] [155] [156] [157])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrleast.mf
 Ok [158] [160] [161] [162] [163] [164] [165] [166] [167] [168] [169] [170]
[171] [172] [173] [174] [175] [176] [177] [178] [179] [180] [181] [182]
[183] [184] [185] [186] [187] [188])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbraces.mf
 Ok [94] [126] [23] [40] [41] [60] [124] [62] [91] [93] [92] [123] [125]
[95] [127] [32])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expunct.mf
 Ok [14] [15] [19] [20] [13] [18] [33] [39] [42] [43] [44] [46] [47] [58]
[59] [61] [96] [189] [17] [45] [16] [21] [22])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccent.mf
 Ok [0] [1] [2] [3] [4] [5] [6] [7] [8] [9] [10] [11] [12])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exsign.mf
 Ok [24] [34] [35] [36] [37] [64] [191] [159])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlig.mf
 Ok [25] [26] [28] [27] [29] [30] [31])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exromp.mf
 Ok [38] [63] [190])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrulett.mf
 Ok [65] [66] [67] [68] [69] [70] [71] [72] [73] [74] [75] [76] [77] [78]
[79] [80] [81] [82] [83] [84] [85] [86] [87] [88] [89] [90])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrllett.mf
 Ok [97] [98] [99] [100] [101] [102] [103] [104] [105] [106] [107] [108]
[109] [110] [111] [112] [113] [114] [115] [116] [117] [118] [119] [120]
[121] [122])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrdigit.mf
 Ok [48] [49] [50] [51] [52] [53] [54] [55] [56] [57])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrligtb.mf
 Ok) ) ) )
(some charht values had to be adjusted by as much as 0.07016pt)
Font metrics written on ecsx1200.tfm.
Output written on ecsx1200.600gf (256 characters, 54960 bytes).
Transcript written on ecsx1200.log.
mktexpk: /data/.texlive2022/texmf-var/fonts/pk/ljfour/jknappen/ec/ecsx1200.600pk: successfully generated.
kpathsea: Running mktexpk --mfmode / --bdpi 600 --mag 1+0/600 --dpi 600 ecsx2074
mktexpk: Running mf-nowin -progname=mf \mode:=ljfour; mag:=1+0/600; nonstopmode; input ecsx2074
This is METAFONT, Version 2.71828182 (TeX Live 2022) (preloaded base=mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecsx2074.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbase.mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecsx.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exroman.mf
 Ok (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccess.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expseudo.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exruwest.mf
 Ok [192] [193] [194] [195] [196] [197] [198] [199] [200] [201] [202] [203]
[204] [205] [206] [207] [208] [209] [210] [211] [212] [213] [214] [215]
[216] [217] [218] [219] [220] [221] [222] [223])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlwest.mf
 Ok [224] [225] [226] [227] [228] [229] [230] [231] [232] [233] [234] [235]
[236] [237] [238] [239] [240] [241] [242] [243] [244] [245] [246] [247]
[248] [249] [250] [251] [252] [253] [254] [255])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrueast.mf
 Ok [128] [129] [130] [131] [132] [133] [134] [135] [136] [137] [138] [139]
[140] [141] [142] [143] [144] [145] [146] [147] [148] [149] [150] [151]
[152] [153] [154] [155] [156] [157])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrleast.mf
 Ok [158] [160] [161] [162] [163] [164] [165] [166] [167] [168] [169] [170]
[171] [172] [173] [174] [175] [176] [177] [178] [179] [180] [181] [182]
[183] [184] [185] [186] [187] [188])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbraces.mf
 Ok [94] [126] [23] [40] [41] [60] [124] [62] [91] [93] [92] [123] [125]
[95] [127] [32])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expunct.mf
 Ok [14] [15] [19] [20] [13] [18] [33] [39] [42] [43] [44] [46] [47] [58]
[59] [61] [96] [189] [17] [45] [16] [21] [22])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccent.mf
 Ok [0] [1] [2] [3] [4] [5] [6] [7] [8] [9] [10] [11] [12])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exsign.mf
 Ok [24] [34] [35] [36] [37] [64] [191] [159])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlig.mf
 Ok [25] [26] [28] [27] [29] [30] [31])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exromp.mf
 Ok [38] [63] [190])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrulett.mf
 Ok [65] [66] [67] [68] [69] [70] [71] [72] [73] [74] [75] [76] [77] [78]
[79] [80] [81] [82] [83] [84] [85] [86] [87] [88] [89] [90])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrllett.mf
 Ok [97] [98] [99] [100] [101] [102] [103] [104] [105] [106] [107] [108]
[109] [110] [111] [112] [113] [114] [115] [116] [117] [118] [119] [120]
[121] [122])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrdigit.mf
 Ok [48] [49] [50] [51] [52] [53] [54] [55] [56] [57])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrligtb.mf
 Ok) ) ) )
(some charht values had to be adjusted by as much as 0.13551pt)
Font metrics written on ecsx2074.tfm.
Output written on ecsx2074.600gf (256 characters, 93548 bytes).
Transcript written on ecsx2074.log.
mktexpk: /data/.texlive2022/texmf-var/fonts/pk/ljfour/jknappen/ec/ecsx2074.600pk: successfully generated.
Latexmk: Getting log file 'out/main.log'
Latexmk: Missing bbl file 'out/main.bbl' in following:
 No file main.bbl.
Latexmk: Run number 1 of rule 'cusdep glo gls out/main'
This is makeindex, version 2.16 [TeX Live 2022] (kpathsea + Thai support).
Scanning style file ./out/main.ist.............................done (29 attributes redefined, 0 ignored).
Scanning input file out/main.glo....done (1 entries accepted, 0 rejected).
Sorting entries...done (0 comparisons).
Generating output file out/main.gls....done (6 lines written, 0 warnings).
Output written in out/main.gls.
Transcript written in out/main.glg.
Latexmk: Run number 1 of rule 'biber out/main'
Latexmk: Biber did't find bib file [main.bib]
Latexmk: Run number 2 of rule 'pdflatex'
This is pdfTeX, Version 3.141592653-2.6-1.40.24 (TeX Live 2022) (preloaded format=pdflatex)
 \write18 enabled.
entering extended mode
/usr/local/bin/pygmentize
kpathsea: Running mktexpk --mfmode / --bdpi 600 --mag 1+0/600 --dpi 600 eccc1095
mktexpk: Running mf-nowin -progname=mf \mode:=ljfour; mag:=1+0/600; nonstopmode; input eccc1095
This is METAFONT, Version 2.71828182 (TeX Live 2022) (preloaded base=mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/eccc1095.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbase.mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/eccc.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/excsc.mf
 Ok (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccess.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expseudo.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exruwest.mf
 Ok [192] [193] [194] [195] [196] [197] [198] [199] [200] [201] [202] [203]
[204] [205] [206] [207] [208] [209] [210] [211] [212] [213] [214] [215]
[216] [217] [218] [219] [220] [221] [222] [223])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrueast.mf
 Ok [128] [129] [130] [131] [132] [133] [134] [135] [136] [137] [138] [139]
[140] [141] [142] [143] [144] [145] [146] [147] [148] [149] [150] [151]
[152] [153] [154] [155] [156] [157])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrulett.mf
 Ok [65] [66] [67] [68] [69] [70] [71] [72] [73] [74] [75] [76] [77] [78]
[79] [80] [81] [82] [83] [84] [85] [86] [87] [88] [89] [90])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbraces.mf
 Ok [94] [126] [23] [40] [41] [60] [124] [62] [91] [93] [92] [123] [125]
[95] [127] [32])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expunct.mf
 Ok [14] [15] [19] [20] [13] [18] [33] [39] [42] [43] [44] [46] [47] [58]
[59] [61] [96] [189] [17] [45] [16] [21] [22])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exsign.mf
 Ok [24] [34] [35] [36] [37] [64] [191] [159])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exromp.mf
 Ok [38] [63] [190])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrdigit.mf
 Ok [48] [49] [50] [51] [52] [53] [54] [55] [56] [57])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expseudo.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/excspl.mf
 Ok [25] [26] [27] [28] [29] [30] [31] [158])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrulett.mf
 Ok [97] [98] [99] [100] [101] [102] [103] [104] [105] [106] [107] [108]
[109] [110] [111] [112] [113] [114] [115] [116] [117] [118] [119] [120]
[121] [122])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exruwest.mf
 Ok [224] [225] [226] [227] [228] [229] [230] [231] [232] [233] [234] [235]
[236] [237] [238] [239] [240] [241] [242] [243] [244] [245] [246] [247]
[248] [249] [250] [251] [252] [253] [254] [255])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrueast.mf
 Ok [160] [161] [162] [163] [164] [165] [166] [167] [168] [169] [170] [171]
[172] [173] [174] [175] [176] [177] [178] [179] [180] [181] [182] [183]
[184] [185] [186] [187] [188])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txpseudo.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/txaccent.mf
 Ok [0] [1] [2] [3] [4] [5] [6] [7] [8] [9] [10] [11] [12])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/excligtb.mf
 Ok) ) ) )
(some charht values had to be adjusted by as much as 0.0856pt)
Font metrics written on eccc1095.tfm.
Output written on eccc1095.600gf (256 characters, 56408 bytes).
Transcript written on eccc1095.log.
mktexpk: /data/.texlive2022/texmf-var/fonts/pk/ljfour/jknappen/ec/eccc1095.600pk: successfully generated.
kpathsea: Running mktexpk --mfmode / --bdpi 600 --mag 1+0/600 --dpi 600 ecti1095
mktexpk: Running mf-nowin -progname=mf \mode:=ljfour; mag:=1+0/600; nonstopmode; input ecti1095
This is METAFONT, Version 2.71828182 (TeX Live 2022) (preloaded base=mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecti1095.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbase.mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecti.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/extextit.mf
 Ok (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccess.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expseudo.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exileast.mf
 Ok [158] [160] [161] [162] [163] [164] [165] [166] [167] [168] [169] [170]
[171] [172] [173] [174] [175] [176] [177] [178] [179] [180] [181] [182]
[183] [184] [185] [186] [187] [188])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exilwest.mf
 Ok [224] [225] [226] [227] [228] [229] [230] [231] [232] [233] [234] [235]
[236] [237] [238] [239] [240] [241] [242] [243] [244] [245] [246] [247]
[248] [249] [250] [251] [252] [253] [254] [255])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exruwest.mf
 Ok [192] [193] [194] [195] [196] [197] [198] [199] [200] [201] [202] [203]
[204] [205] [206] [207] [208] [209] [210] [211] [212] [213] [214] [215]
[216] [217] [218] [219] [220] [221] [222] [223])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrueast.mf
 Ok [128] [129] [130] [131] [132] [133] [134] [135] [136] [137] [138] [139]
[140] [141] [142] [143] [144] [145] [146] [147] [148] [149] [150] [151]
[152] [153] [154] [155] [156] [157])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbraces.mf
 Ok [94] [126] [23] [40] [41] [60] [124] [62] [91] [93] [92] [123] [125]
[95] [127] [32])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expunct.mf
 Ok [14] [15] [19] [20] [13] [18] [33] [39] [42] [43] [44] [46] [47] [58]
[59] [61] [96] [189] [17] [45] [16] [21] [22])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccent.mf
 Ok [0] [1] [2] [3] [4] [5] [6] [7] [8] [9] [10] [11] [12])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exsign.mf
 Ok [24] [34] [35] [36] [37] [64] [191] [159])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exilig.mf
 Ok [25] [26] [27] [28] [29] [30] [31])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exitalp.mf
 Ok [38] [63] [190])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrulett.mf
 Ok [65] [66] [67] [68] [69] [70] [71] [72] [73] [74] [75] [76] [77] [78]
[79] [80] [81] [82] [83] [84] [85] [86] [87] [88] [89] [90])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exillett.mf
 Ok [97] [98] [99] [100] [101] [102] [103] [104] [105] [106] [107] [108]
[109] [110] [111] [112] [113] [114] [115] [116] [117] [118] [119] [120]
[121] [122])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exidigit.mf
 Ok [48] [49] [50] [51] [52] [53] [54] [55] [56] [57])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exiligtb.mf
 Ok) ) ) )
(some charht values had to be adjusted by as much as 0.06808pt)
Font metrics written on ecti1095.tfm.
Output written on ecti1095.600gf (256 characters, 57920 bytes).
Transcript written on ecti1095.log.
mktexpk: /data/.texlive2022/texmf-var/fonts/pk/ljfour/jknappen/ec/ecti1095.600pk: successfully generated.
kpathsea: Running mktexpk --mfmode / --bdpi 600 --mag 1+0/600 --dpi 600 ecsx1095
mktexpk: Running mf-nowin -progname=mf \mode:=ljfour; mag:=1+0/600; nonstopmode; input ecsx1095
This is METAFONT, Version 2.71828182 (TeX Live 2022) (preloaded base=mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecsx1095.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbase.mf)
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/ecsx.mf
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exroman.mf
 Ok (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccess.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expseudo.mf
 Ok) (/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exruwest.mf
 Ok [192] [193] [194] [195] [196] [197] [198] [199] [200] [201] [202] [203]
[204] [205] [206] [207] [208] [209] [210] [211] [212] [213] [214] [215]
[216] [217] [218] [219] [220] [221] [222] [223])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlwest.mf
 Ok [224] [225] [226] [227] [228] [229] [230] [231] [232] [233] [234] [235]
[236] [237] [238] [239] [240] [241] [242] [243] [244] [245] [246] [247]
[248] [249] [250] [251] [252] [253] [254] [255])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrueast.mf
 Ok [128] [129] [130] [131] [132] [133] [134] [135] [136] [137] [138] [139]
[140] [141] [142] [143] [144] [145] [146] [147] [148] [149] [150] [151]
[152] [153] [154] [155] [156] [157])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrleast.mf
 Ok [158] [160] [161] [162] [163] [164] [165] [166] [167] [168] [169] [170]
[171] [172] [173] [174] [175] [176] [177] [178] [179] [180] [181] [182]
[183] [184] [185] [186] [187] [188])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exbraces.mf
 Ok [94] [126] [23] [40] [41] [60] [124] [62] [91] [93] [92] [123] [125]
[95] [127] [32])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/expunct.mf
 Ok [14] [15] [19] [20] [13] [18] [33] [39] [42] [43] [44] [46] [47] [58]
[59] [61] [96] [189] [17] [45] [16] [21] [22])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exaccent.mf
 Ok [0] [1] [2] [3] [4] [5] [6] [7] [8] [9] [10] [11] [12])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exsign.mf
 Ok [24] [34] [35] [36] [37] [64] [191] [159])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrlig.mf
 Ok [25] [26] [28] [27] [29] [30] [31])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exromp.mf
 Ok [38] [63] [190])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrulett.mf
 Ok [65] [66] [67] [68] [69] [70] [71] [72] [73] [74] [75] [76] [77] [78]
[79] [80] [81] [82] [83] [84] [85] [86] [87] [88] [89] [90])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrllett.mf
 Ok [97] [98] [99] [100] [101] [102] [103] [104] [105] [106] [107] [108]
[109] [110] [111] [112] [113] [114] [115] [116] [117] [118] [119] [120]
[121] [122])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrdigit.mf
 Ok [48] [49] [50] [51] [52] [53] [54] [55] [56] [57])
(/usr/local/texlive/2022/texmf-dist/fonts/source/jknappen/ec/exrligtb.mf
 Ok) ) ) )
(some charht values had to be adjusted by as much as 0.06375pt)
Font metrics written on ecsx1095.tfm.
Output written on ecsx1095.600gf (256 characters, 50464 bytes).
Transcript written on ecsx1095.log.
mktexpk: /data/.texlive2022/texmf-var/fonts/pk/ljfour/jknappen/ec/ecsx1095.600pk: successfully generated.
Latexmk: Getting log file 'out/main.log'
Latexmk: Biber did't find bib file [main.bib]
Collected warning summary (may duplicate other messages):1
  NONE
Collected error summary (may duplicate other messages):
  pdflatex: Command for 'pdflatex' gave return code 1
      Refer to 'out/main.log' for details
Latexmk: If appropriate, the -f option can be used to get latexmk
  to try to force complete processing.
INFO[13:36:18] Successfully run latex. Code: {<nil> 12}  
```
===

For example, the final image looks like below and is used to reduce compiling time because that
image is only rebuild if the base image changed or you modified the `Dockerfile`.
```text
textool-demo-latex-swpyt1fzzibqm7lxb7rrukqg1i
```
You find the final PDF file next to the `main.tex`, in out example it's called `demo.pdf`:
```text
demo-latex
├── .latex
├── demo.pdf
├── Dockerfile
├── latexmkrc
├── main.bib
├── main.tex
└── out
    ├── main.acn
    ├── main.aux
    ├── main.bbl
    ├── main.bcf
    ├── main.blg
    ├── main.fdb_latexmk
    ├── main.fls
    ├── main.glg
    ├── main.glo
    ├── main.gls
    ├── main.ist
    ├── main.log
    ├── main.lol
    ├── main.out
    ├── main.pdf
    ├── main.run.xml
    ├── main.toc
    └── _minted-main
        ├── 949BF7DF46352D63B16C28DB265BC139BE66D8DAD8160124D3F09E7AA39CC794.pygtex
        ├── C4B5797D4DFB619B0A1E16736F6A254C09DC1E0ED436C2808BF79F64BB7A146D.pygtex
        ├── FAD5CDE6E73A9B1FF7D7DC84F651A868813F6471C2D7EF07FEF2B7666DDC3D20.pygtex
        └── pastie.pygstyle

2 directories, 27 files
```

Use the `clean` command to remove all build files.
```text textool -p demo-latex clean


demo-latex
├── .latex
├── demo.pdf
├── Dockerfile
├── latexmkrc
├── main.bib
└── main.tex

0 directories, 6 files
```

Use the `watch` command to recompile the project after any changes to the project 
(so `.tex` files as well as images etc).

```text textool -p demo-latex watch
INFO[13:37:14] Current binary is the latest version 2.1.4   
INFO[13:37:14] Opening config file for  reading. Path:demo-latex/.latex 
INFO[13:37:14] Now pulling the base image: ghcr.io/alexander-lindner/latex:base 
INFO[13:37:14] Image is already available. Not pulling.     
INFO[13:37:14] It is necessary to build the file before using it 
Step 1/3 : FROM ghcr.io/alexander-lindner/latex:base
 ---> 4feb1d254540
Step 2/3 : RUN tlmgr update --self
 ---> Using cache
 ---> ebe3b47210f0
Step 3/3 : RUN tlmgr install koma-script xetex xstring float fontspec abstract cleveref hyperref hyphen-german babel-german soul listings minted fvextra fancyvrb upquote lineno xcolor catchfile framed etoolbox glossaries mfirstuc etoolbox textcase xfor datatool tracklang xkeyval glossaries-german csquotes biber biblatex
 ---> Using cache
 ---> 6858eab40605
Successfully built 6858eab40605
Successfully tagged textool-tmp-xv/bnyoroi8lrb7vrggiiihfam:latest
INFO[13:37:14] Creating the docker container                
INFO[13:37:14] The final file does not exists. For watching it has to exists. Therefore, a normal build is executed before... 
INFO[13:37:14] Creating the docker container                
INFO[13:37:14] Starting the docker container                
Rc files read:
  latexmkrc
Latexmk: Run number 1 of rule 'pdflatex'
This is pdfTeX, Version 3.141592653-2.6-1.40.24 (TeX Live 2022) (preloaded format=pdflatex)
 \write18 enabled.
entering extended mode
/usr/local/bin/pygmentize

...

INFO[13:37:29] Successfully run latex. Code: {<nil> 12}     
INFO[13:37:29] Starting the docker container                
INFO[13:37:29] Adding background watcher for changed files... 
Rc files read:
  latexmkrc
======= Need to update make_preview_continuous for target files
Viewing pdf
Latexmk: Run number 1 of rule 'biber out/main'
Latexmk: Biber did't find bib file [main.bib]
Latexmk: Run number 1 of rule 'pdflatex'
This is pdfTeX, Version 3.141592653-2.6-1.40.24 (TeX Live 2022) (preloaded format=pdflatex)
 \write18 enabled.
entering extended mode
/usr/local/bin/pygmentize
Latexmk: Getting log file 'out/main.log'
INFO[13:37:37] Copy re-rendered file....                    
Latexmk: Biber did't find bib file [main.bib]
Latexmk: Run number 2 of rule 'pdflatex'
This is pdfTeX, Version 3.141592653-2.6-1.40.24 (TeX Live 2022) (preloaded format=pdflatex)
 \write18 enabled.
entering extended mode
/usr/local/bin/pygmentize
Latexmk: Getting log file 'out/main.log'
INFO[13:37:39] Copy re-rendered file....                    
Latexmk: Biber did't find bib file [main.bib]
Latexmk: Failure to make the files correctly
    ==> You will need to change a source file before I do another run <==
Collected warning summary (may duplicate other messages):1
  NONE
Collected error summary (may duplicate other messages):
  pdflatex: Command for 'pdflatex' gave return code 1
      Refer to 'out/main.log' for details
=== Watching for updated files. Use ctrl/C to stop ...
```

Use `[CTRL] + [C]` to stop watching.


Finally, use the `bib` command to quickly add an url as bibliography:

```text textool -p demo-latex bib -u https://en.wikipedia.org/wiki/Latex
INFO[13:36:53] Current binary is the latest version 2.1.4   
INFO[13:36:53] Fetching url...                              
INFO[13:36:53] This bibtex entry was successfully added to your main.bib file. 
```


