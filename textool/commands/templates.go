package commands

const MinimalLatex = `\documentclass[ <<twocolumn>>, <<lang>> ]{<<documentclass>>}
\usepackage[utf8]{inputenc}
\usepackage[T1,EU1]{fontenc}
\usepackage[<<lang>>]{babel}
<<glossary>>
<<bibliography>>
<<minted>>
\begin{document}
(Type your content here.)
\end{document}
`
const latexmkrc = `add_cus_dep('glo', 'gls', 0, 'makeglo2gls');
sub makeglo2gls {
    system("makeindex -s '$_[0]'.ist -t '$_[0]'.glg -o '$_[0]'.gls '$_[0]'.glo");
}
$pdflatex=q/xelatex -synctex=1 %O %S/
`

const configFile = `############################################
# alindner's Latex Tools                   #
############################################

fileName: "<<fileName>>"
docker {
	image: "<<dockerImage>>"
}
features {
	documentclass: <<documentclass>>,
	glossary: <<glossary>>,
	minted: <<minted>>,
	bibliography: <<bibliography>>,
	twocolumn: <<twocolumn>>,
	lang: [ <<lang>> ],
}
`
const glossariesTex = `%----------------------------------------------------------------------------------------------------------------------------
% ------------------------------------------------------ Setup Glossaries -------------------------------------------------------
%----------------------------------------------------------------------------------------------------------------------------
\usepackage[acronyms,toc,nonumberlist,section=section]{glossaries}
% ------------------------------------------------------- END -------------------------------------------------------

`
const bibliographyTex = `%----------------------------------------------------------------------------------------------------------------------------
% --------------------------------------------------- Setup bibliography ----------------------------------------------------
%----------------------------------------------------------------------------------------------------------------------------
\usepackage[backend=biber]{biblatex}
\bibliography{main}
\bibliographystyle{ieeetr}
% ------------------------------------------------------- END -------------------------------------------------------

`

const mintedTex = `%----------------------------------------------------------------------------------------------------------------------------
% ------------------------------------------------------ Setup Minted -------------------------------------------------------
% Don't forget: \begin{minted} can't be used in "\newcommand"
%----------------------------------------------------------------------------------------------------------------------------
\usepackage{listings}
\usepackage[outputdir=../out]{minted}
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
const exampleContentDefault = ``

const biberTex = `@software{Lindner_alindner_s_latex_collection_2019,
author = {Lindner, Alexander},
license = {GPL-3.0-or-later},
month = {3},
title = {{alindner's latex collection}},
url = {https://github.com/alexander-lindner/latex},
version = {1.0},
year = {2019}
}`

const MinimalDockerFile = `FROM ghcr.io/alexander-lindner/latex:base

RUN tlmgr update --self
RUN tlmgr install minted

`
