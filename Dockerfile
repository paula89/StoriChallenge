FROM golang:1.22.2
WORKDIR /usr/src/app
COPY . ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/sendEmails

CMD ["sendEmails"]