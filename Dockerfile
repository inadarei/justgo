FROM golang:1.7-alpine
MAINTAINER Irakli Nadareishvili

# Note: if you change GOPATH here, you need to change in .env for docker-compose.yml, too
ENV PORT=3737
ENV GOPATH=/go
ENV PATH=${GOPATH}/bin:${PATH}
ENV APP_USER=appuser
ENV APP_ENV=production

COPY . ${GOPATH}/src/app
WORKDIR ${GOPATH}/src/app

USER root

RUN adduser -s /bin/false -u 7007 -D ${APP_USER} \
 && echo "Installing git and ssh support" \ 
 && apk update && apk upgrade \
 && apk add --no-cache bash git openssh \
 && echo "Installing infrastructural go packages…" \
 ## && go get github.com/pilu/fresh \
 ## && go get github.com/codegangsta/gin \
 && go get github.com/githubnemo/CompileDaemon \
 && go get github.com/Masterminds/glide \
 ## If you are checking-in your deps, you can remove following line
 && echo "Building project…" \
 && glide install \
 && go-wrapper install \
 && echo "Fixing permissions..." \
 && chown -R ${APP_USER} ${GOPATH} \
 && echo "Cleaning up installation caches to reduce image size" \
 && rm -rf /root/src /tmp/* /usr/share/man /var/cache/apk/* 

USER ${APP_USER}

EXPOSE 3000

CMD ["go-wrapper", "run"]