#!/usr/bin/env bash

mkdir -p {out,src,.run}

curl https://raw.githubusercontent.com/alexander-lindner/latex/master/installer/run.sh > run.sh
chmod +x run.sh
touch src/main.bib
touch src/main.tex



cat > src/main.tex <<EOF
\documentclass[a4paper]{report}
\usepackage[utf8]{inputenc}
\usepackage[T1]{fontenc}
\usepackage[backend=biber,style=ieee,backref]{biblatex}
\usepackage[acronyms,toc,nonumberlist,section=section]{glossaries}
\usepackage[outputdir=../out]{minted}
\addbibresource{main.bib}

\makeglossaries

\begin{document}
    %your latex file
\end{document}
EOF