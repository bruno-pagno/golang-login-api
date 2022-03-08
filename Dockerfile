FROM golang

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY .env ./
COPY *.go ./

RUN go build -o /output

CMD ["/output"]