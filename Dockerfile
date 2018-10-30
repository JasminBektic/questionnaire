FROM golang

RUN rm -rf /go/src/github.com/jasminbektic/questionnaire
COPY . /go/src/github.com/jasminbektic/questionnaire
WORKDIR /go/src/github.com/jasminbektic/questionnaire

RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gorilla/mux
RUN go get -u golang.org/x/crypto/bcrypt

CMD go run main.go

EXPOSE 8080