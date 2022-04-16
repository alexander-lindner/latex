package helper

const MinimalLatex = `\documentclass[ <<twocolumn>>, <<lang>> ]{<<documentclass>>}
\usepackage[utf8]{inputenc}
\usepackage[T1]{fontenc}
\usepackage[<<lang>>]{babel}
\usepackage{abstract}
\usepackage{hyperref}
\usepackage{cleveref}
<<glossary>>
<<bibliography>>
<<minted>>
<<glossaryHead>>
\begin{document}
<<content>>
\end{document}
`
const Latexmkrc = `add_cus_dep('glo', 'gls', 0, 'makeglo2gls');
sub makeglo2gls {
    system("makeindex -s '$_[0]'.ist -t '$_[0]'.glg -o '$_[0]'.gls '$_[0]'.glo");
}

$ENV{'OUTPUT_PDF_NAME'} //= "main";
$filename = $ENV{'OUTPUT_PDF_NAME'};
# Output a pdf
$pdf_mode = 1;

$pdflatex = "pdflatex --shell-escape %O %S"; # -jobname=$filename  will break the watch command somehow
# $pdflatex = "pdflatex -jobname=$filename --shell-escape %O %S";
`
const ConfigHeader = `##############################################
#            textools config file            #
# https://github.com/alexander-lindner/latex #
##############################################`
const ConfigFile = ConfigHeader + `
texFile: "main.tex"
fileName: "<<fileName>>"
docker {
	file: "<<dockerFile>>"
}
features {
	documentclass: <<documentclass>>,
	glossary: <<glossary>>,
	minted: <<minted>>,
	bibliography: <<bibliography>>,
	twocolumn: <<twocolumn>>,
	lang: [ <<lang>> ],
	examples: <<examples>>,
}
`
const GlossariesTex = `%----------------------------------------------------------------------------------------------------------------------------
% ------------------------------------------------------ Setup Glossaries -------------------------------------------------------
%----------------------------------------------------------------------------------------------------------------------------
\usepackage[acronyms,toc,nonumberlist,section=section]{glossaries}
% ------------------------------------------------------- END -------------------------------------------------------

`
const GlossariesContentHead = `
% ------------------------------------------------------- EXAMPLE Glossaries -------------------------------------------------------
\makeglossaries
\newglossaryentry{latex}{
name={LaTeX},
description={LaTeX (/ˈlɑːtɛx/ LAH-tekh or /ˈleɪtɛx/ LAY-tekh, often stylized as LaTeX) is a software system for document preparation.}
}
\newacronym{ieee}{IEEE}{Institute of Electrical and Electronics Engineers }
% ------------------------------------------------------- END -------------------------------------------------------

`
const GlossariesContent = `
\section{Glossaries}
Using glossaries is very simple.
Define them before and reference it like this:
\Gls{latex}
This phrase is also shown in the list of glossary entries.

You may also add short abbreviations like the same way:
\begin{itemize}
    \item First usage: \Gls{ieee}
    \item Second usage: \Gls{ieee}
    \item Third usage: \Gls{ieee}
\end{itemize}

`
const GlossariesContentEnd = `
    \printglossary
`
const BibliographyTex = `%----------------------------------------------------------------------------------------------------------------------------
% --------------------------------------------------- Setup bibliography ----------------------------------------------------
%----------------------------------------------------------------------------------------------------------------------------
\usepackage[backend=biber,style=numeric]{biblatex}
\bibliography{main}
% ------------------------------------------------------- END -------------------------------------------------------

`

const BibliographyContent = `
\section{Bibliography}

BibLaTeX is a complete reimplementation of the bibliographic facilities provided by LaTeX~\cite{site:ctan_paket_biblatex}. 
Formatting of the bibliography is entirely controlled by LaTeX macros, and a working knowledge of LaTeX should be sufficient to design new bibliography and citation styles.
BibLaTeX uses its own data backend program called "biber" to read and process the bibliographic data.
With biber, BibLaTeX has many features rivalling or surpassing other bibliography systems. To mention a few:
\begin{itemize}
\item    Full Unicode support
\item     Highly customisable sorting using the Unicode Collation Algorithm + CLDR tailoring
\item     Highly customisable bibliography labels
\item     Complex macro-based on-the-fly data modification without changing your data sources
\item     A tool mode for transforming bibliographic data sources
\item     Multiple bibliographies and lists of bibliographic information in the same document with different sorting
\item     Highly customisable data source inheritance rules
\item     Polyglossia and babel suppport for automatic language switching for bibliographic entries and citations
\item     Automatic bibliography data recoding (UTF-8 -> latin1, LaTeX macros -> UTF-8 etc)
\item     Remote data sources
\item     Highly sophisticated automatic name and name list disambiguation system
\item     Highly customisable data model so users can define their own bibliographic data types
\item     Validation of bibliographic data against a data model
\item     Subdivided and/or filtered bibligraphies, bibliographies per chapter, section etc.
\end{itemize}

`
const BibliographyContentEnd = `
\printbibliography

`

