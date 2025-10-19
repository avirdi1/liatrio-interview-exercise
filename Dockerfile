FROM golang:1.22 

WORKDIR /app

# these dependency files don't change so docker uses cache 
COPY go.mod go.sum ./

RUN go mod download

COPY . .

# makes port 80 available so the app can be reached by web browsers
EXPOSE 80

# compiles go code into an executable file for ease of access
RUN go build -o /liatrio_exercise .

CMD ["/liatrio_exercise"]