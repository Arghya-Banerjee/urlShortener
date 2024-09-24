FROM golang:alpine

WORKDIR /project/url-shortener

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /project/url-shortener/build/myapp .

EXPOSE 8000

ENTRYPOINT [ "/project/url-shortener/build/myapp" ]
