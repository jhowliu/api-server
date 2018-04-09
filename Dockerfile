# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN apk update && apk add --no-cache git -y
RUN go get github.com/jhowliu/service github.com/gorilla/mux
RUN cd /src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app
 
# final stage
FROM centurylink/ca-certs
COPY --from=build-env /src/app /
ENTRYPOINT ["/app"]