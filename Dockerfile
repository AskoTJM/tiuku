FROM golang:1.15.1

WORKDIR /tiuku
VOLUME  ["/tiuku"]



RUN go get -d github.com/gorilla/mux
#GORM V2 uses these
#RUN go get -u gorm.io/gorm
#RUN go get -u gorm.io/driver/mysql

#GORM V1 uses these
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/jinzhu/gorm/dialects/mysql

#COPY go.mod ./
#RUN go mod download
#COPY . .

#RUN go build -o main
#EXPOSE 8080
#ENTRYPOINT [ "go","./main" ]
#CMD tail -fn0 /dev/null
#CMD ./main && tail -f /dev/null
#CMD [./main]
#RUN /ail -fn0 /dev/null