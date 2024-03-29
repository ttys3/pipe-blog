FROM node:10.15.3 as NODE_BUILD
WORKDIR /go/src/github.com/b3log/pipe/
ADD . /go/src/github.com/b3log/pipe/
RUN cd console && npm install && npm run build && cd ../theme && npm install && npm run build && \
    rm -rf node_modules && cd ../console && rm -rf node_modules

FROM golang:alpine as GO_BUILD
WORKDIR /go/src/github.com/b3log/pipe/
COPY --from=NODE_BUILD /go/src/github.com/b3log/pipe/ /go/src/github.com/b3log/pipe/
ENV GO111MODULE=on
RUN apk add --no-cache gcc musl-dev git && go build -i -v

FROM alpine:latest
LABEL maintainer="Liang Ding<d@b3log.org>"
WORKDIR /opt/pipe/
COPY --from=GO_BUILD /go/src/github.com/b3log/pipe/ /opt/pipe/
RUN apk add --no-cache ca-certificates tzdata

ENV TZ=Asia/Shanghai
EXPOSE 8089

ENTRYPOINT [ "/opt/pipe/pipe" ]
