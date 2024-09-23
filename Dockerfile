FROM golang:alpine

WORKDIR /project/url-shortener

COPY go.* ./

COPY . .
RUN go build -o /project/url-shortener/build/myapp .

EXPOSE 8080
ENTRYPOINT [ "/project/url-shortener/build/myapp" ]
