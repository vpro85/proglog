FROM golang:1.17-alpine AS build
WORKDIR /go/src/proglog
COPY . .
RUN CG0_ENABLED=0 go build -o /go/bin/proglog ./cmd/proglog

FROM scratch
COPY --from=build /go/bin/proglog /bin/proglog
ENTRYPOINT ["/bin/proglog"]