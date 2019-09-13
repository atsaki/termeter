#!/usr/bin/env -S docker build --compress -t pvtmert/termeter -f

FROM debian as build

RUN apt update
RUN apt install -y gcc git curl

RUN curl -skL https://dl.google.com/go/go1.13.linux-amd64.tar.gz \
	| tar --strip-components=0 -xzC /usr/local

ENV PATH "$PATH:/usr/local/go/bin"

WORKDIR /root/go/src/github.com/atsaki/termeter
COPY ./ ./
RUN echo get build test install | xargs -n1 | xargs -n1 -I% -- go % .
RUN echo get build test install | xargs -n1 | xargs -n1 -I% -- go % ./cmd/termeter

FROM debian
WORKDIR /data
COPY --from=build /root/go/bin/termeter /usr/local/bin/termeter
ENTRYPOINT [ "/usr/local/bin/termeter" ]
CMD        [ ]
