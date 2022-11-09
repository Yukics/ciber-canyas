FROM golang:1.19.3

WORKDIR /app

CMD ["go", "run", "main.go"]
