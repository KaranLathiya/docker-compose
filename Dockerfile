FROM golang:1.21.0-alpine
# MAINTAINER : Karan
RUN mkdir /app
ADD . /app

# RUN go mod tidy
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
# RUN go get -u github.com/gin-gonic/gin
# CMD ["go","run","main.go"]
EXPOSE 8080