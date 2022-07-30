FROM golang:1.18 as build
WORKDIR /app
COPY . ./
ENV CGO_ENABLED=0
RUN go mod vendor
RUN go build -o julius

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add tzdata
RUN echo "America/Sao_Paulo" > /etc/timezone
WORKDIR /root/
COPY --from=build /app/julius ./
COPY frontend ./frontend
ENV GIN_MODE=release
RUN ls -lsh
RUN pwd
CMD ["/root/julius"] 