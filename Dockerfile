FROM golang:1.17 as builder
RUN CGO_ENABLED=0 go get github.com/buzzsurfr/sonobuoy/...

FROM scratch
COPY --from=builder /go/bin/sonobuoy /sonobuoy
EXPOSE 2869
CMD ["/sonobuoy"]