FROM golang:1.15-alpine3.13 AS build

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories;\
    apk add --no-cache make git
COPY . /go/src/git.inspur.com/middleware/idx-component/docker/filesync/idx-ceco
WORKDIR /go/src/git.inspur.com/middleware/idx-component/docker/filesync/idx-ceco
RUN go env -w GOPROXY=https://goproxy.cn,direct
ARG VERSION 
ARG GIT_INFO
RUN VERSION=$VERSION gitInfo=$GIT_INFO make filesync-controller noRace=T


FROM alpine:3.13

COPY --from=build /go/src/git.inspur.com/middleware/idx-component/docker/filesync/idx-ceco/filesync-controller /filesync-controller

CMD ["/filesync-controller"]