const MintedTex = `%----------------------------------------------------------------------------------------------------------------------------
% ------------------------------------------------------ Setup Minted -------------------------------------------------------
% Don't forget: \begin{minted} can't be used in "\newcommand"
%----------------------------------------------------------------------------------------------------------------------------
\usepackage{fvextra}
\usepackage{csquotes}
\usepackage{listings}
\usepackage[outputdir=./out]{minted}
\usepackage{amssymb}
\usemintedstyle{pastie}

\newenvironment{longlisting}{\captionsetup{type=listing}}{}
\newcommand{\codefile}[4]{
    \begin{longlisting}
        \inputminted[escapeinside=||,linenos,breaklines,breakanywhere,frame=single,tabsize=2,obeytabs]{#1}{#2}
        \caption{#3}
        #4
    \end{longlisting}
}
\newcommand{\codefilelines}[6]{
    \begin{longlisting}
        \inputminted[escapeinside=||,linenos,breaklines,breakanywhere,frame=single,tabsize=2,obeytabs,firstline=#5,lastline=#6]{#1}{#2}
        \caption{#3}
        #4
    \end{longlisting}
}
\newminted[bashCode]{bash}{breaklines,linenos,breakanywhere,frame=single,xleftmargin=20pt} % \begin{bashCode} \end{bashCode}

%-------------------------------
% ------ Markdownishes Highlight
%-------------------------------
\usepackage{soul}
\newcommand{\highlight}[1]{
    \begingroup
    \sethlcolor{lightgray}%
    \hl{#1}%
    \endgroup
}

%-------------------------
% ------ Inlinecode ------
%-------------------------
\newcommand{\inlinecode}[2]{\setlength{\fboxsep}{2pt}\colorbox{lightgray}{\mintinline[escapeinside=||,linenos,breaklines,breakanywhere,frame=single,tabsize=2,obeytabs]{#1}{#2}}}
% ------------------------------------------------------- END -------------------------------------------------------

`
const MintedContent = `
\section{Listings (Minted)}
This is an example of the listing feature:

\begin{listing}
        \caption[How to add a progressbar to a shell script]{When scripting in bash or any other shell in *NIX, while running a command that will take more than a few seconds, a progress bar is needed. This is an example of such a bar. Source: \url{https://stackoverflow.com/a/28044986/9479657}}
\begin{bashCode}
#!/bin/bash
function ProgressBar {
# Process data
    let _progress=(${1}*100/${2}*100)/100
    let _done=(${_progress}*4)/10
    let _left=40-$_done
    _fill=$(printf "%${_done}s")
    _empty=$(printf "%${_left}s")
printf "\rProgress : [${_fill// /#}${_empty// /-}] ${_progress}%%"

}
_start=1
_end=100

for number in $(seq ${_start} ${_end})
do
    sleep 0.1
    ProgressBar ${number} ${_end}
done
printf '\nFinished!\n'
\end{bashCode}
\end{listing}

Another possibility is to use \inlinecode{tex}{ \inlinecode{c}{...} }, which produces exactly the output you saw.

Also, you can use a common \highlight{listing} environment like showed in \cref{listing:3}.
\begin{listing}
    \begin{minted}[linenos,breaklines,breakanywhere,xleftmargin=20pt]{go}
func PathExists(path string) bool {
    if _, err := os.Stat(path); err == nil {
        return true
    } else if errors.Is(err, os.ErrNotExist) {
        return false
    } else {
        log.Panic("Couldn't fetch stats for "+path, err)
        return false
    }
}
    \end{minted}
\caption{Example of a go module, which checks, if a path exists.}
\label{listing:3}

\end{listing}
`
const MintedContentEnd = `
\listoflistings
`
const ExampleContentDefault = `
	\title{An example file}
    \subtitle{This content shows how to some of the features}
    \author{John Doe}

    \maketitle

    \begin{abstract}
		An funny and short abstract
	\end{abstract}

    \tableofcontents
`

const BiberTex = `@software{Lindner_alindner_s_latex_collection_2019,
author = {Lindner, Alexander},
license = {GPL-3.0-or-later},
month = {3},
title = {{alindner's latex collection}},
url = {https://github.com/alexander-lindner/latex},
version = {1.0},
year = {2019}
}
@online{ site:ctan_paket_biblatex,
	title = { CTAN: Paket BibLaTeX },
	date = { 2022-03-02 },
	url = { https://ctan.org/pkg/biblatex?lang=de },
	note = { [accessed 02-March-2022] }
}
`
const MinimalDockerFile = `FROM {{image}}

RUN tlmgr update --self
RUN tlmgr install {{packages}}

`
