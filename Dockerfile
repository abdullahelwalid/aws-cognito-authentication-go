FROM golang:1.22.2-alpine3.19

WORKDIR app/

COPY . .

RUN go install github.com/air-verse/air@latest

RUN go mod download

#CMD ["go", "run", "cmd/main/main.go"]
CMD ["air", "-c", ".air.toml"]
