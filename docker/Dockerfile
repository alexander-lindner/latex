FROM ubuntu:latest
MAINTAINER Alexander Lindner <25225552+alexander-lindner@users.noreply.github.com>
ENV DEBIAN_FRONTEND noninteractive
LABEL org.opencontainers.image.title="latex compiler"
LABEL org.opencontainers.image.version="texlive-2023"
LABEL org.opencontainers.image.licenses="GPL-2.0,LPPL-1.3c"
LABEL org.opencontainers.image.url="https://github.com/alexander-lindner/latex"
LABEL org.opencontainers.image.description="A minimal, basic image for compiling .tex documents. You can use `tlmgr` to install packages. If you want all dependencies, use the `:full` image."

RUN apt-get update -q && \
    apt-get upgrade -y && \
    apt-get install -qy build-essential wget libfontconfig1 curl git inotify-tools imagemagick inkscape python3-pip ruby-hocon python3-pygments && \
    rm -rf /var/lib/apt/lists/*

# Install TexLive with scheme-basic
RUN wget https://mirror.ctan.org/systems/texlive/tlnet/install-tl-unx.tar.gz; \
	mkdir /install-tl-unx; \
	tar -xvf install-tl-unx.tar.gz -C /install-tl-unx --strip-components=1; \
    echo "selected_scheme scheme-basic" >> /install-tl-unx/texlive.profile; \
	/install-tl-unx/install-tl -profile /install-tl-unx/texlive.profile; \
    rm -r /install-tl-unx; \
	rm install-tl-unx.tar.gz && \
    rm -rf /usr/local/texlive/2023/texmf-dist/doc&& \
    rm -rf /usr/local/texlive/2023/texmf-dist/source

ENV PATH="/usr/local/texlive/2022/bin/x86_64-linux:${PATH}"

ENV HOME /data
WORKDIR /data

RUN tlmgr install latexmk && \
    rm -rf /usr/local/texlive/2022/texmf-dist/doc

RUN echo "#!/bin/bash" > /usr/bin/acroread && chmod +x /usr/bin/acroread

ADD entrypoint.sh /
RUN chmod +x /entrypoint.sh

CMD ["run"]
ENTRYPOINT ["/entrypoint.sh"]
