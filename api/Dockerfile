FROM golang:latest

WORKDIR /go/src/app
COPY . .

# RUN go get -u github.com/swaggo/swag/cmd/swag
# RUN swag init
RUN go get github.com/pilu/fresh

CMD [ "fresh" ]