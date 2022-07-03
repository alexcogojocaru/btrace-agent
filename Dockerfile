FROM golang:1.17-alpine

WORKDIR /data

COPY . .
RUN go build -o /data/bin-agent

CMD [ "/data/bin-agent" ]
