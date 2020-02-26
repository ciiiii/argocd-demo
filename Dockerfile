FROM golang:1.13-alpine
WORKDIR /work
ADD . .
RUN go build -o /bin/echoServer .
WORKDIR /
RUN rm -rf /work
EXPOSE 8000
CMD ["/bin/echoServer"]