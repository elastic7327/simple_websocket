FROM alpine:latest

MAINTAINER Daniel <dan@fastcall.co.kr>

USER root

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh 

# go mercurial subversion openssh-client ca-certificates

# 이쪽에서 빌드를 해도 좋을 것 같습니다.

ADD . /simple_websocket

ENTRYPOINT ["/bin/bash", "-c", "/simple_websocket/server"]

CMD ["/simple_websocket/server"]

ENV PORT 8000

EXPOSE 8000
