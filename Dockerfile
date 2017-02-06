FROM golang:1.7-alpine
MAINTAINER Irakli Nadareishvili

ENV REFRESHED_AT 2017-01-25_0600_EST

# Note: if you change GOPATH here, you need to change in .env for docker-compose.yml, too
ENV PORT=3737
ENV APP_NAME=github.com/inadarei/justgo
ENV GOPATH=/go
ENV PATH=${GOPATH}/bin:${PATH}
ENV APP_USER=appuser
ENV APP_ENV=production

ADD . ${GOPATH}/src/${APP_NAME}
WORKDIR ${GOPATH}/src/${APP_NAME}

USER root

RUN adduser -s /bin/false -u 7007 -D ${APP_USER} \
 && echo "Installing git and ssh support" \ 
 && apk update && apk upgrade \
 && apk add --no-cache bash git openssh \
 && echo "Installing infrastructural go packages..." \
 && go get github.com/codegangsta/gin \
 && go get github.com/tools/godep \
 && go get -u github.com/Masterminds/glide \
 && glide install \
 && echo "Fixing permissions..." \
 && chown -R ${APP_USER} ${GOPATH} \
 && echo "Cleaning up installation caches to reduce image size" \
 && rm -rf /root/src /tmp/* /usr/share/man /var/cache/apk/* 

USER ${APP_USER}

EXPOSE 3000