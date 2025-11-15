FROM golang:1.22-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o app main.go

FROM alpine
WORKDIR /
COPY --from=build /src/app .
EXPOSE 8080
CMD ["./app"]
