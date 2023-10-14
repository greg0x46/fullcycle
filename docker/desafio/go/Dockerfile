FROM golang:alpine AS go-builder

WORKDIR /go/src/app

COPY ./main.go .
RUN go build main.go && \
    rm -rf main.go

CMD ["./main"]

FROM scratch
WORKDIR /app
COPY --from=go-builder /go/src/app/main .
CMD [ "./main" ]