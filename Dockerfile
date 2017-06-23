FROM golang:1.8
LABEL name "http2-push"
WORKDIR /go/src/app
COPY . .
RUN go-wrapper download
RUN go-wrapper install
EXPOSE 8000
ENTRYPOINT ["go-wrapper", "run"]