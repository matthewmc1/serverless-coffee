FROM golang:1.16-alpine as BUILD

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /coffee-service


FROM scratch
COPY --from=build /coffee-service /coffee-service
EXPOSE 8080
ENTRYPOINT [ "/product-service" ]
