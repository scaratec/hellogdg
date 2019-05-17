FROM golang:latest
WORKDIR /go/src/github.com/scaratec/hellogdg/
COPY . .
RUN CGO_ENABLED=0 go build -v -o hellogdg

FROM scratch
COPY --from=0 /go/src/github.com/scaratec/hellogdg/hellogdg .
ENTRYPOINT ["/hellogdg"]
