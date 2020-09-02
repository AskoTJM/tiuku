FROM golang:1.15

WORKDIR /tiuku
#COPY go.mod ./
#RUN go mod download
COPY . .
RUN go build -o main


EXPOSE 8080

CMD ["./main"]
#RUN tail -fn0 /dev/null