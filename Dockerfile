FROM golang:1.7-alpine
MAINTAINER Irakli Nadareishvili

ENV REFRESHED_AT 2017-01-21_0600_EST

# Note: if you change APP_PATH and GOPATH here, you need to change docker-compose.yml, too
ENV APP_PATH=/opt/application \
    PATH=/opt/application/bin:${PATH} \
    PORT=3737 \
    GOPATH=/opt/application \
    APP_USER=appuser \
    APP_ENV=production

ADD ./ ${APP_PATH}
WORKDIR ${APP_PATH}
EXPOSE 3000

USER root

RUN adduser -s /bin/false -u 7007 -D ${APP_USER} \
 && mkdir -p ${APP_PATH}/bin \ 
 && chown -R ${APP_USER} ${APP_PATH} \
 && echo "Installing git and ssh support" \ 
 && apk update && apk upgrade \
 && apk add --no-cache bash git openssh \
 && echo "Installing go packages..." \
 && go get github.com/codegangsta/gin \
 && echo "Cleaning up caches to reduce image size" \
 && rm -rf /root/src /tmp/* /usr/share/man /var/cache/apk/* 

USER ${APP_USER}