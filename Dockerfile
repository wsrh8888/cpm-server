FROM golang:1.16-alpine as builder

WORKDIR /docker_cpm
COPY . .

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w CGO_ENABLED=0
RUN go env
RUN go mod tidy
RUN go build -o docker_cpm .

EXPOSE 7777

# CMD [ "./docker_cpm" ]