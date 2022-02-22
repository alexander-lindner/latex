FROM ubuntu:latest
MAINTAINER Alexander Lindner <25225552+alexander-lindner@users.noreply.github.com>
ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update -q && \
    apt-get upgrade -y && \
    apt-get install -qy build-essential wget libfontconfig1 curl git inotify-tools imagemagick inkscape python3-pip && \
    rm -rf /var/lib/apt/lists/*

# Install TexLive with scheme-basic
RUN wget http://mirror.ctan.org/systems/texlive/tlnet/install-tl-unx.tar.gz; \
	mkdir /install-tl-unx; \
	tar -xvf install-tl-unx.tar.gz -C /install-tl-unx --strip-components=1; \
    echo "selected_scheme scheme-basic" >> /install-tl-unx/texlive.profile; \
	/install-tl-unx/install-tl -profile /install-tl-unx/texlive.profile; \
    rm -r /install-tl-unx; \
	rm install-tl-unx.tar.gz

ENV PATH="/usr/local/texlive/2021/bin/x86_64-linux:${PATH}"

ENV HOME /data
WORKDIR /data

RUN pip3 install pygments
RUN tlmgr install latexmk

ADD entrypoint.sh /
RUN chmod +x /entrypoint.sh
ENTRYPOINT /entrypoint.sh