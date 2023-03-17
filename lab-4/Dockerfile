FROM golang:1.20.2-bullseye

RUN useradd -ms /bin/bash go

USER go

WORKDIR /home/go/app

COPY . .

RUN go build -o build/app cmd/main.go

CMD ["./build/app"]

EXPOSE 8080
