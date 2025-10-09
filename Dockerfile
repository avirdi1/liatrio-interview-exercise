FROM golang:1.22 

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV PORT=3000

EXPOSE 3000

RUN go build -o /liatrio_exercise .

CMD ["/liatrio_exercise"]