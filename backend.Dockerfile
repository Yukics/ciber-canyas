FROM golang:1.19.3

WORKDIR /app

# RUN go install

CMD ["go", "run", "main.go"]
